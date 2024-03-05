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
		Code:           "DESC1",
		Discount:       1,
		MinBasketValue: 100,
	}

	// Case FindByCode
	t.Run("Should return false", func(t *testing.T) {
		ok, err := repo.Exists(coupon.Code)
		assert.Equal(t, false, ok)
		assert.NoError(t, err, "should be nil")
	})

	// Case Create
	t.Run("Should return nil", func(t *testing.T) {
		ok, err := repo.Exists(coupon.Code)
		assert.Equal(t, false, ok)
		assert.NoError(t, err, "should be nil")

		err = repo.Create(coupon)
		assert.NoError(t, err, "should be nil")
	})

	// Case FindByCode already exists
	t.Run("Should return true", func(t *testing.T) {
		ok, err := repo.Exists(coupon.Code)
		assert.Equal(t, true, ok)
		assert.NoError(t, err, "should be nil")
	})

	// Case Get coupon codes
	t.Run("Get coupons", func(t *testing.T) {
		codes, err := repo.Get()
		assert.NotEmpty(t, codes, "shouldn't be empty")
		assert.NoError(t, err, "should be nil")
	})

	// FindByCode
	t.Run("Get coupon by Code", func(t *testing.T) {
		coupon, err := repo.FindByCode("DESC1")
		assert.NotEmpty(t, coupon, "shouldn't be empty")
		assert.NoError(t, err, "should be nil")
	})

	db.Exec("DROP TABLE coupons")
	if err := db.AutoMigrate(&model.Coupon{}); err != nil {
		t.Error(err)
	}

	t.Run("Get empty coupons", func(t *testing.T) {
		codes, err := repo.Get()
		assert.Empty(t, codes, "should be empty")
		assert.NoError(t, err, "should be nil")
	})
}
