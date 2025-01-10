package resources

import (
	"github.com/google/uuid"
	"time"
)

type DBResource struct {
	ID         uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Name       string    `gorm:"type:varchar(255);not null;unique" json:"name"`
	Type       string    `gorm:"type:varchar(100);not null" json:"type"`
	Region     string    `gorm:"type:varchar(100);not null" json:"region"`
	CustomerID uuid.UUID `gorm:"type:uuid;not null;constraint:OnDelete:CASCADE;" json:"customer_id"`
	CreatedAt  time.Time `gorm:"type:timestamp;default:current_timestamp" json:"created_at"`
	UpdatedAt  time.Time `gorm:"type:timestamp;default:current_timestamp on update current_timestamp" json:"updated_at"`
}

type Resource struct {
	ID         uuid.UUID `json:"id"`
	Name       string    `json:"name"`
	Type       string    `json:"type"`
	Region     string    `json:"region"`
	CustomerID uuid.UUID `json:"customer_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
