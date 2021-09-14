package database

import (
	"log"
	"time"

	"github.com/rafamarquesrmb/rest_go_todo/src/database/migrations"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

// const dbname string = "gotodo.db"

func StartDB() {

	database, err := gorm.Open(sqlite.Open("gotodo.db"), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})

	if err != nil {
		log.Fatal("error: ", err)
	}

	db = database

	config, _ := db.DB()

	config.SetMaxIdleConns(100)
	config.SetMaxOpenConns(100)
	config.SetConnMaxLifetime(time.Hour)

	migrations.RunMigrations(db)
}

func GetDatabase() *gorm.DB {
	return db
}

func CloseConn() error {
	config, err := db.DB()
	if err != nil {
		return err
	}

	err = config.Close()
	if err != nil {
		return err
	}

	return nil
}
