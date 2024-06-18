package usecases

import (
	"github.com/mr-kmh/go-gin-starterkit/internal/features/users/entities"
	"testing"
)

type MockRepository struct {
	user []entities.User
}

func (r *MockRepository) GetAll() ([]entities.User, error) {
	r.user = make([]entities.User, 0)
	return r.user, nil
}

func TestNewService(t *testing.T) {
	repo := &MockRepository{}
	services := NewServices(repo)
	users, err := services.GetUsers()

	if err != nil {
		t.Error("[user service] error should be nil", err)
	}

	if users == nil {
		t.Error("[user service] users is nil")
	}
}
