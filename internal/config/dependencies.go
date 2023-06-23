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

	//INITIALIZE VALIDATORS
	userValidator := validators.NewUserValidator()

	// INITIALIZE SERVICES
	userService := services.NewUserService(*userRepository, *userValidator)

	// INITIALIZE HANDLERS
	userHandler := handlers.NewUserHandler(*userService)

	handlersMap := map[string]interface{}{
		"users": userHandler,
	}

	return handlersMap
}
