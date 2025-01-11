package notificationconsumer

import (
	"aqua-backend/internal/repositories/notification"
	"context"
	"log"

	"aqua-backend/pkg/rabbitmq"
	//"aqua-task/internal/repositories/notifications"
	//"aqua-task/pkg/rabbitmq"
)

func StartNotificationConsumer(rmq *rabbitmq.RabbitMQ, queue string, repo *notification.SQLRepository) {
	msgs, err := rmq.Channel.Consume(queue, "", true, false, false, false, nil)
	if err != nil {
		log.Fatalf("Failed to consume messages: %v", err)
	}
	for msg := range msgs {
		userID := msg.Headers["user_id"].(string)
		message := string(msg.Body)

		// Insert into DB
		if err := repo.InsertNotification(context.Background(), userID, message); err != nil {
			log.Printf("Failed to insert notification: %v", err)
		} else {
			log.Printf("Notification saved for user %s: %s", userID, message)
		}
	}
}
