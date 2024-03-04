package repository

import (
	"coupon_service/internal/domain/model"
	"fmt"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestMain(m *testing.M) {
	if err := godotenv.Load("../../../config/.env"); err != nil {
		panic(err)
	}

	os.Exit(m.Run())
}

func TestCouponRepository(t *testing.T) {
	var (
		host     = os.Getenv("POSTGRES_HOST")
		user     = os.Getenv("POSTGRES_USER")
		password = os.Getenv("POSTGRES_PASSWORD")
		dbName   = os.Getenv("POSTGRES_DB_NAME_TEST")
		port     = os.Getenv("POSTGRES_PORT")
	)

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", host, user, password, dbName, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})

	if err != nil {
		t.Error(err)
	}

	// Migrate
	if err := db.AutoMigrate(&model.Coupon{}); err != nil {
		t.Error(err)
	}

	// At the end, drop table coupons
	defer db.Exec("DROP TABLE coupons")

	// Create new repository
	repo := NewCouponRepository(db)

	// Coupon
	coupon := &model.Coupon{
		Code:           "12345",
		Discount:       10,
		MinBasketValue: 100,
	}

	// Case FindByCode
	t.Run("Should return false ", func(t *testing.T) {
		ok, err := repo.FindByCode(coupon.Code)
		assert.Equal(t, false, ok)
		assert.NoError(t, err, "should be nil")
	})

	// Case Create
	t.Run("Should return nil", func(t *testing.T) {
		ok, err := repo.FindByCode(coupon.Code)
		assert.Equal(t, false, ok)
		assert.NoError(t, err, "should be nil")

		err = repo.Create(coupon)
		assert.NoError(t, err, "should be nil")
	})

	// Case FindByCode already exists
	t.Run("Should return true", func(t *testing.T) {
		ok, err := repo.FindByCode(coupon.Code)
		assert.Equal(t, true, ok)
		assert.NoError(t, err, "should be nil")
	})
}
