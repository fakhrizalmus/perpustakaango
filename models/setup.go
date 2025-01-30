package model

import (
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := os.Getenv("DB")
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Buku{}, &Kategori{}, &User{})
	DB = db
}
