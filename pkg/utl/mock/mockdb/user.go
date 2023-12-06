package mockdb

import (
	"github.com/go-pg/pg/v9/orm"
)

// User database mock
type User struct {
	CreateFn         func(orm.DB, stems.User) (stems.User, error)
	ViewFn           func(orm.DB, int) (stems.User, error)
	FindByUsernameFn func(orm.DB, string) (stems.User, error)
	FindByTokenFn    func(orm.DB, string) (stems.User, error)
	ListFn           func(orm.DB, *stems.ListQuery, stems.Pagination) ([]stems.User, error)
	DeleteFn         func(orm.DB, stems.User) error
	UpdateFn         func(orm.DB, stems.User) error
}

// Create mock
func (u *User) Create(db orm.DB, usr stems.User) (stems.User, error) {
	return u.CreateFn(db, usr)
}

// View mock
func (u *User) View(db orm.DB, id int) (stems.User, error) {
	return u.ViewFn(db, id)
}

// FindByUsername mock
func (u *User) FindByUsername(db orm.DB, uname string) (stems.User, error) {
	return u.FindByUsernameFn(db, uname)
}

// FindByToken mock
func (u *User) FindByToken(db orm.DB, token string) (stems.User, error) {
	return u.FindByTokenFn(db, token)
}

// List mock
func (u *User) List(db orm.DB, lq *stems.ListQuery, p stems.Pagination) ([]stems.User, error) {
	return u.ListFn(db, lq, p)
}

// Delete mock
func (u *User) Delete(db orm.DB, usr stems.User) error {
	return u.DeleteFn(db, usr)
}

// Update mock
func (u *User) Update(db orm.DB, usr stems.User) error {
	return u.UpdateFn(db, usr)
}
