package domain

import (
	"strings"

	"github.com/google/uuid"
)

// Client - Client Entity
type User struct {
	UserId    uuid.UUID `json:"userId" param:"userId"`
	UserName  string    `json:"userName" form:"userName" query:"userName" validate:"required"`
	Email     string    `json:"email" form:"email" query:"email" `
	CompanyID int       `json:"companyId" form:"companyId" query:"companyId"`
}

func (u *User) get() interface{} {
	return u
}

// IsValid - Check User
func (u *User) IsValid() bool {
	return (len(strings.TrimSpace(u.UserName)) > 0)
}
