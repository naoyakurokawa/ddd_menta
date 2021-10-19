package config

import (
	// "fmt"
	// "os"

	// "github.com/joho/godotenv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// NewDB DBと接続する
func NewDB() *gorm.DB {
	// err := godotenv.Load()

	// if err != nil {
	// 	fmt.Printf("読み込み出来ませんでした: %v", err)
	// }

	// USER := os.Getenv("DB_USER")
	// PASS := os.Getenv("DB_PASS")
	// PROTOCOL := "tcp(" + os.Getenv("DB_ADDRESS") + ")"
	// DB_NAME := os.Getenv("DB_NAME")
	// CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DB_NAME

	// db, err := gorm.Open("mysql", CONNECT)
	// if err != nil {
	// 	panic(err)
	// }
	db, err := gorm.Open("mysql", "ddd_menta:ddd_menta@tcp(dockerMySQL)/ddd_menta")
	if err != nil {
		panic(err)
	}

	return db
}
