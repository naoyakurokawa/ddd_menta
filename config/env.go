package config

import (
	"time"

	"github.com/kelseyhightower/envconfig"
)

var (
	JST = time.FixedZone("Asia/Tokyo", 9*60*60)
)

type ENV struct {
	GOEnv      string `envconfig:"GO_ENV" default:"local" required:"true"`
	DBUser     string `required:"true" default:"ddd_menta"`
	DBPassword string `required:"true" default:"ddd_menta"`
	DBHost     string `required:"true" default:"localhost"`
	DBPort     string `required:"true" default:"3306"`
	DBName     string `required:"true" default:"ddd_menta"`
	MailHost   string `required:"true" default:"localhost"`
	MailPort   int    `required:"true" default:"1025"`
}

var Env ENV

func init() {
	if err := envconfig.Process("", &Env); err != nil {
		panic(err)
	}
}
