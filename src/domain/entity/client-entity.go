package domain

import (
	"strings"
	"time"

	"github.com/google/uuid"
)

// A Client entity
type Client struct {
	ID        uuid.UUID `json:"id" param:"id"`
	Name      string    `json:"name" form:"name" query:"name" validate:"required"`
	Email     string    `json:"email" form:"email" query:"email" validate:"required"`
	Age       int       `json:"age" form:"age" query:"age" validate:""`
	CompanyID int       `json:"companyId" form:"companyId" query:"companyId" validate:"required"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

func (c Client) IsValid() bool {
	return (len(strings.TrimSpace(c.Name)) > 0)
}
