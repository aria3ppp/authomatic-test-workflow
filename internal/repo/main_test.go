package repo_test

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/aria3ppp/watch-server/internal/config"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// To run this test suite set TEST_DB_INTEGRATION env
const ENV_TEST_DB_INTEGRATION = "TEST_DB_INTEGRATION"

var db *sql.DB

// setup test cases
func setup() (func() error, error) {
	// run migrations
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatalf("postgres.WithInstance error: %s", err)
	}
	migrator, err := migrate.NewWithDatabaseInstance(
		"file://../../migrations",
		"postgres", driver,
	)
	if err != nil {
		log.Fatalf("migrate.NewWithDatabaseInstance error: %s", err)
	}

	err = migrator.Up()
	if err != nil && err != migrate.ErrNoChange {
		return nil, fmt.Errorf("migrate up error: %w", err)
	}

	// prepare teardown
	teardownFunc := func() error {
		// drop migrations
		err = migrator.Drop()
		if err != nil {
			return err
		}
		return nil
	}

	return teardownFunc, nil
}

func TestMain(m *testing.M) {
	// run only when TEST_DB_INTEGRATION env is set
	if os.Getenv(ENV_TEST_DB_INTEGRATION) == "" {
		fmt.Printf(
			"integration tests skipped: to enable, set %s env!\n",
			ENV_TEST_DB_INTEGRATION,
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
