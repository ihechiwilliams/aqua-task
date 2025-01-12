package notification

import (
	"time"

	"github.com/google/uuid"
)

type Notification struct {
	ID        uuid.UUID `gorm:"primaryKey"`
	UserID    string    `gorm:"index"`
	Message   string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}
