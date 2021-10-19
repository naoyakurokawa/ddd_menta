package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"

	"github.com/jinzhu/gorm"
)

// NewDB DBと接続する
func NewDB() *gorm.DB {
	err := godotenv.Load(fmt.Sprintf("../%s.env", os.Getenv("GO_ENV")))
	if err != nil {
		// .env読めなかった場合の処理
	}

	USER := os.Getenv("DB_USER")
	PASS := os.Getenv("DB_PASS")
	PROTOCOL := "tcp(" + os.Getenv("DB_ADDRESS") + ")"
	DB_NAME := os.Getenv("DB_NAME")
	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DB_NAME

	db, err := gorm.Open("mysql", CONNECT)
	if err != nil {
		panic(err)
	}

	return db
}
