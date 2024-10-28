package repository

import (
	"GoAuth/src/models"
	"os"
)

type IRepository[T models.Model] interface {
	List() ([]*T, error)
	Get(id uint) (*T, error)
	GetByColumn(columns map[string]any) (*T, error)
	Create(model T) (*T, error)
	Update(model T) (*T, error)
	Delete(model T) error
	HardDelete(model T) error
}

func GetRepository[T models.Model]() IRepository[T] {
	dbType := os.Getenv("DB_DRIVER")

	switch dbType {
	case "postgresql":
		return NewPostgresqlRepository[T]()
	default:
		return NewPostgresqlRepository[T]()
	}
}
