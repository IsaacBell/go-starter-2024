package transport

import stems "github.com/Soapstone-Services/template-01"

// User model response
// swagger:response userResp
type swaggUserResponse struct {
	// in:body
	Body struct {
		*stems.User
	}
}

// Users model response
// swagger:response userListResp
type swaggUserListResponse struct {
	// in:body
	Body struct {
		Users []stems.User `json:"users"`
		Page  int          `json:"page"`
	}
}
