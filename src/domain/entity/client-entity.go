package domain

import "strings"

// Client - Client Entity
type Client struct {
	ID        string
	Name      string
	Age       int
	CompanyID int
}

// IsValid - Check Client
func (c Client) IsValid() bool {
	if len(strings.TrimSpace(c.Name)) > 0 {
		return true
	}
	return false
}
