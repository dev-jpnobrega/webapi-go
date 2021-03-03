package domain

import (
	"strings"
	"time"

	"github.com/google/uuid"
)

// Client - Client Entity
type Client struct {
	ID        uuid.UUID
	Name      string `json:"name" form:"name" query:"name" validate:"required"`
	Email     string `json:"email" form:"email" query:"email" validate:"required"`
	Age       int    `json:"age" form:"age" query:"age" validate:""`
	CompanyID int    `json:"companyId" form:"companyId" query:"companyId" validate:"required"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

// IsValid - Check Client
func (c Client) IsValid() bool {
	if len(strings.TrimSpace(c.Name)) > 0 {
		return true
	}
	return false
}
