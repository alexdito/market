package common

import (
	"application/internal/app/config"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type App struct {
	DataBase *gorm.DB
}

func GetApp() (*App, error) {
	DataBaseConnection, err := getConnection()

	return &App{
		DataBase: DataBaseConnection,
	}, err
}

func getConnection() (*gorm.DB, error) {
	conf := config.New()

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Europe/Moscow",
		conf.DB.Host, conf.DB.User, conf.DB.Password, conf.DB.DBName, conf.DB.Port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	return db, err
}
