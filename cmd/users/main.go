package main

import (
	"fmt"
	"github.com/mr-kmh/go-gin-starterkit/internal/features/users/app"
	"github.com/spf13/viper"
)

func LoadConfig() {
	viper.AddConfigPath(".")
	viper.SetConfigFile("./configs/config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
}

func main() {
	LoadConfig()
	app.Run()
}
