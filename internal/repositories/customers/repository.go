package customers

import (
	"context"
	"errors"

	"aqua-backend/internal/constants"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
)

const (
	tableName = "customers"
)

type Repository interface {
	CreateCustomer(ctx context.Context, customer *DBCustomer) (*Customer, error)
	GetCustomerByID(ctx context.Context, id uuid.UUID) (*Customer, error)
}

type SQLRepository struct {
	db *gorm.DB
}

func (s *SQLRepository) GetCustomerByID(ctx context.Context, id uuid.UUID) (*Customer, error) {
	var customer DBCustomer

	err := s.db.Clauses(dbresolver.Write).WithContext(ctx).Table(tableName).Where("id = ?", id).First(&customer).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return FromDBCustomer(&customer), err
}

func (s *SQLRepository) CreateCustomer(ctx context.Context, customer *DBCustomer) (*Customer, error) {
	if customer.ID == uuid.Nil {
		customer.ID = uuid.New()
	}

	result := s.db.WithContext(ctx).Table(tableName).Create(customer)
	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, constants.ErrRecordAlreadyExists
	}

	return FromDBCustomer(customer), nil
}

func NewSQLRepository(db *gorm.DB) *SQLRepository {
	return &SQLRepository{
		db: db,
	}
}
