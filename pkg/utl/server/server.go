package server

import (
	"context"
	"crypto/subtle"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-playground/validator"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"

	// "github.com/Soapstone-Services/go-template-2024/pkg/utl/middleware/secure"

	echoPrometheus "github.com/globocom/echo-prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// instantiate new Echo server
func New() *echo.Echo {
	e := echo.New()

	if !isProductionEnv() {
		defer logServerRoutes(e)
	}

	e.Use(echoPrometheus.MetricsMiddleware())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			username, password, ok := c.Request().BasicAuth()
			if ok {
				// TODO - remove after testing
				fmt.Println("u: ", username)
				fmt.Println("p: ", password)
			} else {
				return echo.NewHTTPError(
					http.StatusUnauthorized,
					"Please provide valid credentials",
				)
			}

			authUsername := "changeme"
			val, ok := os.LookupEnv("V1_AUTH_USER")
			fmt.Println("au: ", authUsername)
			if !ok {
				fmt.Printf("HTTP Basic Auth username not set!\n")
			} else {
				authUsername = val
			}

			authPassword := "changeme"
			val, ok = os.LookupEnv("V1_AUTH_PASS")
			if !ok {
				fmt.Printf("HTTP Basic Auth password not set!\n")
			} else {
				authPassword = val
			}

			if subtle.ConstantTimeCompare([]byte(username), []byte(authUsername)) == 1 &&
				subtle.ConstantTimeCompare([]byte(password), []byte(authPassword)) == 1 {
				return next(c)
			}

			return echo.NewHTTPError(http.StatusUnauthorized, "Please provide valid credentials")
		}
	})

	v1 := e.Group("/v1")
	v1.POST("/upload", upload)
	v1.GET("/", healthCheck)
	v1.GET("/health", healthCheck)

	// expose metrics for analytics collection
	// see Prometheus: https://prometheus.io/
	e.GET("/metrics", prometheusHandlerFunc())

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.GET("/", healthCheck)

	e.Validator = &CustomValidator{V: validator.New()}
	custErr := &customErrHandler{e: e}
	e.HTTPErrorHandler = custErr.handler
	e.Binder = &CustomBinder{b: &echo.DefaultBinder{}}
	return e
}

func logServerRoutes(e *echo.Echo) error {
	data, err := json.MarshalIndent(e.Routes(), "", "  ")
	if err != nil {
		return err
	}
	os.WriteFile("local/_routes.json", data, 0644)

	return nil
}

func prometheusHandlerFunc() echo.HandlerFunc {
	return echo.WrapHandler(promhttp.Handler())
}

func healthCheck(c echo.Context) error {
	if c.Request().Header.Get("Content-Type") == "application/json" {
		return c.JSON(http.StatusOK, "OK")
	}
	return c.String(http.StatusOK, "OK")
}

func upload(c echo.Context) error {
	// Read form fields
	name := c.FormValue("name")
	email := c.FormValue("description")

	//-----------
	// Read file
	//-----------

	// Source
	file, err := c.FormFile("file")
	if err != nil {
		return err
	}
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// Destination
	dst, err := os.Create(file.Filename)
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	return c.HTML(http.StatusOK, fmt.Sprintf("<p>File %s uploaded successfully with fields name=%s and email=%s.</p>", file.Filename, name, email))
}

func isProductionEnv() bool {
	return os.Getenv("ENVIRONMENT_NAME") == "production"
}

// config for the current machine
type Config struct {
	Port                string
	ReadTimeoutSeconds  int
	WriteTimeoutSeconds int
	Debug               bool
}

// Spin up an echo server
func Start(e *echo.Echo, cfg *Config) {
	s := &http.Server{
		Addr:         cfg.Port,
		ReadTimeout:  time.Duration(cfg.ReadTimeoutSeconds) * time.Second,
		WriteTimeout: time.Duration(cfg.WriteTimeoutSeconds) * time.Second,
	}
	e.Debug = cfg.Debug

	// Start the server runtime
	go func() {
		err := e.StartServer(s)
		if err != nil {
			e.Logger.Info("::[Server] Server crash on startup.::")
			e.Logger.Error(err)
		} else {
			e.Logger.Info("::Server successfully booted.::")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 10 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)

	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
