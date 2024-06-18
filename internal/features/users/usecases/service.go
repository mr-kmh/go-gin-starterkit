package usecases

import "github.com/mr-kmh/go-gin-starterkit/internal/features/users/entities"

type Services struct {
	repository IRepository
}

func NewServices(repository IRepository) *Services {
	return &Services{repository: repository}
}

func (s *Services) GetUsers() ([]entities.User, error) {
	return s.repository.GetAll()
}
