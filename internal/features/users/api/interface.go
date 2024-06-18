package api

import "github.com/mr-kmh/go-gin-starterkit/internal/features/users/entities"

type IService interface {
	GetUsers() ([]entities.User, error)
}
