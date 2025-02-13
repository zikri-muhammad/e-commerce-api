package main

import (
	"e-commerce-api/configs"
	"e-commerce-api/internal/handlers"
	"e-commerce-api/internal/models"
	"e-commerce-api/internal/repositories"
	"e-commerce-api/internal/routers"
	"e-commerce-api/internal/services"
	"e-commerce-api/internal/utils"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func loadEnvFile() error {
	envPaths := []string{
		".env",       // Try current directory
		"../../.env", // Try two levels up
		filepath.Join(os.Getenv("GOPATH"), "src/e-commerce-api/.env"), // Try GOPATH
	}

	for _, path := range envPaths {
		if err := godotenv.Load(path); err == nil {
			log.Printf("Loaded .env file from: %s", path)
			return nil
		}
	}

	return fmt.Errorf("no .env file found")
}

func main() {
	// Find project root and load .env
	rootPath, err := utils.FindRootPath()
	if err != nil {
		log.Fatalf("Error finding project root: %v", err)
	}

	if err := godotenv.Load(filepath.Join(rootPath, ".env")); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Load configurations
	configs.LoadDBCfg()
	dbConfig := configs.DBCfg()
	appConfig := configs.LoadConfig()

	// Setelah LoadDBCfg()
	log.Printf("Database config: %+v", dbConfig)
	log.Printf("DB User: '%s'", dbConfig.User)
	log.Printf("DB Password: '%s'", dbConfig.Password) // Hati-hati dengan logging password di production

	// Setup database connection
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.Name,
	)
	// Tambahkan log untuk debug
	log.Printf("Connecting to database with DSN: %s", dsn)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Configure connection pool
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get database instance: %v", err)
	}

	sqlDB.SetMaxIdleConns(dbConfig.MaxIdleConn)
	sqlDB.SetMaxOpenConns(dbConfig.MaxOpenConn)
	sqlDB.SetConnMaxLifetime(dbConfig.MaxConnLifetime)

	// Setelah koneksi database berhasil, tambahkan:
	if err := db.AutoMigrate(&models.Product{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	// Initialize Fiber app
	app := fiber.New()

	// Setup the repository with DB connection
	productRepo := repositories.NewProductRepository(db)

	// Setup the service, injecting the repository
	productService := services.NewProductService(*productRepo)

	// Setup the handler, injecting the service
	productHandler := handlers.NewProductHandler(productService)

	// Setup routes
	routers.SetupProductRoutes(app, productHandler)

	// Run the server
	serverPort := fmt.Sprintf(":%s", appConfig.ServerPort)
	log.Fatal(app.Listen(serverPort))
}
