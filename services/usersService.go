package services

import (
	"net/http"
	"example.com/GoDoctor/models"
	"example.com/GoDoctor/repositories"
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
	responseErr := validateUser(user)
	if responseErr != nil {
		return nil, responseErr
	}

	return rs.usersRepository.CreateUser(user)
}

func (rs UsersService) UpdateUser(user *models.Users) *models.ResponseError {
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

func (rs UsersService) DeleteUser(userId int64) *models.ResponseError {
	responseErr := validateUserId(userId)
	if responseErr != nil {
		return responseErr
	}

	return rs.usersRepository.DeleteUser(userId)
}

func (rs UsersService) GetUser(userId int64) (*models.Users, *models.ResponseError) {
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

	return rs.usersRepository.GetAllUsers()
}


func validateUser(user *models.Users) *models.ResponseError {
	if user.FirstName == "" {
		return &models.ResponseError{
			Message: "Invalid first name",
			Status:  http.StatusBadRequest,
		}
	}

	if user.LastName == "" {
		return &models.ResponseError{
			Message: "Invalid last name",
			Status:  http.StatusBadRequest,
		}
	}

	if user.Username == "" {
		return &models.ResponseError{
			Message: "Invalid user name",
			Status: http.StatusBadRequest,
		}
	}
	

	 if user.Password == "" {
                return &models.ResponseError{
                        Message: "Invalid password",
                        Status: http.StatusBadRequest,
				}
        }

	 if user.Role == "" {
                return &models.ResponseError{
                        Message: "Invalid role",
                        Status: http.StatusBadRequest,
				}
        }

	 if user.Email == "" {
                return &models.ResponseError{
                        Message: "Invalid user email",
                        Status: http.StatusBadRequest,
				}
        }

	 if user.Phone == "" {
                return &models.ResponseError{
                        Message: "Invalid phone number",
                        Status: http.StatusBadRequest,
				}
        }

	 if user.Address == "" {
                return &models.ResponseError{
                        Message: "Invalid address",
                        Status: http.StatusBadRequest,
				}
        }

	return nil
}

func validateUserId(userId int64) *models.ResponseError {
	if userId == 0 {
		return &models.ResponseError{
			Message: "Invalid user ID",
			Status:  http.StatusBadRequest,
		}
	}

	return nil
}

