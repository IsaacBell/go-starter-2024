package server

import (
	"context"
	"encoding/json"
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

	v1 := e.Group("/v1")

	/* TODO: uncomment and test */
	v1.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		// authUsername := "changeme"
		// val, ok := os.LookupEnv("V1_AUTH_USER")
		// fmt.Println("HERE")
		// if !ok {
		// 	fmt.Printf("HTTP Basic Auth username not set!\n")
		// } else {
		// 	authUsername = val
		// }

		// authPassword := "changeme"
		// val, ok = os.LookupEnv("V1_AUTH_PASS")
		// if !ok {
		// 	fmt.Printf("HTTP Basic Auth password not set!\n")
		// } else {
		// 	authPassword = val
		// }

		// if username == authUsername && password == authPassword {
		if username == "changeme" && password == "changeme" {
			return true, nil
		}

		return false, nil
	}))

	// expose metrics for analytics collection
	// see Prometheus: https://prometheus.io/
	e.GET("/metrics", prometheusHandlerFunc())

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.GET("/", healthCheck)
	e.GET("/v1", healthCheck)

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
