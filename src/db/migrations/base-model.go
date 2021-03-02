package migration

import (
	"time"

	"github.com/google/uuid"
)

// BaseModel define
type BaseModel struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
