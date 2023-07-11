package dependencies

import (
	"fmt"
	custom_errors "khodza/rest-api/internal/app/errors"
	"khodza/rest-api/internal/app/handlers"
	"khodza/rest-api/internal/app/repositories"
	"khodza/rest-api/internal/app/services"
	"khodza/rest-api/internal/app/validators"
	"khodza/rest-api/internal/database"
	"khodza/rest-api/internal/logger"

	"go.uber.org/zap"
)

func InitDependencies() (*custom_errors.GlobalErrorHandler, map[string]interface{}, *zap.Logger, error) {
	// INITIALIZE LOGGER
	logger, err := logger.CreateLogger()
	if err != nil {
		return nil, nil, nil, fmt.Errorf("error on initializing logger")
	}
	// Get the DB instance
	db := database.GetDB()
	if db == nil {
		return nil, nil, nil, fmt.Errorf("error on getting db instance")
	}

	// INITIALIZE REPOSITORIES
	userRepository := repositories.NewUserRepository(db)
	productRepository := repositories.NewProductRepository(db)
	orderRepository := repositories.NewOrderRepository(db)

	//INITIALIZE VALIDATORS
	userValidator := validators.NewUserValidator()
	productValidator := validators.NewProductValidator()

	// INITIALIZE SERVICES
	userService := services.NewUserService(userRepository, userValidator)
	productService := services.NewProductService(productRepository, productValidator)
	orderService := services.NewOrderService(orderRepository, productService)
	paymentService := services.NewPaymentService(orderService)

	// INITIALIZE HANDLERS
	userHandler := handlers.NewUserHandler(userService, logger)
	productHandler := handlers.NewProductHandler(productService, logger)
	orderHandler := handlers.NewOrderHandler(orderService, logger)
	paymentHandler := handlers.NewPaymentHandler(paymentService, logger)

	//INITIALIZE Global Error Handler
	globalErrorHandler := custom_errors.NewGlobalErrorHandler(logger)

	handlersMap := map[string]interface{}{
		"users":    userHandler,
		"products": productHandler,
		"orders":   orderHandler,
		"payments": paymentHandler,
	}

	return globalErrorHandler, handlersMap, logger, nil
}
