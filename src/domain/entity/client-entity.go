package domain

import (
	"strings"
	"time"

	"github.com/google/uuid"
)

// Client - Client Entity
type Client struct {
	ID        uuid.UUID
	Name      string
	Age       int
	CompanyID int
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
