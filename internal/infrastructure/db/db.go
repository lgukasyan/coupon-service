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

	db, err := gorm.Open(postgres.Open(configDB.DSN), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})

	db.DisableForeignKeyConstraintWhenMigrating = true

	if err != nil {
		log.Fatal(err)
	}

	return db, nil
}
