package config

import (
	"khodza/rest-api/internal/app/handlers"
	"khodza/rest-api/internal/app/repositories"
	"khodza/rest-api/internal/app/services"
	"khodza/rest-api/internal/app/validators"
)

func InitDependencies() map[string]interface{} {
	// INITIALIZE REPOSITORIES
	userRepository := repositories.NewUserRepository(GetDB())
	productRepository := repositories.NewProductRepository(GetDB())

	//INITIALIZE VALIDATORS
	userValidator := validators.NewUserValidator()
	productValidator := validators.NewProductValidator()

	// INITIALIZE SERVICES
	userService := services.NewUserService(*userRepository, *userValidator)
	productService := services.NewProductService(*productRepository, *productValidator)

	// INITIALIZE HANDLERS
	userHandler := handlers.NewUserHandler(*userService)
	productHandler := handlers.NewProductHandler(*productService)

	handlersMap := map[string]interface{}{
		"users":    userHandler,
		"products": productHandler,
	}

	return handlersMap
}
