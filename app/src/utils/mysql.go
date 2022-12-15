package utils

import (
	"fmt"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func UseDB() *gorm.DB {
	if db == nil {
		for {
			p_db, err := InitializeDB()
			if err == nil {
				db = p_db
				return db
			}
			fmt.Println(err.Error())
			time.Sleep(3 * time.Second)
		}
	}
	return db
}

func InitializeDB() (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("MARIADB_USERNAME"),
		os.Getenv("MARIADB_PASSWORD"),
		os.Getenv("MARIADB_HOST"),
		os.Getenv("MARIADB_PORT"),
		os.Getenv("MARIADB_NAME"),
	)

	return gorm.Open(
		mysql.Open(dsn),
		&gorm.Config{},
	)
}

func ShouldInitializeDB() *gorm.DB {
	db, err := InitializeDB()

	if err != nil {
		panic(err.Error())
	}

	return db
}
