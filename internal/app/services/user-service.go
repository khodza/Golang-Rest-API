package services

import (
	"khodza/rest-api/internal/app/models"
	"khodza/rest-api/internal/app/repositories"
	"khodza/rest-api/internal/app/validators"
)

type UserServiceInterface interface {
	GetUsers() ([]models.User, error)
	CreateUser(user models.User) (models.User, error)
	GetUser(userID int) (models.User, error)
	UpdateUser(userID int, user models.User) (models.User, error)
	DeleteUser(userID int) error
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

func (s *UserService) GetUsers() ([]models.User, error) {

	users, err := s.userRepository.GetUsers()
	if err != nil {
		return []models.User{}, err
	}
	return users, err
}

func (s *UserService) CreateUser(user models.User) (models.User, error) {
	//validation
	if err := s.validator.ValidateUserCreate(&user); err != nil {
		return models.User{}, err
	}

	newUser, err := s.userRepository.CreateUser(user)
	if err != nil {
		return models.User{}, err
	}

	return newUser, err
}

func (s *UserService) GetUser(userID int) (models.User, error) {
	user, err := s.userRepository.GetUser(userID)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (s *UserService) UpdateUser(userID int, user models.User) (models.User, error) {
	if err := s.validator.ValidateUserUpdate(&user); err != nil {
		return models.User{}, err
	}

	updatedUser, err := s.userRepository.UpdateUser(userID, user)
	if err != nil {
		return models.User{}, err
	}

	return updatedUser, nil
}

func (s *UserService) DeleteUser(userID int) error {
	if err := s.userRepository.DeleteUser(userID); err != nil {
		return err
	}
	return nil
}
