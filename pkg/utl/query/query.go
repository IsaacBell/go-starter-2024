package query

import (
	"github.com/Soapstone-Services/go-template-2024"
	"github.com/labstack/echo/v4"
)

// List prepares data for list queries
func List(u template.AuthUser) (*template.ListQuery, error) {
	switch true {
	case u.Role <= template.AdminRole: // user is SuperAdmin or Admin
		return nil, nil
	case u.Role == template.CompanyAdminRole:
		return &template.ListQuery{Query: "company_id = ?", ID: u.CompanyID}, nil
	case u.Role == template.LocationAdminRole:
		return &template.ListQuery{Query: "location_id = ?", ID: u.LocationID}, nil
	default:
		return nil, echo.ErrForbidden
	}
}
