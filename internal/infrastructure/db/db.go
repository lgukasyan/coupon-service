package db

import (
	"coupon_service/config"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type IDatabase interface {
	Connect() (*gorm.DB, error)
}

type Database struct{}

func New() IDatabase { return &Database{} }

func (d *Database) Connect() (*gorm.DB, error) {
	configDB, err := config.NewDBConfig()
	if err != nil {
		return nil, err
	}

	db, err := gorm.Open(postgres.New(postgres.Config{DSN: configDB.DSN}))

	if err != nil {
		log.Fatal(err)
	}

	return db, nil
}
