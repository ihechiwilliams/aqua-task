package notification

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type TestSuiteSQLRepository struct {
	suite.Suite

	ctx    context.Context
	db     *gorm.DB
	mockDB sqlmock.Sqlmock
	repo   *SQLRepository
}

func (s *TestSuiteSQLRepository) SetupTest() {
	s.ctx = context.Background()
	db, mockDB, _ := sqlmock.New()
	s.mockDB = mockDB

	dialector := postgres.New(postgres.Config{Conn: db})
	gormDB, _ := gorm.Open(dialector, &gorm.Config{})

	s.db = gormDB

	s.repo = NewSQLRepository(s.db)
}

func (s *TestSuiteSQLRepository) TearDownSubTest() {
	s.Require().NoError(s.mockDB.ExpectationsWereMet())
}

func TestRepositorySuite(t *testing.T) {
	suite.Run(t, new(TestSuiteSQLRepository))
}

func (s *TestSuiteSQLRepository) TestInsertNotification() {
	notificationParams := &Notification{
		UserID:  "123456",
		Message: "Test Message",
	}

	s.Run("insert new notification if no conflicting record exists", func() {
		// Mock the database expectations
		s.mockDB.ExpectBegin()
		s.mockDB.ExpectExec(`INSERT INTO "notifications" \("id","user_id","message","created_at"\) VALUES \(\$1,\$2,\$3,\$4\)`).
			WithArgs(sqlmock.AnyArg(), notificationParams.UserID, notificationParams.Message, sqlmock.AnyArg()).
			WillReturnResult(sqlmock.NewResult(1, 1))
		s.mockDB.ExpectCommit()

		// Call the repository function
		err := s.repo.InsertNotification(s.ctx, notificationParams.UserID, notificationParams.Message)

		// Assertions
		s.Require().NoError(err)
	})

	// Failure scenario: Database returns an error
	s.Run("return error if insert fails due to database error", func() {
		// Mock the database expectations
		s.mockDB.ExpectBegin()
		s.mockDB.ExpectExec(`INSERT INTO "notifications" \("id","user_id","message","created_at"\) VALUES \(\$1,\$2,\$3,\$4\)`).
			WithArgs(sqlmock.AnyArg(), notificationParams.UserID, notificationParams.Message, sqlmock.AnyArg()).
			WillReturnError(fmt.Errorf("database error"))
		s.mockDB.ExpectRollback()

		// Call the repository function
		err := s.repo.InsertNotification(s.ctx, notificationParams.UserID, notificationParams.Message)

		// Assertions
		s.Require().Error(err)
		s.EqualError(err, "database error")
	})
}

func (s *TestSuiteSQLRepository) TestGetNotificationsByUserID() {
	userID := "123456"
	notifications := []*Notification{
		{
			ID:        uuid.New(),
			UserID:    userID,
			Message:   "Notification 1",
			CreatedAt: time.Now(),
		},
		{
			ID:        uuid.New(),
			UserID:    userID,
			Message:   "Notification 2",
			CreatedAt: time.Now(),
		},
	}

	s.Run("return notifications when records exist for user", func() {
		rows := sqlmock.NewRows([]string{"id", "user_id", "message", "created_at"}).
			AddRow(notifications[0].ID, notifications[0].UserID, notifications[0].Message, notifications[0].CreatedAt).
			AddRow(notifications[1].ID, notifications[1].UserID, notifications[1].Message, notifications[1].CreatedAt)

		s.mockDB.ExpectQuery(`SELECT \* FROM "notifications" WHERE user_id = \$1`).
			WithArgs(userID).
			WillReturnRows(rows)

		result, err := s.repo.GetNotificationsByUserID(s.ctx, userID)

		s.Require().NoError(err)
		s.Require().Len(result, len(notifications))
		s.Equal(notifications[0].ID, result[0].ID)
		s.Equal(notifications[1].Message, result[1].Message)
	})

	s.Run("return empty slice when no notifications exist for user", func() {
		rows := sqlmock.NewRows([]string{"id", "user_id", "message", "created_at"}) // No rows

		s.mockDB.ExpectQuery(`SELECT \* FROM "notifications" WHERE user_id = \$1`).
			WithArgs(userID).
			WillReturnRows(rows)

		result, err := s.repo.GetNotificationsByUserID(s.ctx, userID)

		s.Require().NoError(err)
		s.Empty(result)
	})

	s.Run("return error when database query fails", func() {
		s.mockDB.ExpectQuery(`SELECT \* FROM "notifications" WHERE user_id = \$1`).
			WithArgs(userID).
			WillReturnError(fmt.Errorf("database error"))

		result, err := s.repo.GetNotificationsByUserID(s.ctx, userID)

		s.Require().Error(err)
		s.Nil(result)
		s.EqualError(err, "database error")
	})
}

