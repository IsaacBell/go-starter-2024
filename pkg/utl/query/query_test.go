package query_test

import (
	"testing"

	"github.com/labstack/echo/v4"

	"github.com/stretchr/testify/assert"

	"github.com/Soapstone-Services/go-template-2024"
	"github.com/Soapstone-Services/go-template-2024/pkg/utl/query"
)

func TestList(t *testing.T) {
	type args struct {
		user template.AuthUser
	}
	cases := []struct {
		name     string
		args     args
		wantData *template.ListQuery
		wantErr  error
	}{
		{
			name: "Super admin user",
			args: args{user: template.AuthUser{
				Role: template.SuperAdminRole,
			}},
		},
		{
			name: "Company admin user",
			args: args{user: template.AuthUser{
				Role:      template.CompanyAdminRole,
				CompanyID: 1,
			}},
			wantData: &template.ListQuery{
				Query: "company_id = ?",
				ID:    1},
		},
		{
			name: "Location admin user",
			args: args{user: template.AuthUser{
				Role:       template.LocationAdminRole,
				CompanyID:  1,
				LocationID: 2,
			}},
			wantData: &template.ListQuery{
				Query: "location_id = ?",
				ID:    2},
		},
		{
			name: "Normal user",
			args: args{user: template.AuthUser{
				Role: template.UserRole,
			}},
			wantErr: echo.ErrForbidden,
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			q, err := query.List(tt.args.user)
			assert.Equal(t, tt.wantData, q)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
