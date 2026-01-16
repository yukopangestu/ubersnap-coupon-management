package route

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo, db *sql.DB) {
	e.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})

	// Example of using db
	e.GET("/db-check", func(c echo.Context) error {
		if err := db.Ping(); err != nil {
			return c.String(http.StatusInternalServerError, "Database connection failed")
		}
		return c.String(http.StatusOK, "Database connected")
	})
}
