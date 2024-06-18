package app

import (
	"github.com/gin-gonic/gin"
	"github.com/mr-kmh/go-gin-starterkit/internal/features/users/api"
	"github.com/mr-kmh/go-gin-starterkit/internal/features/users/entities"
	"github.com/mr-kmh/go-gin-starterkit/internal/features/users/infrastructure"
	"github.com/mr-kmh/go-gin-starterkit/internal/features/users/usecases"
	"github.com/mr-kmh/go-gin-starterkit/pkg/database"
	"github.com/spf13/viper"
	"log"
)

func Run() {
	port := viper.GetString("server.port")
	config := &database.PostgresConfig{
		Host:     viper.GetString("postgres.host"),
		User:     viper.GetString("postgres.user"),
		Password: viper.GetString("postgres.password"),
		DBName:   viper.GetString("postgres.dbname"),
		Port:     viper.GetString("postgres.port"),
		SSLMode:  viper.GetString("postgres.sslmode"),
		TimeZone: viper.GetString("postgres.timezone"),
	}

	db, _ := database.NewPostgres(config)
	db.Migrate(entities.User{})
	defer db.Close()

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	err := r.SetTrustedProxies([]string{"192.168.1.0/24"})
	if err != nil {
		log.Fatalf("Error setting trusted proxies: %v", err)
	}

	repo := infrastructure.NewRepository(db.GetInstance())
	services := usecases.NewServices(repo)
	handlers := api.NewREST(services)
	api.REST(r, handlers)

	log.Fatal(r.Run(":" + port))
}
