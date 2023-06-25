package config

import (
	"khodza/rest-api/internal/app/handlers"
	"khodza/rest-api/internal/app/repositories"
	"khodza/rest-api/internal/app/services"
	"khodza/rest-api/internal/app/validators"

	"go.uber.org/zap"
)

func InitDependencies(logger *zap.Logger) map[string]interface{} {

	// INITIALIZE UTILS
	handlerUtilities := handlers.NewHandlerUtilities(logger)
	// INITIALIZE REPOSITORIES
	userRepository := repositories.NewUserRepository(GetDB())
	productRepository := repositories.NewProductRepository(GetDB())
	orderRepository := repositories.NewOrderRepository(GetDB())

	//INITIALIZE VALIDATORS
	userValidator := validators.NewUserValidator()
	productValidator := validators.NewProductValidator()

	// INITIALIZE SERVICES
	userService := services.NewUserService(*userRepository, *userValidator)
	productService := services.NewProductService(*productRepository, *productValidator)
	orderService := services.NewOrderService(*orderRepository, *productService)

	// INITIALIZE HANDLERS
	userHandler := handlers.NewUserHandler(*userService, *handlerUtilities)
	productHandler := handlers.NewProductHandler(*productService, *handlerUtilities)
	orderHandler := handlers.NewOrderHandler(*orderService, *handlerUtilities)

	handlersMap := map[string]interface{}{
		"users":    userHandler,
		"products": productHandler,
		"orders":   orderHandler,
	}

	return handlersMap
}
