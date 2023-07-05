package services

import (
	"database/sql"
	"khodza/rest-api/internal/app/models"
	"khodza/rest-api/internal/app/repositories"
	"khodza/rest-api/internal/app/validators"
	"net/http"

	"github.com/lib/pq"
)

type UserServiceInterface interface {
	GetUsers() ([]models.User, CustomError)
	CreateUser(user models.User) (models.User, CustomError)
	GetUser(userID int) (models.User, CustomError)
	UpdateUser(userID int, user models.User) (models.User, CustomError)
	DeleteUser(userID int) CustomError
}

type UserService struct {
	userRepository repositories.UserRepositoryInterface
	validator      validators.UserValidatorInterface
}

func NewUserService(userRepository repositories.UserRepositoryInterface, userValidator validators.UserValidatorInterface) UserServiceInterface {
	return &UserService{
		userRepository: userRepository,
		validator:      userValidator,
	}
}

// ValidationError represents a validation error with a nice status code and message
type CustomError struct {
	StatusCode int    `json:"status"`
	Message    string `json:"message"`
	Err        error  `json:"-"`
}

func (s *UserService) GetUsers() ([]models.User, CustomError) {

	users, err := s.userRepository.GetUsers()
	if err != nil {
		return []models.User{}, CustomError{
			StatusCode: http.StatusInternalServerError,
			Message:    "failed to get users",
			Err:        err,
		}
	}
	return users, CustomError{}
}

func (s *UserService) CreateUser(user models.User) (models.User, CustomError) {
	if err := s.validator.ValidateUserCreate(&user); err != nil {
		// Return a validation error with a nice status code
		return models.User{}, CustomError{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Err:        err,
		}
	}

	newUser, err := s.userRepository.CreateUser(user)
	if err != nil {
		pqErr, _ := err.(*pq.Error)
		if pqErr.Code == "23505" {

			return models.User{}, CustomError{
				StatusCode: http.StatusBadRequest,
				Message:    "email already exists",
				Err:        err,
			}

		}

		return models.User{}, CustomError{
			StatusCode: http.StatusInternalServerError,
			Message:    "failed to create user",
			Err:        err,
		}
	}

	return newUser, CustomError{}
}

func (s *UserService) GetUser(userID int) (models.User, CustomError) {
	user, err := s.userRepository.GetUser(userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.User{}, CustomError{
				StatusCode: 404,
				Message:    "No user found with given id",
				Err:        err,
			}
		}

		return models.User{}, CustomError{
			StatusCode: http.StatusInternalServerError,
			Message:    "Something went wrong while getting user",
			Err:        err,
		}
	}
	return user, CustomError{}
}

func (s *UserService) UpdateUser(userID int, user models.User) (models.User, CustomError) {
	if err := s.validator.ValidateUserUpdate(&user); err != nil {
		// Return a validation error with a nice status code
		return models.User{}, CustomError{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Err:        err,
		}
	}

	updatedUser, err := s.userRepository.UpdateUser(userID, user)
	if err != nil {
		//no found error
		if err == sql.ErrNoRows {
			return models.User{}, CustomError{
				StatusCode: 404,
				Message:    "No user found with given id",
				Err:        err,
			}
		}

		//duplicate error
		pqErr, _ := err.(*pq.Error)
		if pqErr.Code == "23505" {

			return models.User{}, CustomError{
				StatusCode: http.StatusBadRequest,
				Message:    "email already exists",
				Err:        err,
			}

		}
		//other error
		return models.User{}, CustomError{
			StatusCode: http.StatusInternalServerError,
			Message:    "Something went wrong while updating user",
			Err:        err,
		}

	}

	return updatedUser, CustomError{}
}

func (s *UserService) DeleteUser(userID int) CustomError {
	if err := s.userRepository.DeleteUser(userID); err != nil {
		//no found error
		if err == sql.ErrNoRows {
			return CustomError{
				StatusCode: 404,
				Message:    "No user found with given id",
				Err:        err,
			}
		}
		return CustomError{
			StatusCode: http.StatusInternalServerError,
			Message:    "Something went wrong while deleting user",
			Err:        err,
		}
	}
	return CustomError{}
}
