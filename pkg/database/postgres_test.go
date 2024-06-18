package database

import (
	"log"
	"testing"
)

var config = &PostgresConfig{
	Host:     "localhost",
	User:     "postgres",
	Password: "12345678",
	DBName:   "testing",
	Port:     "5432",
	SSLMode:  "disable",
	TimeZone: "Asia/Yangon",
}

func TestNewPostgres(t *testing.T) {
	postgre, err := NewPostgres(config)

	if err != nil {
		t.Error(err)
	}

	if postgre == nil {
		t.Fatal("postgre is nil")
	}

	log.Print("database name ", postgre.db.Name())

	postgre.Close()
}
