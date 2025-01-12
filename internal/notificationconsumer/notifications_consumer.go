package notificationconsumer

import (
	"aqua-backend/internal/repositories/notification"
	"context"
	"encoding/json"
	"log"

	"aqua-backend/pkg/rabbitmq"
	//"aqua-task/internal/repositories/notifications"
	//"aqua-task/pkg/rabbitmq"
)

type NotificationMessage struct {
	UserID  string `json:"user_id"`
	Message string `json:"message"`
}

func StartNotificationConsumer(rmq *rabbitmq.RabbitMQ, queue string, repo *notification.SQLRepository) {
	msgs, err := rmq.Channel.Consume(queue, "", true, false, false, false, nil)
	if err != nil {
		log.Fatalf("Failed to consume messages: %v", err)
	}
	for msg := range msgs {
		log.Println("Received message:", string(msg.Body))

		// Parse the JSON message
		var notificationMessage NotificationMessage
		if err := json.Unmarshal(msg.Body, &notificationMessage); err != nil {
			log.Printf("Failed to parse message: %v", err)
			continue
		}

		userID := notificationMessage.UserID
		message := notificationMessage.Message

		// Insert into DB
		if err := repo.InsertNotification(context.Background(), userID, message); err != nil {
			log.Printf("Failed to insert notification: %v", err)
		} else {
			log.Printf("Notification saved for user %s: %s", userID, message)
		}
	}
}
