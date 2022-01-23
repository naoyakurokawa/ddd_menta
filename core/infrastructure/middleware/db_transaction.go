package middleware

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/naoyakurokawa/ddd_menta/config"
)

func DBMiddlewareFunc(db *gorm.DB) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return echo.HandlerFunc(func(c echo.Context) error {
			tx := db.Begin()
			c.Set(config.DBKey, tx)
			if err := next(c); err != nil {
				tx.Rollback()
				log.Printf("failed to handle transaction Rollback: %+v", err)
				return err
			}
			tx.Commit()
			return nil
		})
	}
}
