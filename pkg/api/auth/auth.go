package auth

import (
	"net/http"

	"github.com/Soapstone-Services/go-template-2024"
	"github.com/labstack/echo/v4"
)

// Custom errors
var (
	ErrInvalidCredentials = echo.NewHTTPError(http.StatusUnauthorized, "Username or password does not exist")
)

// Authenticate tries to authenticate the user provided by username and password
func (a Auth) Authenticate(c echo.Context, user, pass string) (template.AuthToken, error) {
	u, err := a.udb.FindByUsername(a.db, user)
	if err != nil {
		return template.AuthToken{}, err
	}

	if !a.sec.HashMatchesPassword(u.Password, pass) {
		return template.AuthToken{}, ErrInvalidCredentials
	}

	if !u.Active {
		return template.AuthToken{}, template.ErrUnauthorized
	}

	token, err := a.tg.GenerateToken(u)
	if err != nil {
		return template.AuthToken{}, template.ErrUnauthorized
	}

	u.UpdateLastLogin(a.sec.Token(token))

	if err := a.udb.Update(a.db, u); err != nil {
		return template.AuthToken{}, err
	}

	return template.AuthToken{Token: token, RefreshToken: u.Token}, nil
}

// Refresh refreshes jwt token and puts new claims inside
func (a Auth) Refresh(c echo.Context, refreshToken string) (string, error) {
	user, err := a.udb.FindByToken(a.db, refreshToken)
	if err != nil {
		return "", err
	}
	return a.tg.GenerateToken(user)
}

// Me returns info about currently logged user
func (a Auth) Me(c echo.Context) (template.User, error) {
	au := a.rbac.User(c)
	return a.udb.View(a.db, au.ID)
}
