package server_test

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/aria3ppp/watch-server/internal/app"
	"github.com/aria3ppp/watch-server/internal/config"
	"github.com/aria3ppp/watch-server/internal/dto"
	"github.com/aria3ppp/watch-server/internal/hasher"
	"github.com/aria3ppp/watch-server/internal/repo"
	"github.com/aria3ppp/watch-server/internal/search"
	appServer "github.com/aria3ppp/watch-server/internal/server"
	"github.com/aria3ppp/watch-server/internal/token"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// To run this test suite set TEST_E2E env
const ENV_TEST_E2E = "TEST_E2E"

type SetupOpt int

const (
	OptEnableLogger SetupOpt = 1 << iota
	OptEnableDefaultUser
	OptEnableDefaultSeries
)

type Defaults struct {
	user   *DefaultUser
	series *DefaultSeries
}
type DefaultUser struct {
	id          int
	email       string
	password    string
	auth        string
	refreshAuth string
	reqObject   *dto.UserCreateRequest
}
type DefaultSeries struct {
	id        int
	reqObject *dto.SeriesCreateRequest
}

// setup test cases
func setup(
	options ...SetupOpt,
) (testServer *httptest.Server, appInstance *app.Application, defaults *Defaults, teardownFunc func() error, err error) {
	var opts SetupOpt
	if len(options) > 0 {
		opts = options[0]
	}

	// run migrations
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return nil, nil, nil, nil, fmt.Errorf(
			"postgres.WithInstance error: %w",
			err,
		)
	}
	migrator, err := migrate.NewWithDatabaseInstance(
		"file://../../migrations",
		"postgres", driver,
	)
	if err != nil {
		return nil, nil, nil, nil, fmt.Errorf(
			"migrate.NewWithDatabaseInstance error: %w",
			err,
		)
	}
	err = migrator.Up()
	if err != nil && err != migrate.ErrNoChange {
		return nil, nil, nil, nil, fmt.Errorf("migrate up error: %w", err)
	}

	// set server to run in production mode
	config.Config.Servic.Server.Production = true

	// initialize server
	repo := repo.NewRepository(db)
	hasher := hasher.NewBcrypt()
	tokenService := token.NewJWT(
		token.JWTConfig{
			Key:           []byte(config.Config.Servic.Token.SecretKey),
			SigningMethod: jwt.SigningMethodHS512,
			AccessDuration: time.Minute * time.Duration(
				config.Config.Servic.Token.Access.Duration.InMinutes,
			),
			RefreshDuration: time.Minute * time.Duration(
				config.Config.Servic.Token.Refresh.Duration.InMinutes,
			),
		},
	)
	esClient, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{"http://localhost:9200"},
	})
	if err != nil {
		return nil, nil, nil, nil, fmt.Errorf(
			"elasticsearch.NewDefaultClient error: %w",
			err,
		)
	}
	searchService, err := search.NewElasticSearch(esClient)
	if err != nil {
		return nil, nil, nil, nil, fmt.Errorf(
			"search.NewElasticSearch error: %w",
			err,
		)
	}
	appInstance = app.NewApplication(repo, tokenService, searchService, hasher)
	echo := echo.New()
	logger := zap.NewNop()
	if OptEnableLogger&opts != 0 {
		logger, err = zap.NewDevelopment()
		if err != nil {
			return nil, nil, nil, nil, fmt.Errorf(
				"zap.NewDevelopment error: %w",
				err,
			)
		}
	}
	server := appServer.NewServer(appInstance, echo, tokenService, logger)
	testServer = httptest.NewServer(server.GetHandler())

	var defaultUser *DefaultUser
	if OptEnableDefaultUser&opts != 0 || OptEnableDefaultSeries&opts != 0 {
		var (
			email    = "frank@prog.net"
			password = "pa$$W0RD1"
		)
		req := &dto.UserCreateRequest{Email: email, Password: password}
		id, err := appInstance.UserCreate(
			context.Background(),
			req,
		)
		if err != nil {
			return nil, nil, nil, nil, err
		}
		accessToken, refreshToken, err := appInstance.UserLogin(
			context.Background(),
			&dto.UserLoginRequest{
				Email:    email,
				Password: password,
			},
		)
		if err != nil {
			return nil, nil, nil, nil, err
		}
		auth := "Bearer " + accessToken
		refreshAuth := "Bearer " + refreshToken

		defaultUser = &DefaultUser{
			id:          id,
			email:       email,
			password:    password,
			auth:        auth,
			refreshAuth: refreshAuth,
			reqObject:   req,
		}
	}

	var defaultSeries *DefaultSeries
	if OptEnableDefaultSeries&opts != 0 {
		req := &dto.SeriesCreateRequest{Title: "default series"}
		id, err := appInstance.SeriesCreate(
			context.Background(),
			defaultUser.id,
			req,
		)
		if err != nil {
			return nil, nil, nil, nil, err
		}
		defaultSeries = &DefaultSeries{id: id, reqObject: req}
	}

	// prepare teardown
	teardownFunc = func() error {
		// drop migrations
		err = migrator.Drop()
		if err != nil {
			return err
		}
		// close server
		testServer.Close()
		return nil
	}

	defaults = &Defaults{
		user:   defaultUser,
		series: defaultSeries,
	}
	return testServer, appInstance, defaults, teardownFunc, nil
}

