package repository

import (
	"GoAuth/src/database"
	"GoAuth/src/models"
	"fmt"
)

type PostgresqlRepository[T models.Model] struct{}

// NewPostgresqlRepository creates a new instance of PostgresqlRepository
func NewPostgresqlRepository[T models.Model]() *PostgresqlRepository[T] {
	return &PostgresqlRepository[T]{}
}

// List method retrieves all
func (r *PostgresqlRepository[T]) List() ([]T, error) {
	var model []T
	err := database.GetInstance().GetClient().Find(&model).Error
	if err != nil {
		return nil, err
	}

	return model, nil
}

// ListByColumn method retrieves all base on condition
func (r *PostgresqlRepository[T]) ListByColumn(columns map[string]any) ([]*T, error) {
	var model []*T
	query := database.GetInstance().GetClient().Model(&model)

	for key, value := range columns {
		query = query.Where(fmt.Sprintf("%s = ?", key), value)
	}
	err := query.Find(&model).Error
	if err != nil {
		return nil, err
	}

	return model, nil
}

// ListByColumnWithGreaterThan method retrieves all base on condition and some greater that
func (r *PostgresqlRepository[T]) ListByColumnWithGreaterThan(columns map[string]any, greater map[string]any) ([]*T, error) {
	var model []*T
	query := database.GetInstance().GetClient().Model(&model)

	for key, value := range columns {
		query = query.Where(fmt.Sprintf("%s = ?", key), value)
	}

	for key, value := range greater {
		query = query.Where(fmt.Sprintf("%s > ?", key), value)
	}

	err := query.Find(&model).Error
	if err != nil {
		return nil, err
	}

	return model, nil
}

// Get method retrieves a models by its ID
func (r *PostgresqlRepository[T]) Get(id uint) (*T, error) {
	var model T
	err := database.GetInstance().GetClient().First(&model, id).Error
	if err != nil {
		return nil, err
	}

	return &model, nil
}

// GetByColumn retrieves a models by it's columns
func (r *PostgresqlRepository[T]) GetByColumn(columns map[string]any) (*T, error) {
	var model T

	query := database.GetInstance().GetClient().Model(&model)

	for key, value := range columns {
		query = query.Where(fmt.Sprintf("%s = ?", key), value)
	}
	err := query.First(&model).Error
	if err != nil {
		return nil, err
	}

	return &model, nil
}

// Create method Create a models
func (r *PostgresqlRepository[T]) Create(model T) (*T, error) {
	err := database.GetInstance().GetClient().Create(&model).Error
	if err != nil {
		return nil, err
	}

	return &model, err
}

// Update method Update a models
func (r *PostgresqlRepository[T]) Update(model T) (*T, error) {
	err := database.GetInstance().GetClient().Save(&model).Error
	if err != nil {
		return nil, err
	}

	return &model, nil
}

// Delete method delete a models by its ID
func (r *PostgresqlRepository[T]) Delete(model T) error {
	err := database.GetInstance().GetClient().Delete(&model).Error
	if err != nil {
		return err
	}

	return nil
}

// HardDelete method that  delete a models by its ID hardly
func (r *PostgresqlRepository[T]) HardDelete(model T) error {
	err := database.GetInstance().GetClient().Unscoped().Delete(&model).Error
	if err != nil {
		return err
	}

	return nil
}
