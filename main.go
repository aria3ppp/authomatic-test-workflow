package main

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/aria3ppp/watch-server/internal/app"
	"github.com/aria3ppp/watch-server/internal/config"
	"github.com/aria3ppp/watch-server/internal/hasher"
	"github.com/aria3ppp/watch-server/internal/repo"
	"github.com/aria3ppp/watch-server/internal/search"
	"github.com/aria3ppp/watch-server/internal/server"
	"github.com/aria3ppp/watch-server/internal/token"
	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

func main() {
	if err := config.Load("config.yaml"); err != nil {
		panic("failed loading configs: " + err.Error())
	}

	var logFile *os.File
	if config.Config.Servic.Server.Production {
		var err error
		logFile, err = os.Create(config.Config.Servic.Server.Logfile)
		if err != nil {
			panic(
				fmt.Sprintf(
					"failed creating log file %q: %s",
					config.Config.Servic.Server.Logfile,
					err,
				),
			)
		}
	}

	logger := newLogger(logFile)

	db, err := sql.Open("postgres", config.Config.Servic.Database.DSN)
	if err != nil {
		logger.Panic("failed openning databse connection", zap.Error(err))
	}
	err = db.Ping()
	if err != nil {
		logger.Panic("failed ping database connection", zap.Error(err))
	}

	repository := repo.NewRepository(db)
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

	var esLogger elastictransport.Logger
	if config.Config.Servic.Server.Production {
		esLogger = &esCustomLogger{logger}
	} else {
		esLogger = &elastictransport.ColorLogger{
			Output:             os.Stdout,
			EnableRequestBody:  true,
			EnableResponseBody: true,
		}
	}
	esClient, err := elasticsearch.NewClient(elasticsearch.Config{
		Logger: esLogger,
	})
	if err != nil {
		logger.Panic("failed creating elasticsearch client", zap.Error(err))
	}
	searchService, err := search.NewElasticSearch(esClient)
	if err != nil {
		logger.Panic(
			"failed new elasticsearch service instantiation",
			zap.Error(err),
		)
	}

	application := app.NewApplication(
		repository,
		tokenService,
		searchService,
		hasher,
	)

	server := server.NewServer(application, echo.New(), tokenService, logger)
	server.Run(":" + strconv.Itoa(int(config.Config.Servic.Server.Port)))
}
