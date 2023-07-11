package service_test

import (
	"errors"
	custom_errors "khodza/rest-api/internal/app/errors"
	"khodza/rest-api/internal/app/models"
	"khodza/rest-api/internal/app/services"
	"khodza/rest-api/internal/app/services/mocks"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGetUsers(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockUserRepositoryInterface(ctrl)
	mockValidator := mocks.NewMockUserValidatorInterface(ctrl)

	// Set up expectations for the mock repository
	mockRepo.EXPECT().GetUsers().Return([]models.User{{ID: 1, Username: "Izzat", Email: "pro@gmail.com"}}, nil)

	service := services.NewUserService(mockRepo, mockValidator)

	users, err := service.GetUsers()

	assert.NoError(t, err)
	assert.Equal(t, 1, users[0].ID)
	assert.Equal(t, "Izzat", users[0].Username)
	assert.Equal(t, "pro@gmail.com", users[0].Email)
}

func TestGetUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockUserRepositoryInterface(ctrl)
	mockValidator := mocks.NewMockUserValidatorInterface(ctrl)

	mockRepo.EXPECT().GetUser(1).Return(models.User{ID: 1, Username: "Izzat", Email: "pro@gmail.com"}, nil)

	service := services.NewUserService(mockRepo, mockValidator)

	user, err := service.GetUser(1)

	assert.NoError(t, err)
	assert.Equal(t, 1, user.ID)
	assert.Equal(t, "Izzat", user.Username)
	assert.Equal(t, "pro@gmail.com", user.Email)
}

func TestGetUser_UserNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockUserRepositoryInterface(ctrl)
	mockValidator := mocks.NewMockUserValidatorInterface(ctrl)

	mockRepo.EXPECT().GetUser(1).Return(models.User{}, custom_errors.ErrUserNotFound)

	service := services.NewUserService(mockRepo, mockValidator)
	user, err := service.GetUser(1)

	assert.Error(t, err)
	assert.Equal(t, custom_errors.ErrUserNotFound, err)
	assert.Equal(t, models.User{}, user)
}

func TestCreateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockUserRepositoryInterface(ctrl)
	mockValidator := mocks.NewMockUserValidatorInterface(ctrl)

	// Set up expectations for the mock repository
	mockRepo.EXPECT().CreateUser(gomock.Any()).Return(models.User{ID: 1, Username: "Izzat", Email: "pro@gmail.com"}, nil)

	mockValidator.EXPECT().ValidateUserCreate(gomock.Any()).Return(nil)

	service := services.NewUserService(mockRepo, mockValidator)

	user := models.User{Username: "Izzat", Email: "pro@gmail.com"}
	createdUser, err := service.CreateUser(user)
	assert.NoError(t, err)
	assert.Equal(t, 1, createdUser.ID)
	assert.Equal(t, "Izzat", user.Username)
	assert.Equal(t, "pro@gmail.com", user.Email)
}

func TestCreateUser_ValidationFailed(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockUserRepositoryInterface(ctrl)
	mockValidator := mocks.NewMockUserValidatorInterface(ctrl)

	// Set up expectations for the mock validator
	mockValidator.EXPECT().ValidateUserCreate(gomock.Any()).Return(errors.New(custom_errors.ValidationErr))

	service := services.NewUserService(mockRepo, mockValidator)

	user := models.User{Username: "Izzat", Email: "pro@gmail.com"}
	createdUser, err := service.CreateUser(user)

	assert.Error(t, err)
	assert.Equal(t, custom_errors.IsValidationErr(err.Error()), true)
	assert.Equal(t, models.User{}, createdUser)
}

