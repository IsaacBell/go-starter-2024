package api

import (
	"crypto/sha1"
	"errors"
	"os"
	"strings"

	"github.com/go-pg/pg/v9"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"

	"github.com/Soapstone-Services/go-template-2024/pkg/api/auth"
	al "github.com/Soapstone-Services/go-template-2024/pkg/api/auth/logging"
	at "github.com/Soapstone-Services/go-template-2024/pkg/api/auth/transport"
	"github.com/Soapstone-Services/go-template-2024/pkg/api/password"
	pl "github.com/Soapstone-Services/go-template-2024/pkg/api/password/logging"
	pt "github.com/Soapstone-Services/go-template-2024/pkg/api/password/transport"
	"github.com/Soapstone-Services/go-template-2024/pkg/api/user"
	ul "github.com/Soapstone-Services/go-template-2024/pkg/api/user/logging"
	ut "github.com/Soapstone-Services/go-template-2024/pkg/api/user/transport"
	authMw "github.com/Soapstone-Services/go-template-2024/pkg/utl/middleware/auth"

	"github.com/Soapstone-Services/go-template-2024/pkg/utl/config"
	"github.com/Soapstone-Services/go-template-2024/pkg/utl/postgres"
	"github.com/Soapstone-Services/go-template-2024/pkg/utl/rbac"
	"github.com/Soapstone-Services/go-template-2024/pkg/utl/secure"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/rs/zerolog"
	"github.com/ziflex/lecho/v3"

	errorUtils "github.com/Soapstone-Services/go-template-2024/pkg/utl/errors"
	"github.com/Soapstone-Services/go-template-2024/pkg/utl/jwt"
	"github.com/Soapstone-Services/go-template-2024/pkg/utl/server"
)

// Start starts the API service
func Start(cfg *config.Configuration) error {
	postgresDb, err := configurePG(cfg)
	errorUtils.CheckErr(err)

	pointer, err := influxClient()
	influx := *pointer
	if err != nil {
		defer influx.Close()
	} else {
		errorUtils.CheckErr(err)
	}

	sec := secure.New(cfg.App.MinPasswordStr, sha1.New())
	rbac := rbac.Service{}

	if !isProd() {
		os.Setenv("JWT_SECRET", strings.Repeat("12345678", 8))
	}

	jwt, err := jwt.New(cfg.JWT.SigningAlgorithm, os.Getenv("JWT_SECRET"), cfg.JWT.DurationMinutes, cfg.JWT.MinSecretLength)
	if err != nil {
		return err
	}

	// var _zLvl zerolog.Level
	var eLvl log.Lvl

	if !cfg.Server.Debug {
		// eLvl, _zLvl = lecho.MatchZeroLevel(zerolog.DebugLevel)
		eLvl, _ = lecho.MatchZeroLevel(zerolog.DebugLevel)
	} else {
		// _zLvl, eLvl = lecho.MatchEchoLevel(log.WARN)
		_, eLvl = lecho.MatchEchoLevel(log.WARN)
	}

	logger := lecho.New(
		os.Stdout,
		lecho.WithLevel(eLvl),
		lecho.WithTimestamp(),
		lecho.WithCaller(),
		lecho.WithFields(map[string]interface{}{
			"service": "go-template-2024",
			"type":    "api",
		}),
	)

	e := server.New()
	e.Logger = logger
	e.Use(lecho.Middleware(lecho.Config{
		Logger:  logger,
		NestKey: "request",
		Enricher: func(c echo.Context, logger zerolog.Context) zerolog.Context {
			return logger
		},
		HandleError: true,
	}))

	// e.Static("/swaggerui", cfg.App.SwaggerUIPath)

	authMiddleware := authMw.Middleware(jwt)

	at.NewHTTP(al.New(auth.Initialize(postgresDb, jwt, sec, rbac), logger), e, authMiddleware)

	v1 := e.Group("/v1")
	// v1.Use(authMiddleware)

	ut.NewHTTP(ul.New(user.Initialize(postgresDb, rbac, sec), logger), v1)
	pt.NewHTTP(pl.New(password.Initialize(postgresDb, rbac, sec), logger), v1)

	server.Start(e, &server.Config{
		Port:                cfg.Server.Port,
		ReadTimeoutSeconds:  cfg.Server.ReadTimeout,
		WriteTimeoutSeconds: cfg.Server.WriteTimeout,
		Debug:               cfg.Server.Debug,
	})

	return nil
}

func configurePG(cfg *config.Configuration) (*pg.DB, error) {
	url := postgresAddr() // try to use .env first

	if url == "" {
		url = cfg.DB.Url // fall back to yaml config
	}

	postgresDb, err := postgres.New(cfg.DB.Url, cfg.DB.Timeout, cfg.DB.LogQueries)
	// fmt.Println("DB Addr: ", postgresDb.Options().Addr)

	return postgresDb, err
}

func influxClient() (*influxdb2.Client, error) {
	const bucket = "test"
	const org = "soapstone"
	// You can generate a Token from the "Tokens Tab" in the UI
	token := os.Getenv("INFLUXDB_TOKEN")

	hostUrl := os.Getenv("HOST")

	if token == "" || hostUrl == "" {
		return nil, errors.New("InfluxDB couldn't be located.")
	}

	client := influxdb2.NewClient(hostUrl, token)
	return &client, nil
}

func isProd() bool {
	return os.Getenv("ENVIRONMENT_NAME") == "production"
}

func postgresAddr() string {
	user   := os.Getenv("PG_USER")
	pass   := os.Getenv("PG_PASS")
	dbUrl  := os.Getenv("PG_URL")
	dbName := os.Getenv("PG_DB")

	if user == "" || pass == "" || dbUrl == "" || dbName == "" {
		return ""
	}

	return "postgres://" + user + ":" + pass + "@" + dbUrl + "/" + dbName + "?sslmode=disable"
}
