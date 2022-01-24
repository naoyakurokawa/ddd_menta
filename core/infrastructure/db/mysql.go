package db

import (
	"sync"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var (
	once sync.Once
	db   *gorm.DB
	err  error
)

// NewDB DBと接続する
func NewDB() *gorm.DB {
	//todo: 環境変数化
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

	once.Do(func() {
		db, err = gorm.Open("mysql", "ddd_menta:ddd_menta@tcp(localhost)/ddd_menta?parseTime=true")
		if err != nil {
			panic(err)
		}
	})

	return db
}
