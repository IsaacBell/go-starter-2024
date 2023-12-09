package user

import (
	"github.com/go-pg/pg/v9"
	"github.com/go-pg/pg/v9/orm"
	"github.com/labstack/echo/v4"

	"github.com/Soapstone-Services/go-template-2024"
	"github.com/Soapstone-Services/go-template-2024/pkg/api/user/platform/pgsql"
)

// Service represents user application interface
type Service interface {
	Create(echo.Context, template.User) (template.User, error)
	List(echo.Context, template.Pagination) ([]template.User, error)
	View(echo.Context, int) (template.User, error)
	Delete(echo.Context, int) error
	Update(echo.Context, Update) (template.User, error)
}

// New creates new user application service
func New(db *pg.DB, udb UDB, rbac RBAC, sec Securer) *User {
	return &User{db: db, udb: udb, rbac: rbac, sec: sec}
}

// Initialize initalizes User application service with defaults
func Initialize(db *pg.DB, rbac RBAC, sec Securer) *User {
	return New(db, pgsql.User{}, rbac, sec)
}

// User represents user application service
type User struct {
	db   *pg.DB
	udb  UDB
	rbac RBAC
	sec  Securer
}

// Securer represents security interface
type Securer interface {
	Hash(string) string
}

// UDB represents user repository interface
type UDB interface {
	Create(orm.DB, template.User) (template.User, error)
	View(orm.DB, int) (template.User, error)
	List(orm.DB, *template.ListQuery, template.Pagination) ([]template.User, error)
	Update(orm.DB, template.User) error
	Delete(orm.DB, template.User) error
}

// RBAC represents role-based-access-control interface
type RBAC interface {
	User(echo.Context) template.AuthUser
	EnforceUser(echo.Context, int) error
	AccountCreate(echo.Context, template.AccessRole, int, int) error
	IsLowerRole(echo.Context, template.AccessRole) error
}
