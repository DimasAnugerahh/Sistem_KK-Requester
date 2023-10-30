package config

import (
	"fmt"
	"kk-requester/model/domain"

	"os"
	"log"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func InitDB() (*gorm.DB, error) {
	var err error
	_, err = os.Stat(".env")
	if err == nil {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Failed to fetch .env file")
		}
	}

	dbuser := os.Getenv("DB_USER")
	dbpass := os.Getenv("DB_PASSWORD")
	dbhost := os.Getenv("DB_HOST")
	dbport := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbuser, dbpass, dbhost, dbport, dbname)

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return DB, err
	}

	return DB, nil
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
