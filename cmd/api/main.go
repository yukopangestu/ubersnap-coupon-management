package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	route "github.com/username/go-webapp/configs"
	"github.com/username/go-webapp/internal/handler"
	"github.com/username/go-webapp/internal/model"
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

	// Default to localhost/3306 if not set
	if dbHost == "" {
		dbHost = "127.0.0.1"
	}
	if dbPort == "" {
		dbPort = "3306"
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort, dbName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Fatal error connecting to database: %v", err)
	}

	// 3. Auto Migrate
	log.Println("Running AutoMigrate...")
	if err := db.AutoMigrate(&model.User{}, &model.Coupon{}, &model.CouponUsage{}); err != nil {
		log.Fatalf("Fatal error during migration: %v", err)
	}
	log.Println("AutoMigrate complete.")

	// 4. Initialize Echo
	e := echo.New()

	// 5. Setup Routes
	// Pass Echo instance and GORM DB connection to routes
	route.SetupRoutes(e, db)

	// Wrap the existing standard handler
	e.GET("/", echo.WrapHandler(http.HandlerFunc(handler.HelloHandler)))

	// 6. Start Server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	e.Logger.Fatal(e.Start(":" + port))
}
