package v1

import (
	"net/http"

	"aqua-backend/internal/api/server"
	"aqua-backend/internal/repositories/notification"

	"github.com/gin-gonic/gin"
	openapi_types "github.com/oapi-codegen/runtime/types"
	"github.com/samber/lo"
)

type NotificationsHandler struct {
	notificationRepo notification.Repository
}

func NewNotificationHandler(notificationRepo notification.Repository) *NotificationsHandler {
	return &NotificationsHandler{
		notificationRepo: notificationRepo,
	}
}

func (a *API) V1DeleteNotification(c *gin.Context, notificationID openapi_types.UUID) {
	err := a.notificationHandler.notificationRepo.DeleteNotificationByID(c.Request.Context(), notificationID.String())
	if err != nil {
		server.ProcessingError(err, c)

		return
	}

	c.JSON(http.StatusNoContent, gin.H{"msg": "notifications deleted"})
}

func (a *API) V1GetUserNotifications(c *gin.Context, userID string) {
	result, err := a.notificationHandler.notificationRepo.GetNotificationsByUserID(c.Request.Context(), userID)
	if err != nil {
		server.ProcessingError(err, c)

		return
	}

	c.JSON(http.StatusOK, lo.Map(result, func(notif *notification.Notification, _ int) server.UserNotificationResponseData {
		return serializeNotificationToAPIResponseData(notif)
	}))
}

func (a *API) V1DeleteUserNotifications(c *gin.Context, userID string) {
	err := a.notificationHandler.notificationRepo.DeleteAllNotificationsByUserID(c.Request.Context(), userID)
	if err != nil {
		server.ProcessingError(err, c)

		return
	}

	c.JSON(http.StatusNoContent, gin.H{"msg": "notifications deleted"})
}

func serializeNotificationToAPIResponseData(notifs *notification.Notification) server.UserNotificationResponseData {
	return server.UserNotificationResponseData{
		Id:      notifs.ID,
		Message: notifs.Message,
		UserId:  notifs.UserID,
	}
}
