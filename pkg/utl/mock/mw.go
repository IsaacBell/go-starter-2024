package mock

// JWT mock
type JWT struct {
	GenerateTokenFn func(stems.User) (string, error)
}

// GenerateToken mock
func (j JWT) GenerateToken(u stems.User) (string, error) {
	return j.GenerateTokenFn(u)
}
