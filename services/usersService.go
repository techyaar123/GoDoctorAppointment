package services

import (
	"net/http"
	"example.com/GoDoctor/models"
	"example.com/GoDoctor/repositories"
	"strconv"
	"time"
)

type UsersService struct {
	usersRepository *repositories.UsersRepository
}

func NewUsersService(usersRepository *repositories.UsersRepository) *UsersService {
	return &UsersService{
		usersRepository: usersRepository,
	}
}

func (rs UsersService) CreateUser(user *models.Users) (*models.Users, *models.ResponseError) {
	responseErr := validateRunner(user)
	if responseErr != nil {
		return nil, responseErr
	}

	return rs.usersRepository.CreateUser(user)
}

func (rs UsersService) UpdateUser(user *models.User) *models.ResponseError {
	responseErr := validateUserId(user.ID)
	if responseErr != nil {
		return responseErr
	}

	responseErr = validateUser(user)
	if responseErr != nil {
		return responseErr
	}

	return rs.usersRepository.UpdateUser(user)
}

func (rs UsersService) DeleteUser(userId string) *models.ResponseError {
	responseErr := validateUserId(userId)
	if responseErr != nil {
		return responseErr
	}

	return rs.usersRepository.DeleteUser(userId)
}

func (rs UsersService) GetUser(userId string) (*models.Users, *models.ResponseError) {
	responseErr := validateUserId(userId)
	if responseErr != nil {
		return nil, responseErr
	}

	user, responseErr := rs.usersRepository.GetUser(userId)
	if responseErr != nil {
		return nil, responseErr
	}

	return user, nil
}

func (rs UsersService) GetUsersBatch() ([]*models.Users, *models.ResponseError) {

	return rs.runnersRepository.GetAllRunners()
}


func validateUser(user *models.User) *models.ResponseError {
	if user.FirstName == "" {
		return &models.ResponseError{
			Message: "Invalid first name",
			Status:  http.StatusBadRequest,
		}
	}

	if user.LastName == "" {
		return &models.ResponseError{
			Message: "Invalid last name",
			Status:  http.StatusBadRequest}
	}

	if user.Username == "" {
		return &models.ReponseError(
			Message: "Invalid user name"
			Status: http.StatusBadRequest}
		)
	}

	 if user.Password == "" {
                return &models.ReponseError(
                        Message: "Invalid password"
                        Status: http.StatusBadRequest}
                )
        }

	 if user.Role == "" {
                return &models.ReponseError(
                        Message: "Invalid role"
                        Status: http.StatusBadRequest}
                )
        }

	 if user.Email == "" {
                return &models.ReponseError(
                        Message: "Invalid user email"
                        Status: http.StatusBadRequest}
                )
        }

	 if user.Phone == "" {
                return &models.ReponseError(
                        Message: "Invalid phone number"
                        Status: http.StatusBadRequest}
                )
        }

	 if user.Address == "" {
                return &models.ReponseError(
                        Message: "Invalid address"
                        Status: http.StatusBadRequest}
                )
        }

	return nil
}

func validateUserId(userId string) *models.ResponseError {
	if userId == "" {
		return &models.ResponseError{
			Message: "Invalid user ID",
			Status:  http.StatusBadRequest,
		}
	}

	return nil
}