func (s *TestSuiteSQLRepository) TestDeleteNotificationByID() {
	notificationID := "123"

	s.Run("successfully delete notification when record exists", func() {
		s.mockDB.ExpectBegin()
		s.mockDB.ExpectExec(`DELETE FROM "notifications" WHERE id = \$1`).
			WithArgs(notificationID).
			WillReturnResult(sqlmock.NewResult(1, 1)) // 1 row affected
		s.mockDB.ExpectCommit()

		err := s.repo.DeleteNotificationByID(s.ctx, notificationID)

		s.Require().NoError(err)
	})

	s.Run("return error when no notification found with given ID", func() {
		s.mockDB.ExpectBegin()
		s.mockDB.ExpectExec(`DELETE FROM "notifications" WHERE id = \$1`).
			WithArgs(notificationID).
			WillReturnResult(sqlmock.NewResult(0, 0)) // No rows affected
		s.mockDB.ExpectCommit()

		err := s.repo.DeleteNotificationByID(s.ctx, notificationID)

		s.Require().Error(err)
		s.EqualError(err, "no resource found with the given ID")
	})

	s.Run("return error when database query fails", func() {
		s.mockDB.ExpectBegin()
		s.mockDB.ExpectExec(`DELETE FROM "notifications" WHERE id = \$1`).
			WithArgs(notificationID).
			WillReturnError(fmt.Errorf("database error"))
		s.mockDB.ExpectRollback()

		err := s.repo.DeleteNotificationByID(s.ctx, notificationID)

		s.Require().Error(err)
		s.EqualError(err, "database error")
	})
}

func (s *TestSuiteSQLRepository) TestDeleteAllNotificationsByUserID() {
	userID := "123"

	s.Run("successfully delete notifications for the user", func() {
		s.mockDB.ExpectBegin()
		s.mockDB.ExpectExec(`DELETE FROM "notifications" WHERE user_id = \$1`).
			WithArgs(userID).
			WillReturnResult(sqlmock.NewResult(0, 3)) // 3 rows affected
		s.mockDB.ExpectCommit()

		err := s.repo.DeleteAllNotificationsByUserID(s.ctx, userID)

		s.Require().NoError(err)
	})

	s.Run("no notifications found for the user", func() {
		s.mockDB.ExpectBegin()
		s.mockDB.ExpectExec(`DELETE FROM "notifications" WHERE user_id = \$1`).
			WithArgs(userID).
			WillReturnResult(sqlmock.NewResult(0, 0)) // 0 rows affected
		s.mockDB.ExpectCommit()

		err := s.repo.DeleteAllNotificationsByUserID(s.ctx, userID)

		s.Require().NoError(err)
	})

	s.Run("return error when database query fails", func() {
		s.mockDB.ExpectBegin()
		s.mockDB.ExpectExec(`DELETE FROM "notifications" WHERE user_id = \$1`).
			WithArgs(userID).
			WillReturnError(fmt.Errorf("database error"))
		s.mockDB.ExpectRollback()

		err := s.repo.DeleteAllNotificationsByUserID(s.ctx, userID)

		s.Require().Error(err)
		s.EqualError(err, "database error")
	})
}
