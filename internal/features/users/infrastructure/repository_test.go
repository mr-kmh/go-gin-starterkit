package infrastructure

import (
	"github.com/mr-kmh/go-gin-starterkit/internal/features/users/entities"
	"github.com/mr-kmh/go-gin-starterkit/pkg/database"
	"testing"
)

func setup(t *testing.T) (*database.Postgres, func()) {
	config := &database.PostgresConfig{
		Host:     "localhost",
		User:     "postgres",
		Password: "12345678",
		DBName:   "testing",
		Port:     "5432",
		SSLMode:  "disable",
		TimeZone: "Asia/Yangon",
	}

	postgres, err := database.NewPostgres(config)
	if err != nil {
		t.Fatal("Failed to connect database ", err)
	}

	postgres.Migrate(&entities.User{})

	return postgres, func() {
		postgres.Close()
	}
}

func TestGetAll(t *testing.T) {
	postgres, close := setup(t)
	defer close()

	repo := NewRepository(postgres.GetInstance())
	users, err := repo.GetAll()

	if err != nil {
		t.Error("[user repository] err must be nil")
	}

	if len(users) != 0 {
		t.Log("[user repository] users muse be empty")
	}
}
