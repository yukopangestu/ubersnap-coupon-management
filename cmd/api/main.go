package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"

	route "github.com/username/go-webapp/configs"
	"github.com/username/go-webapp/internal/handler"
)

func main() {
	// 1. Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using default environment variables")
	}

	// 2. Database Connection
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	// Default to localhost/3306 if not set (for local dev without docker-compose networking)
	if dbHost == "" {
		dbHost = "127.0.0.1"
	}
	if dbPort == "" {
		dbPort = "3306"
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&allowPublicKeyRetrieval=true", dbUser, dbPass, dbHost, dbPort, dbName)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Verify connection
	if err := db.Ping(); err != nil {
		log.Printf("Warning: Could not connect to database: %v\n", err)
	} else {
		log.Println("Successfully connected to the database")
	}

	// 3. Initialize Echo
	e := echo.New()

	// 4. Setup Routes
	// Pass Echo instance and DB connection to routes
	route.SetupRoutes(e, db)

	// Wrap the existing standard handler
	e.GET("/", echo.WrapHandler(http.HandlerFunc(handler.HelloHandler)))

	// 5. Start Server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	e.Logger.Fatal(e.Start(":" + port))
}
