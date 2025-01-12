package v1

import (
	"context"

	"aqua-backend/internal/repositories/notification"
	pb "aqua-backend/proto"
)

type NotificationServer struct {
	Repo notification.Repository
	pb.UnimplementedNotificationServiceServer
}

func (s *NotificationServer) GetNotifications(ctx context.Context, req *pb.GetNotificationsRequest) (*pb.GetNotificationsResponse, error) {
	// Logic to fetch notifications from the database
	notifications, err := s.Repo.GetNotificationsByUserID(ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	grpcNotifications := make([]*pb.Notification, 0, len(notifications))

	for _, n := range notifications {
		grpcNotifications = append(grpcNotifications, &pb.Notification{
			Id:        n.ID.String(),
			Message:   n.Message,
			CreatedAt: n.CreatedAt.Format("2006-01-02T15:04:05Z"),
		})
	}

	return &pb.GetNotificationsResponse{Notifications: grpcNotifications}, nil
}

func (s *NotificationServer) ClearNotification(ctx context.Context, req *pb.ClearNotificationRequest) (*pb.ClearNotificationResponse, error) {
	err := s.Repo.DeleteNotificationByID(ctx, req.NotificationId)
	if err != nil {
		return nil, err
	}

	return &pb.ClearNotificationResponse{Message: "Notification cleared"}, nil
}

func (s *NotificationServer) ClearAllNotifications(ctx context.Context, req *pb.ClearAllNotificationsRequest) (*pb.ClearAllNotificationsResponse, error) {
	err := s.Repo.DeleteAllNotificationsByUserID(ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	return &pb.ClearAllNotificationsResponse{Message: "All notifications cleared"}, nil
}
