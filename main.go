package main

import (
	"context"
	"github.com/aronipurwanto/go-api-northwind/config"
	"github.com/aronipurwanto/go-api-northwind/controllers"
	"github.com/aronipurwanto/go-api-northwind/internal/soap"
	"github.com/aronipurwanto/go-api-northwind/repositories"
	"github.com/aronipurwanto/go-api-northwind/routes"
	"github.com/aronipurwanto/go-api-northwind/services"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"log"

	redisClient "github.com/aronipurwanto/go-api-northwind/internal/redis"
)

func main() {
	cfg := config.LoadConfig()

	db, err := gorm.Open(sqlserver.Open(cfg.DBSource), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to DB: ", err)
	}
	log.Println("Connected to DB")

	redis := redisClient.NewRedisClient(cfg.RedisHost, cfg.RedisPort, cfg.RedisPass)
	pong, err := redis.Ping(context.Background()).Result()
	if err != nil {
		log.Fatal("Failed to connect to Redis:", err)
	}
	log.Println("Connected to Redis:", pong)

	empRepo := repositories.NewEmployeeRepository(db)
	empService := services.NewEmployeeService(empRepo)
	empController := controllers.NewEmployeeController(empService)

	authRepo := repositories.NewAuthRepository(db)
	authService := services.NewAuthService(authRepo, cfg.JWTSecret, redis)
	authController := controllers.NewAuthController(authService)

	categoryRepo := repositories.NewCategoryRepository(db)
	categoryService := services.NewCategoryService(categoryRepo)
	categoryController := controllers.NewCategoryController(categoryService)

	productRepo := repositories.NewProductRepository(db)
	productService := services.NewProductService(productRepo)
	productController := controllers.NewProductController(productService)

	orderRepo := repositories.NewOrderRepository(db)
	orderService := services.NewOrderService(orderRepo)
	orderController := controllers.NewOrderController(orderService)

	// Init SOAP Client
	soapClient := soap.NewSOAPClient(cfg)
	log.Println("Connected to SOAP client")
	soapController := controllers.NewSOAPController(soapClient)

	app := fiber.New()

	routes.SetupRoutes(app, cfg, redis, authController, empController, categoryController, productController,
		orderController, soapController)
	err = app.Listen(":" + cfg.Port)
	if err != nil {
		log.Fatal("Failed start server: ", err)
	}
	log.Println("Server started")
}
