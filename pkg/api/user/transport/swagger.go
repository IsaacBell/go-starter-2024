package transport

import stems "github.com/Soapstone-Services/go-template-2024"

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
