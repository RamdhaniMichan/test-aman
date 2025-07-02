package main

import (
	"fmt"
	"log"
	"os"
	"test-aman/src/delivery/http"
	"test-aman/src/domain"
	"test-aman/src/middleware"
	"test-aman/src/repository"
	"test-aman/src/usecase"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Database connection
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database")
	}

	// Auto migrate
	err = db.AutoMigrate(&domain.User{}, &domain.Product{}, &domain.Transaction{})
	if err != nil {
		log.Fatal("Failed to migrate database")
	}

	// Repository
	userRepo := repository.NewUserRepository(db)
	productRepo := repository.NewProductRepository(db)
	transactionRepo := repository.NewTransactionRepository(db)

	// Usecase
	userUsecase := usecase.NewUserUsecase(userRepo)
	productUsecase := usecase.NewProductUsecase(productRepo)
	transactionUsecase := usecase.NewTransactionUsecase(transactionRepo, productRepo)

	// Handler
	userHandler := http.NewUserHandler(userUsecase)
	productHandler := http.NewProductHandler(productUsecase)
	transactionHandler := http.NewTransactionHandler(transactionUsecase)

	// Router
	router := gin.Default()

	// Public routes
	router.POST("/register", userHandler.Register)
	router.POST("/login", userHandler.Login)
	router.GET("/products", productHandler.GetAllProducts)
	router.GET("/products/:id", productHandler.GetProduct)

	// Protected routes
	protected := router.Group("/")
	protected.Use(middleware.AuthMiddleware())
	{
		// Merchant routes
		merchant := protected.Group("/")
		merchant.Use(middleware.RoleMiddleware("merchant"))
		{
			merchant.POST("/products", productHandler.CreateProduct)
			merchant.PUT("/products/:id", productHandler.UpdateProduct)
			merchant.DELETE("/products/:id", productHandler.DeleteProduct)
			merchant.GET("/merchant/products", productHandler.GetMerchantProducts)
			merchant.GET("/merchant/transactions", transactionHandler.GetMerchantTransactions)
		}

		// Customer routes
		customer := protected.Group("/")
		customer.Use(middleware.RoleMiddleware("customer"))
		{
			customer.POST("/transactions", transactionHandler.CreateTransaction)
			customer.GET("/transactions", transactionHandler.GetCustomerTransactions)
			customer.GET("/transactions/:id", transactionHandler.GetTransaction)
		}
	}

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	if err := router.Run(":" + port); err != nil {
		log.Fatal("Failed to start server")
	}
}
