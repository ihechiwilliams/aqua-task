package resources

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
)

const (
	tableName = "resources"
)

type Repository interface {
	CreateResourcesByCustomerID(ctx context.Context, customerID uuid.UUID, resources []*DBResource) ([]*Resource, error)
	GetResourcesByCustomerID(ctx context.Context, customerID uuid.UUID) ([]*Resource, error)
	GetResourceByID(ctx context.Context, id uuid.UUID) (*Resource, error)
	UpdateResource(ctx context.Context, resource *Resource) error
	DeleteResource(ctx context.Context, id uuid.UUID) error
}

type SQLRepository struct {
	db *gorm.DB
}

func (s *SQLRepository) CreateResourcesByCustomerID(ctx context.Context, customerID uuid.UUID, resources []*DBResource) ([]*Resource, error) {
	resourceNames := make([]string, 0, len(resources))
	for _, resource := range resources {
		resourceNames = append(resourceNames, resource.Name)
	}

	var existingNames []string
	if err := s.db.Clauses(dbresolver.Write).WithContext(ctx).Table(tableName).
		Where("customer_id = ? AND name IN ?", customerID, resourceNames).
		Pluck("name", &existingNames).Error; err != nil {
		return nil, fmt.Errorf("failed to check existing resource names: %w", err)
	}

	nameSet := make(map[string]struct{})
	for _, name := range existingNames {
		nameSet[name] = struct{}{}
	}

	var filteredResources []*DBResource

	for _, resource := range resources {
		if _, exists := nameSet[resource.Name]; !exists {
			resource.CustomerID = customerID
			filteredResources = append(filteredResources, resource)
		}
	}

	if len(filteredResources) > 0 {
		if err := s.db.WithContext(ctx).Table(tableName).Create(&filteredResources).Error; err != nil {
			return nil, fmt.Errorf("failed to create resources: %w", err)
		}
	}

	return FromDBResourceList(filteredResources), nil
}

func (s *SQLRepository) GetResourcesByCustomerID(ctx context.Context, customerID uuid.UUID) ([]*Resource, error) {
	var resources []*DBResource

	err := s.db.Clauses(dbresolver.Write).WithContext(ctx).Table(tableName).Where("customer_id = ?", customerID).Find(&resources).Error
	if err != nil {
		return nil, err
	}

	return FromDBResourceList(resources), nil
}

func (s *SQLRepository) GetResourceByID(ctx context.Context, id uuid.UUID) (*Resource, error) {
	var resource DBResource

	err := s.db.Clauses(dbresolver.Write).WithContext(ctx).Table(tableName).Where("id = ?", id).First(&resource).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return FromDBResource(&resource), err
}

func (s *SQLRepository) UpdateResource(ctx context.Context, resource *Resource) error {
	return s.db.WithContext(ctx).Save(resource).Error
}

func (s *SQLRepository) DeleteResource(ctx context.Context, id uuid.UUID) error {
	result := s.db.WithContext(ctx).Where("id = ?", id).Delete(&Resource{})
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("no resource found with the given ID")
	}

	return nil
}

func NewSQLRepository(db *gorm.DB) *SQLRepository {
	return &SQLRepository{
		db: db,
	}
}
