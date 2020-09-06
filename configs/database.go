// Package configs database configs
package configs

import (
	"fmt"
	"os"

	"github.com/mcuadros/go-defaults"
	"gorm.io/gorm"
)

// DB GORM DB instance
var DB *gorm.DB

// DBConfig database configuration
type DBConfig struct {
	Host     string `default:"0.0.0.0"`
	Port     string `default:"5432"`
	User     string `default:"postgres"`
	DBName   string `default:"postgres"`
	Password string `default:"postgres"`
}

// BuildDBConfig build DB config
func BuildDBConfig() *DBConfig {
	dbConfig := DBConfig{
		Host:     os.Getenv("PG_HOST"),
		Port:     os.Getenv("PG_PORT"),
		User:     os.Getenv("PG_USER"),
		DBName:   os.Getenv("PG_DBNAME"),
		Password: os.Getenv("PG_PASSWORD"),
	}
	defaults.SetDefaults(&dbConfig)
	return &dbConfig
}

// BuildDSN build Postgresql's DSN
func BuildDSN() string {
	dbConfig := BuildDBConfig()
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		dbConfig.Host,
		dbConfig.User,
		dbConfig.Password,
		dbConfig.DBName,
		dbConfig.Port,
	)
}
