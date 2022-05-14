package db

import (
	"log"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/kelseyhightower/envconfig"
)

var (
	once sync.Once
	db   *gorm.DB
)

type Config struct {
	DbUser     string `required:"true" default:"ddd_menta""`
	DbPassword string `required:"true" default:"ddd_menta""`
	DbHost     string `required:"true" default:"localhost""`
	DbPort     string `required:"true" default:"3306""`
	DbName     string `required:"true" default:"ddd_menta""`
}

func NewDB() *gorm.DB {
	var config Config
	if err := envconfig.Process("", &config); err != nil {
		log.Fatalf("[ERROR] Failed to process env: %s", err.Error())
	}
	dbconf := config.DbUser + ":" + config.DbPassword + "@tcp(" + config.DbHost + ":" + config.DbPort + ")/" + config.DbName + "?parseTime=true"

	once.Do(func() {
		conn, err := gorm.Open("mysql", dbconf)
		if err != nil {
			panic(err)
		}
		db = conn
	})

	return db
}
