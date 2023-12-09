package mockdb

import (
	"github.com/Soapstone-Services/go-template-2024"
	"github.com/go-pg/pg/v9/orm"
)

// User database mock
type User struct {
	CreateFn         func(orm.DB, template.User) (template.User, error)
	ViewFn           func(orm.DB, int) (template.User, error)
	FindByUsernameFn func(orm.DB, string) (template.User, error)
	FindByTokenFn    func(orm.DB, string) (template.User, error)
	ListFn           func(orm.DB, *template.ListQuery, template.Pagination) ([]template.User, error)
	DeleteFn         func(orm.DB, template.User) error
	UpdateFn         func(orm.DB, template.User) error
}

// Create mock
func (u *User) Create(db orm.DB, usr template.User) (template.User, error) {
	return u.CreateFn(db, usr)
}

// View mock
func (u *User) View(db orm.DB, id int) (template.User, error) {
	return u.ViewFn(db, id)
}

// FindByUsername mock
func (u *User) FindByUsername(db orm.DB, uname string) (template.User, error) {
	return u.FindByUsernameFn(db, uname)
}

// FindByToken mock
func (u *User) FindByToken(db orm.DB, token string) (template.User, error) {
	return u.FindByTokenFn(db, token)
}

// List mock
func (u *User) List(db orm.DB, lq *template.ListQuery, p template.Pagination) ([]template.User, error) {
	return u.ListFn(db, lq, p)
}

// Delete mock
func (u *User) Delete(db orm.DB, usr template.User) error {
	return u.DeleteFn(db, usr)
}

// Update mock
func (u *User) Update(db orm.DB, usr template.User) error {
	return u.UpdateFn(db, usr)
}