var db *sql.DB

func TestMain(m *testing.M) {
	// run only when TEST_E2E env is set
	if os.Getenv(ENV_TEST_E2E) == "" {
		fmt.Printf(
			"end-2-end tests skipped: to enable, set %s env!\n",
			ENV_TEST_E2E,
		)
		return
	}

	os.Setenv("DSN", "")
	err := config.Load(filepath.Join("..", "..", "config.yaml"))
	if err != nil {
		log.Fatalf("Failed loading configs: %s", err)
	}
	os.Setenv("DSN", fmt.Sprintf(
		"postgres://%s:%s@localhost:%s/%s?sslmode=disable",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_DB"),
	))

	// // uses a sensible default on windows (tcp/http) and linux/osx (socket)
	// pool, err := dockertest.NewPool("")
	// if err != nil {
	// 	log.Fatalf("Could not connect to docker: %s", err)
	// }

	// // pulls an image, creates a container based on it and runs it
	// resource, err := pool.RunWithOptions(&dockertest.RunOptions{
	// 	Repository: "postgres",
	// 	Tag:        "14",
	// 	Env: []string{
	// 		"POSTGRES_PASSWORD=secret",
	// 		"POSTGRES_USER=user_name",
	// 		"POSTGRES_DB=dbname",
	// 		"listen_addresses = '*'",
	// 	},
	// }, func(config *docker.HostConfig) {
	// 	// set AutoRemove to true so that stopped container goes away by itself
	// 	config.AutoRemove = true
	// 	config.RestartPolicy = docker.RestartPolicy{Name: "no"}
	// })
	// if err != nil {
	// 	log.Fatalf("Could not start resource: %s", err)
	// }

	// hostAndPort := resource.GetHostPort("5432/tcp")
	// databaseUrl := fmt.Sprintf(
	// 	"postgres://user_name:secret@%s/dbname?sslmode=disable",
	// 	hostAndPort,
	// )

	// log.Println("Connecting to database on url: ", databaseUrl)

	// resource.Expire(
	// 	120,
	// ) // Tell docker to hard kill the container in 120 seconds

	// // exponential backoff-retry, because the application in the container might not be ready to accept connections yet
	// pool.MaxWait = 120 * time.Second
	// if err = pool.Retry(func() error {
	// 	db, err = sql.Open("postgres", databaseUrl)
	// 	if err != nil {
	// 		return err
	// 	}
	// 	return db.Ping()
	// }); err != nil {
	// 	log.Fatalf("Could not connect to docker: %s", err)
	// }

	dsn := os.Getenv("DSN")
	db, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Could not connect to database %q: %s", dsn, err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatalf("Could not ping database %q: %s", dsn, err)
	}

	// Run tests
	code := m.Run()

	// // You can't defer this because os.Exit doesn't care for defer
	// if err := pool.Purge(resource); err != nil {
	// 	log.Fatalf("Could not purge resource: %s", err)
	// }

	os.Exit(code)
}
