package auth

import (
	"time"

	"github.com/labstack/echo/v4"

	stems "github.com/Soapstone-Services/go-template-2024"
	"github.com/Soapstone-Services/go-template-2024/pkg/api/auth"
)

// New creates new auth logging service
func New(svc auth.Service, logger echo.Logger) *LogService {
	return &LogService{
		Service: svc,
		logger:  logger,
	}
}

// LogService represents auth logging service
type LogService struct {
	auth.Service
	logger echo.Logger
}

const name = "auth"

// Authenticate logging
func (ls *LogService) Authenticate(c echo.Context, user, password string) (resp stems.AuthToken, err error) {
	defer func(begin time.Time) {
		ls.logger.Info(
			c,
			name, "Authenticate request", err,
			map[string]interface{}{
				"req":  user,
				"took": time.Since(begin),
			},
		)
	}(time.Now())
	return ls.Service.Authenticate(c, user, password)
}

// Refresh logging
func (ls *LogService) Refresh(c echo.Context, req string) (token string, err error) {
	defer func(begin time.Time) {
		ls.logger.Info(
			c,
			name, "Refresh request", err,
			map[string]interface{}{
				"req":  req,
				"resp": token,
				"took": time.Since(begin),
			},
		)
	}(time.Now())
	return ls.Service.Refresh(c, req)
}

// Me logging
func (ls *LogService) Me(c echo.Context) (resp stems.User, err error) {
	defer func(begin time.Time) {
		ls.logger.Info(
			c,
			name, "Me request", err,
			map[string]interface{}{
				"resp": resp,
				"took": time.Since(begin),
			},
		)
	}(time.Now())
	return ls.Service.Me(c)
}