func TestCreateUser_UserAlreadyExists(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockUserRepositoryInterface(ctrl)
	mockValidator := mocks.NewMockUserValidatorInterface(ctrl)

	mockValidator.EXPECT().ValidateUserCreate(gomock.Any()).Return(nil)
	mockRepo.EXPECT().CreateUser(gomock.Any()).Return(models.User{}, custom_errors.ErrEmailExist)

	service := services.NewUserService(mockRepo, mockValidator)
	user := models.User{Username: "Izzat", Email: "pro@gmail.com"}

	createdUser, err := service.CreateUser(user)

	assert.Error(t, err)
	assert.Equal(t, custom_errors.ErrEmailExist, err)
	assert.Equal(t, models.User{}, createdUser)
}
func TestUserUpdate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockUserRepositoryInterface(ctrl)
	mockValidator := mocks.NewMockUserValidatorInterface(ctrl)

	mockValidator.EXPECT().ValidateUserUpdate(gomock.Any()).Return(nil)

	// Set up expectations for the mock repository
	mockRepo.EXPECT().UpdateUser(1, gomock.Any()).Return(models.User{ID: 1, Username: "Updated", Email: "updated@example.com"}, nil)

	service := services.NewUserService(mockRepo, mockValidator)

	user := models.User{ID: 1, Username: "Updated", Email: "updated@example.com"}

	updatedUser, err := service.UpdateUser(1, user)

	assert.NoError(t, err)
	assert.Equal(t, "Updated", updatedUser.Username)
	assert.Equal(t, "updated@example.com", updatedUser.Email)
}

func TestUserUpdate_ValidationFailed(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockUserRepositoryInterface(ctrl)
	mockValidator := mocks.NewMockUserValidatorInterface(ctrl)

	mockValidator.EXPECT().ValidateUserUpdate(gomock.Any()).Return(errors.New(custom_errors.ValidationErr))

	service := services.NewUserService(mockRepo, mockValidator)
	user := models.User{Email: "dadfafgmail.com"}
	updatedUser, err := service.UpdateUser(1, user)

	assert.Error(t, err)
	assert.Equal(t, custom_errors.ValidationErr, err.Error())
	assert.Equal(t, models.User{}, updatedUser)
}

func TestUserUpdate_EmailAlreadyExist(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockUserRepositoryInterface(ctrl)
	mockValidator := mocks.NewMockUserValidatorInterface(ctrl)

	mockValidator.EXPECT().ValidateUserUpdate(gomock.Any()).Return(nil)
	mockRepo.EXPECT().UpdateUser(gomock.Any(), gomock.Any()).Return(models.User{}, custom_errors.ErrEmailExist)

	service := services.NewUserService(mockRepo, mockValidator)
	user := models.User{Username: "Izzat", Email: "pro@gmail.com"}

	updatedUser, err := service.UpdateUser(1, user)

	assert.Error(t, err)
	assert.Equal(t, custom_errors.ErrEmailExist, err)
	assert.Equal(t, models.User{}, updatedUser)
}

func TestUserUpdate_UserAlreadyExists(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockUserRepositoryInterface(ctrl)
	mockValidator := mocks.NewMockUserValidatorInterface(ctrl)

	mockValidator.EXPECT().ValidateUserUpdate(gomock.Any()).Return(nil)
	mockRepo.EXPECT().UpdateUser(gomock.Any(), gomock.Any()).Return(models.User{}, custom_errors.ErrEmailExist)

	service := services.NewUserService(mockRepo, mockValidator)
	user := models.User{Username: "Izzat", Email: "pro@gmail.com"}

	createdUser, err := service.UpdateUser(1, user)

	assert.Error(t, err)
	assert.Equal(t, custom_errors.ErrEmailExist, err)
	assert.Equal(t, models.User{}, createdUser)
}
func TestUserUpdate_UserNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockUserRepositoryInterface(ctrl)
	mockValidator := mocks.NewMockUserValidatorInterface(ctrl)

	mockValidator.EXPECT().ValidateUserUpdate(gomock.Any()).Return(nil)
	mockRepo.EXPECT().UpdateUser(gomock.Any(), gomock.Any()).Return(models.User{}, custom_errors.ErrUserNotFound)

	service := services.NewUserService(mockRepo, mockValidator)
	user := models.User{Username: "Izzat", Email: "pro@gmail.com"}

	updatedUser, err := service.UpdateUser(1, user)

	assert.Error(t, err)
	assert.Equal(t, custom_errors.ErrUserNotFound, err)
	assert.Equal(t, models.User{}, updatedUser)
}

func TestDeleteUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockUserRepositoryInterface(ctrl)
	mockValidator := mocks.NewMockUserValidatorInterface(ctrl)

	mockRepo.EXPECT().DeleteUser(1).Return(nil)

	service := services.NewUserService(mockRepo, mockValidator)

	err := service.DeleteUser(1)

	assert.NoError(t, err)
}
