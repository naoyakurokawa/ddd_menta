package middleware

import (
	"log"
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/naoyakurokawa/ddd_menta/config"
)

func DBMiddlewareFunc(db *gorm.DB) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return echo.HandlerFunc(func(c echo.Context) error {
			var tx *gorm.DB
			var err error
			if isTransactionMethod(c.Request().Method) {
				tx = db.Begin()
				c.Set(config.DBKey, tx)
			} else {
				c.Set(config.DBKey, tx)
			}
			if err = next(c); err != nil {
				if tx != nil {
					tx.Rollback()
					log.Printf("failed to handle transaction Rollback: %+v", err)
					return err
				}
				return nil
			}
			if tx != nil {
				tx.Commit()
				return nil
			}
			return nil
		})
	}
}

func isTransactionMethod(method string) bool {
	if method == http.MethodPost ||
		method == http.MethodPut ||
		method == http.MethodDelete {
		return true
	}

	return false
}
