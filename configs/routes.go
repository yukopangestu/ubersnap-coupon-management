package route

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func SetupRoutes(e *echo.Echo, db *gorm.DB) {
	e.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})

	// Example of using db
	e.GET("/db-check", func(c echo.Context) error {
		sqlDB, err := db.DB()
		if err != nil {
			return c.String(http.StatusInternalServerError, "Failed to get generic database object")
		}
		if err := sqlDB.Ping(); err != nil {
			return c.String(http.StatusInternalServerError, "Database connection failed")
		}
		return c.String(http.StatusOK, "Database connected")
	})
}
