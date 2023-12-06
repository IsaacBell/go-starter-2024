package query

import (
	"github.com/labstack/echo/v4"

	stems "github.com/Soapstone-Services/go-template-2024"
)

// List prepares data for list queries
func List(u stems.AuthUser) (*stems.ListQuery, error) {
	switch true {
	case u.Role <= stems.AdminRole: // user is SuperAdmin or Admin
		return nil, nil
	case u.Role == stems.CompanyAdminRole:
		return &stems.ListQuery{Query: "company_id = ?", ID: u.CompanyID}, nil
	case u.Role == stems.LocationAdminRole:
		return &stems.ListQuery{Query: "location_id = ?", ID: u.LocationID}, nil
	default:
		return nil, echo.ErrForbidden
	}
}
