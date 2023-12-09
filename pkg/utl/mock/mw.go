package mock

import "github.com/Soapstone-Services/go-template-2024"

// JWT mock
type JWT struct {
	GenerateTokenFn func(template.User) (string, error)
}

// GenerateToken mock
func (j JWT) GenerateToken(u template.User) (string, error) {
	return j.GenerateTokenFn(u)
}
