package database

import (
	"log"

	"BTPNS_Fulldev/app"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
    var err error
    dsn := "zul:123@tcp(127.0.0.1:3306)/finalgo?charset=utf8mb4&parseTime=True&loc=Local" // Pastikan DSN benar
    DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("failed to connect to database: %v", err)
    }
    // AutoMigrate untuk membuat tabel jika belum ada
    if err := DB.AutoMigrate(&app.User{}, &app.Photo{}); err != nil {
        log.Fatalf("failed to migrate database: %v", err)
    }
}
