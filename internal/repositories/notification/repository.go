package notification

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

const (
	tableName = "notifications"
)

type Repository interface {
	InsertNotification(ctx context.Context, userID, message string) error
	GetNotificationsByUserID(ctx context.Context, userID string) ([]*Notification, error)
	DeleteNotificationByID(ctx context.Context, id string) error
	DeleteAllNotificationsByUserID(ctx context.Context, userID string) error
}

type SQLRepository struct {
	db *gorm.DB
}

func (r *SQLRepository) GetNotificationsByUserID(ctx context.Context, userID string) ([]*Notification, error) {
	var notifications []*Notification

	result := r.db.WithContext(ctx).
		Table(tableName).
		Where("user_id = ?", userID).
		Find(&notifications)

	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return []*Notification{}, nil
	}

	return notifications, nil
}

func (r *SQLRepository) DeleteNotificationByID(ctx context.Context, id string) error {
	result := r.db.WithContext(ctx).
		Table(tableName).
		Where("id = ?", id).Delete(&Notification{})
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("no resource found with the given ID")
	}

	return nil
}

func (r *SQLRepository) DeleteAllNotificationsByUserID(ctx context.Context, userID string) error {
	result := r.db.WithContext(ctx).
		Table(tableName).
		Where("user_id = ?", userID).
		Delete(&Notification{})

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *SQLRepository) InsertNotification(ctx context.Context, userID, message string) error {
	notification := &Notification{
		ID:        uuid.New(),
		UserID:    userID,
		Message:   message,
		CreatedAt: time.Now(),
	}

	return r.db.WithContext(ctx).Table(tableName).Create(notification).Error
}

func NewSQLRepository(db *gorm.DB) *SQLRepository {
	return &SQLRepository{
		db: db,
	}
}
