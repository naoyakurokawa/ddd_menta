package db

import (
	"sync"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/naoyakurokawa/ddd_menta/config"
)

var (
	once sync.Once
	db   *gorm.DB
)

func NewDB() *gorm.DB {
	dbconf := config.Env.DBUser + ":" + config.Env.DBPassword + "@tcp(" + config.Env.DBHost + ":" + config.Env.DBPort + ")/" + config.Env.DBName + "?parseTime=true"

	once.Do(func() {
		conn, err := gorm.Open("mysql", dbconf)
		if err != nil {
			panic(err)
		}
		db = conn
	})

	return db
}
