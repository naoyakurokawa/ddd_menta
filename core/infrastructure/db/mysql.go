package db

import (
	"os"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var (
	once sync.Once
	db   *gorm.DB
)

func NewDB() *gorm.DB {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")
	dbconf := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + name + "?parseTime=true"

	once.Do(func() {
		conn, err := gorm.Open("mysql", dbconf)
		if err != nil {
			panic(err)
		}
		db = conn
	})

	return db
}
