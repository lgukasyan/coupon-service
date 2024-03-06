package db

import (
	"coupon_service/config"
	"log"
	"time"

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

	var db *gorm.DB

	for i := 0; i < 10; i++ {
		log.Printf("trying to connect to the psql db..., try num: %d", i+1)
		db, err = gorm.Open(postgres.Open(configDB.DSN), &gorm.Config{
			DisableForeignKeyConstraintWhenMigrating: true,
		})

		if err == nil && db != nil {
			break
		}

		time.Sleep(time.Second * 2)
	}

	db.DisableForeignKeyConstraintWhenMigrating = true

	if err != nil {
		log.Fatal(err)
	}

	return db, nil
}
