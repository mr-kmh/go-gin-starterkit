package usecases

import "github.com/mr-kmh/go-gin-starterkit/internal/features/users/entities"

type IRepository interface {
	GetAll() ([]entities.User, error)
}
