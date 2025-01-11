package customers

import (
	"context"
	"regexp"
	"testing"
	"time"

	"aqua-backend/internal/constants"

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

func (s *TestSuiteSQLRepository) TestCreateCustomer() {
	createCustomerQuery := `INSERT INTO "customers"`

	customerParams := &DBCustomer{
		Name:  "John Doe",
		Email: "johndoe@test.com",
	}

	customerResult := sqlmock.NewRows(
		[]string{
			"id", "name", "email", "created_at",
			"updated_at",
		},
	).AddRow(
		"3be5d416-48e2-4e55-b18b-5874a1fd9919",
		customerParams.Name,
		customerParams.Email,
		time.Now().UTC(),
		time.Now().UTC(),
	)

	s.Run("do nothing if a conflicting customer exists", func() {
		s.mockDB.ExpectBegin()
		s.mockDB.ExpectQuery(createCustomerQuery).
			WithArgs().WillReturnRows(sqlmock.NewRows([]string{})) // Simulate DO NOTHING
		s.mockDB.ExpectCommit()

		charge, err := s.repo.CreateCustomer(s.ctx, customerParams)

		s.Require().ErrorIs(err, constants.ErrRecordAlreadyExists)
		s.Require().Nil(charge)
	})

	s.Run("insert new customer if no conflicting record exists", func() {
		s.mockDB.ExpectBegin()
		s.mockDB.ExpectQuery(createCustomerQuery).
			WithArgs().WillReturnRows(customerResult)
		s.mockDB.ExpectCommit()

		customer, err := s.repo.CreateCustomer(s.ctx, customerParams)

		s.Require().NoError(err)
		s.Require().NotNil(customer)
		s.Require().Equal(customerParams.Name, customer.Name)
		s.Require().Equal(customerParams.Email, customer.Email)
	})
}

func (s *TestSuiteSQLRepository) TestGetCustomerByID() {
	s.Run("when the record exists", func() {
		id := uuid.New()

		query := regexp.QuoteMeta(`SELECT * FROM "customers" WHERE id = $1`)
		rows := sqlmock.NewRows([]string{"id"}).AddRow(id)

		s.mockDB.ExpectQuery(query).WithArgs(id.String(), 1).WillReturnRows(rows)

		account, err := s.repo.GetCustomerByID(s.ctx, id)

		s.Require().NoError(err)
		s.Equal(account.ID, id)
	})

	s.Run("when the record does not exist", func() {
		id := uuid.New()
		query := regexp.QuoteMeta(`SELECT * FROM "customers" WHERE id = $1`)
		rows := sqlmock.NewRows([]string{"id"})

		s.mockDB.ExpectQuery(query).WithArgs(id.String(), 1).WillReturnRows(rows)

		account, err := s.repo.GetCustomerByID(s.ctx, id)

		s.Nil(err)
		s.Nil(account)
	})
}
