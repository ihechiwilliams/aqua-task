package v1_test

import (
	"context"
	"testing"
	"time"

	v1 "aqua-backend/internal/api/v1"
	"aqua-backend/internal/repositories/notification"
	notificationsMock "aqua-backend/internal/repositories/notification/mocks"
	pb "aqua-backend/proto"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type notificationGrpcSuite struct {
	suite.Suite
	notificationsRepo *notificationsMock.MockRepository
}

func (n *notificationGrpcSuite) SetupSubTest() {
	n.notificationsRepo = notificationsMock.NewMockRepository(n.T())
}

func TestNotificationsGrpc(t *testing.T) {
	suite.Run(t, new(notificationGrpcSuite))
}

func (n *notificationGrpcSuite) Test_GetNotifications() {
	n.Run("when request is successful", func() {
		server := v1.NotificationServer{Repo: n.notificationsRepo}

		ctx := context.Background()
		userID := "test-user"
		notifications := []*notification.Notification{
			{ID: uuid.New(), Message: "Test message 1", CreatedAt: time.Now()},
			{ID: uuid.New(), Message: "Test message 2", CreatedAt: time.Now()},
		}

		n.notificationsRepo.On("GetNotificationsByUserID", ctx, userID).Return(notifications, nil)

		req := &pb.GetNotificationsRequest{UserId: userID}
		resp, err := server.GetNotifications(ctx, req)

		assert.NoError(n.T(), err)
		assert.NotNil(n.T(), resp)
		assert.Len(n.T(), resp.Notifications, len(notifications))
		n.notificationsRepo.AssertExpectations(n.T())
	})
}

func (n *notificationGrpcSuite) Test_ClearNotification() {
	n.Run("when request is successful", func() {
		server := v1.NotificationServer{Repo: n.notificationsRepo}

		ctx := context.Background()
		notificationID := "notification-1"

		n.notificationsRepo.On("DeleteNotificationByID", ctx, notificationID).Return(nil)

		req := &pb.ClearNotificationRequest{NotificationId: notificationID}
		resp, err := server.ClearNotification(ctx, req)

		assert.NoError(n.T(), err)
		assert.NotNil(n.T(), resp)
		assert.Equal(n.T(), "Notification cleared", resp.Message)
		n.notificationsRepo.AssertExpectations(n.T())
	})
}

func (n *notificationGrpcSuite) Test_ClearAllNotifications() {
	n.Run("when request is successful", func() {
		server := v1.NotificationServer{Repo: n.notificationsRepo}

		ctx := context.Background()
		userID := "test-user"

		n.notificationsRepo.On("DeleteAllNotificationsByUserID", ctx, userID).Return(nil)

		req := &pb.ClearAllNotificationsRequest{UserId: userID}
		resp, err := server.ClearAllNotifications(ctx, req)

		assert.NoError(n.T(), err)
		assert.NotNil(n.T(), resp)
		assert.Equal(n.T(), "All notifications cleared", resp.Message)
		n.notificationsRepo.AssertExpectations(n.T())
	})
}
