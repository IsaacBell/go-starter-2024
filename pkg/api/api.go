package api

import (
	"fmt"
	"os"
	"strings"

	// "github.com/Soapstone-Services/go-template-2024/pkg/utl/zlog"
	// "github.com/Soapstone-Services/go-template-2024/pkg/api/auth"
	// al "github.com/Soapstone-Services/go-template-2024/pkg/api/auth/logging"
	// at "github.com/Soapstone-Services/go-template-2024/pkg/api/auth/transport"
	// "github.com/Soapstone-Services/go-template-2024/pkg/api/password"
	// pl "github.com/Soapstone-Services/go-template-2024/pkg/api/password/logging"
	// pt "github.com/Soapstone-Services/go-template-2024/pkg/api/password/transport"
	// "github.com/Soapstone-Services/go-template-2024/pkg/api/user"
	// ul "github.com/Soapstone-Services/go-template-2024/pkg/api/user/logging"
	// ut "github.com/Soapstone-Services/go-template-2024/pkg/api/user/transport"
	authMw "github.com/Soapstone-Services/go-template-2024/pkg/utl/middleware/auth"

	"github.com/Soapstone-Services/go-template-2024/pkg/utl/config"
	"github.com/Soapstone-Services/go-template-2024/pkg/utl/postgres"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/rs/zerolog"
	"github.com/ziflex/lecho/v3"

	"github.com/Soapstone-Services/go-template-2024/pkg/utl/jwt"
	"github.com/Soapstone-Services/go-template-2024/pkg/utl/server"
)

// Start starts the API service
func Start(cfg *config.Configuration) error {
	db, err := postgres.New(cfg.DB.Url, cfg.DB.Timeout, cfg.DB.LogQueries)
	if err != nil {
		return err
	}
	fmt.Println("DB Addr: ", db.Options().Addr)

	// sec := secure.New(cfg.App.MinPasswordStr, sha1.New())
	// rbac := rbac.Service{}

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
			"service": "template01",
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

	// at.NewHTTP(al.New(auth.Initialize(db, jwt, sec, rbac), log), e, authMiddleware)

	v1 := e.Group("/v1")
	v1.Use(authMiddleware)

	// ut.NewHTTP(ul.New(user.Initialize(db, rbac, sec), log), v1)
	// pt.NewHTTP(pl.New(password.Initialize(db, rbac, sec), log), v1)

	server.Start(e, &server.Config{
		Port:                cfg.Server.Port,
		ReadTimeoutSeconds:  cfg.Server.ReadTimeout,
		WriteTimeoutSeconds: cfg.Server.WriteTimeout,
		Debug:               cfg.Server.Debug,
	})

	return nil
}

func isProd() bool {
	return os.Getenv("ENVIRONMENT_NAME") == "production"
}
