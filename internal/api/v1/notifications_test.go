package v1_test

import (
	"errors"
	"net/http"
	"testing"
	"time"

	"aqua-backend/internal/api"
	v1 "aqua-backend/internal/api/v1"
	"aqua-backend/internal/repositories/notification"
	notificationsMock "aqua-backend/internal/repositories/notification/mocks"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/steinfletcher/apitest"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type notificationsSuite struct {
	suite.Suite
	mux               *gin.Engine
	notificationsRepo *notificationsMock.MockRepository
}

func (n *notificationsSuite) SetupSubTest() {
	n.mux = gin.Default()
	n.notificationsRepo = notificationsMock.NewMockRepository(n.T())

	notificationHandler := v1.NewNotificationHandler(n.notificationsRepo)
	apiService := v1.NewAPI(nil, nil, notificationHandler)
	api.InitRoutes(n.mux, api.NewRoutes(apiService))
}

func TestNotifications(t *testing.T) {
	suite.Run(t, new(notificationsSuite))
}

func (n *notificationsSuite) Test_V1DeleteNotification() {
	path := "/v1/notification/388fa464-590b-4c4d-908b-d045e06bbe81"

	n.Run("when request is successful", func() {
		n.notificationsRepo.On("DeleteNotificationByID",
			mock.Anything, mock.Anything,
		).Return(nil).Once()

		apitest.New().
			Handler(n.mux).
			Delete(path).
			Header("Content-Type", "application/json").
			Expect(n.T()).
			Status(http.StatusNoContent).
			End()
	})

	n.Run("when db fails", func() {
		n.notificationsRepo.On("DeleteNotificationByID",
			mock.Anything, mock.Anything,
		).Return(errors.New("")).Once()

		apitest.New().
			Handler(n.mux).
			Delete(path).
			Header("Content-Type", "application/json").
			Expect(n.T()).
			Status(http.StatusUnprocessableEntity).
			End()
	})
}

func (n *notificationsSuite) Test_V1GetUserNotifications() {
	path := "/v1/notifications/388fa464-590b-4c4d-908b-d045e06bbe89"

	n.Run("when request is successful", func() {
		newNotifications := []*notification.Notification{
			{
				ID:        uuid.New(),
				UserID:    uuid.New().String(),
				Message:   "Test message",
				CreatedAt: time.Now().UTC(),
			},
			{
				ID:        uuid.New(),
				UserID:    uuid.New().String(),
				Message:   "Test message new",
				CreatedAt: time.Now().UTC(),
			},
		}

		n.notificationsRepo.On("GetNotificationsByUserID",
			mock.Anything, mock.Anything,
		).Return(newNotifications, nil).Once()

		apitest.New().
			Handler(n.mux).
			Get(path).
			Header("Content-Type", "application/json").
			Expect(n.T()).
			Status(http.StatusOK).
			End()
	})

	n.Run("when db fails", func() {
		n.notificationsRepo.On("GetNotificationsByUserID",
			mock.Anything, mock.Anything,
		).Return([]*notification.Notification{}, errors.New("")).Once()

		apitest.New().
			Handler(n.mux).
			Get(path).
			Header("Content-Type", "application/json").
			Expect(n.T()).
			Status(http.StatusUnprocessableEntity).
			End()
	})
}

func (n *notificationsSuite) Test_V1DeleteUserNotifications() {
	path := "/v1/notifications/388fa464-590b-4c4d-908b-d045e06bbe81"

	n.Run("when request is successful", func() {
		n.notificationsRepo.On("DeleteAllNotificationsByUserID",
			mock.Anything, mock.Anything,
		).Return(nil).Once()

		apitest.New().
			Handler(n.mux).
			Delete(path).
			Header("Content-Type", "application/json").
			Expect(n.T()).
			Status(http.StatusNoContent).
			End()
	})

	n.Run("when db fails", func() {
		n.notificationsRepo.On("DeleteAllNotificationsByUserID",
			mock.Anything, mock.Anything,
		).Return(errors.New("")).Once()

		apitest.New().
			Handler(n.mux).
			Delete(path).
			Header("Content-Type", "application/json").
			Expect(n.T()).
			Status(http.StatusUnprocessableEntity).
			End()
	})
}
