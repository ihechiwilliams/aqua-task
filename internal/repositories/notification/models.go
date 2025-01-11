package notification

import (
	"github.com/google/uuid"
	"time"
)

type Notification struct {
	ID        uuid.UUID `gorm:"primaryKey"`
	UserID    string    `gorm:"index"`
	Message   string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}
