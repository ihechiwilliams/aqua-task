package notification

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

const (
	tableName = "notifications"
)

type Repository interface {
	InsertNotification(ctx context.Context, userId, message string) error
	GetNotificationsByUserID(ctx context.Context, userId string) ([]*Notification, error)
	DeleteNotificationByID(ctx context.Context, Id string) error
	DeleteAllNotificationsByUserID(ctx context.Context, userId string) error
}

type SQLRepository struct {
	db *gorm.DB
}

func (r *SQLRepository) GetNotificationsByUserID(ctx context.Context, userId string) ([]*Notification, error) {
	var notifications []*Notification

	result := r.db.WithContext(ctx).
		Table(tableName).
		Where("user_id = ?", userId).
		Find(&notifications)

	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return []*Notification{}, nil
	}

	return notifications, nil
}

func (r *SQLRepository) DeleteNotificationByID(ctx context.Context, Id string) error {
	result := r.db.WithContext(ctx).
		Table(tableName).
		Where("id = ?", Id).Delete(&Notification{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("no resource found with the given ID")
	}
	return nil
}

func (r *SQLRepository) DeleteAllNotificationsByUserID(ctx context.Context, userId string) error {
	result := r.db.WithContext(ctx).
		Table(tableName).
		Where("user_id = ?", userId).Delete(&Notification{})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return nil
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
	return r.db.WithContext(ctx).Create(notification).Error
}

func NewSQLRepository(db *gorm.DB) *SQLRepository {
	return &SQLRepository{
		db: db,
	}
}
