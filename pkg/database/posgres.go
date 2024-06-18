package database

import (
	"fmt"
	"log"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var once sync.Once

type PostgresConfig struct {
	Host     string
	User     string
	Password string
	DBName   string
	Port     string
	SSLMode  string
	TimeZone string
}

type Postgres struct {
	db *gorm.DB
}

func dsn(config *PostgresConfig) string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		config.Host, config.User, config.Password, config.DBName, config.Port, config.SSLMode, config.TimeZone)
}

func NewPostgres(config *PostgresConfig) (*Postgres, error) {
	var db *gorm.DB
	var err error

	once.Do(func() {
		db, err = gorm.Open(postgres.New(postgres.Config{
			DriverName:           "pgx",
			DSN:                  dsn(config),
			PreferSimpleProtocol: true,
		}), &gorm.Config{})
	})

	if err != nil {
		return nil, err
	}

	return &Postgres{db: db}, nil
}

func (p *Postgres) Migrate(dst ...any) {
	if err := p.db.AutoMigrate(dst...); err != nil {
		log.Fatal("failed to migrate database")
	}
}

func (p *Postgres) GetInstance() *gorm.DB {
	return p.db
}

func (p *Postgres) Close() {
	log.Print("closing database ", p.db.Name())
	db, err := p.db.DB()
	if err != nil {
		log.Printf("error getting db connection: %s", err.Error())
		return
	}
	err = db.Close()
	if err != nil {
		log.Printf("failed to close db connection: %s", err.Error())
		return
	}
}
