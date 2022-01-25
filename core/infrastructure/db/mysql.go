package db

import (
	"sync"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var (
	once sync.Once
	db   *gorm.DB
)

func NewDB() *gorm.DB {
	//todo: 環境変数化
	once.Do(func() {
		conn, err := gorm.Open("mysql", "ddd_menta:ddd_menta@tcp(localhost)/ddd_menta?parseTime=true")
		if err != nil {
			panic(err)
		}
		db = conn
	})

	return db
}
