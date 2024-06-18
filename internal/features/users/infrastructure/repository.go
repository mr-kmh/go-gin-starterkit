package infrastructure

import (
	"github.com/mr-kmh/go-gin-starterkit/internal/features/users/entities"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) GetAll() ([]entities.User, error) {
	users := []entities.User{}
	r.db.Find(&users)
	return users, nil
}
