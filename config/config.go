package config

import (
	"fmt"
	"kk-requester/model/domain"

	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func InitDB() (*gorm.DB, error) {
	godotenv.Load(".env")

	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"))

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return db, err
	}

	return db, nil
}

func InitialMigration(DB *gorm.DB) {
	DB.AutoMigrate(
		&domain.Account{},

		&domain.KTP{},

		&domain.AktaKelahiran{},

		&domain.SuratPindah{},

		&domain.SuratNikah{},

		&domain.AktaKematian{},

		&domain.RequestKK{},

		&domain.KK{},
	)
}
