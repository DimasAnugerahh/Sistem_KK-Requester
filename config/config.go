package config

import (
	"fmt"
	"kk-requester/model/domain"

	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error

	dbuser := os.Getenv("DB_USERNAME")
	dbpass := os.Getenv("DB_PASSWORD")
	dbhost := os.Getenv("DB_HOST")
	dbport := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbuser, dbpass, dbhost, dbport, dbname)

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Could not open database")
	}
	fmt.Println("Koneksi database berhasil")

}

func InitialMigration() {
	err := DB.AutoMigrate(&domain.Account{}, &domain.KTP{}, &domain.AktaKelahiran{}, &domain.SuratPindah{}, &domain.SuratNikah{}, &domain.AktaKematian{}, &domain.RequestKK{}, &domain.KK{})
	if err != nil {
		log.Fatal("Gagal migrasi database")
	}
	fmt.Println("Migrasi table berhasil")
}
