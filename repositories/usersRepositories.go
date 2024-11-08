package repositories

import (
	"database/sql"
	"fmt"
	"net/http"
	"example.com/GoDoctor/models"
	"strconv"
)

type UsersRepository struct {
	dbHandler   *sql.DB
}

func NewUsersRepository(dbHandler *sql.DB) *UsersRepository {
	return &UsersRepository{
		dbHandler: dbHandler,
	}
}

func (rr UsersRepository) CreateUser(user *models.Users) (*models.Users, *models.ResponseError) {
	query := `
		INSERT INTO Users (username, password, role, first_name, last_name, email, phone_number, address)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?)`

	res, err := rr.dbHandler.Exec(query, user.Username, user.Password, user.Role, user.FirstName,user.LastName,user.Email,user.Phone,user.Address)
	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	userId, err := res.LastInsertId()
	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	return &models.Users{
		ID: int(userId),
		Username:  user.Username,
		Password:  user.Password,
		Role:      user.Role,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Phone:     user.Phone,
		Address:   user.Address,
	}, nil
}

func (rr UsersRepository) UpdateUser(user *models.Users) *models.ResponseError {
	query := `
		UPDATE Users
		SET
			username = ? ,
			password = ?,
			role = ?,
			first_name = ?,
			last_name = ?,
			email = ?,
			phone_number = ?,
			address = ?
		WHERE email = ?`

	res, err := rr.dbHandler.Exec(query, user.Username, user.Password, user.Role, user.FirstName,user.LastName,user.Email,user.Phone,user.Address, user.Email)
	if err != nil {
		return &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	if rowsAffected == 0 {
		return &models.ResponseError{
			Message: "User not found",
			Status:  http.StatusNotFound,
		}
	}

	return nil
}



func (rr UsersRepository) DeleteUser(userId int) *models.ResponseError {
	query := `DELETE FROM Users WHERE user_id = ?`
	res, err := rr.dbHandler.Exec(query, userId)
	if err != nil {
		return &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	if rowsAffected == 0 {
		return &models.ResponseError{
			Message: "User not found",
			Status:  http.StatusNotFound,
		}
	}

	return nil
}

func (rr UsersRepository) GetUser(userId int) (*models.Users, *models.ResponseError) {
	fmt.Println(userId)
	
	query := `
		SELECT *
	FROM Users WHERE user_id=?`

	
	rows, err := rr.dbHandler.Query(query, userId)
	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	defer rows.Close()
	var user_id string
	var  first_name, last_name, username, password, role, email,phone_number,address string
	for rows.Next() {
		err := rows.Scan(&user_id, &first_name, &last_name, &username, &password, &role, &email, &phone_number,&address)
		if err != nil {
			return nil, &models.ResponseError{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
			}
		}
	}

	if rows.Err() != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	return &models.Users{
		ID:           userId,
		FirstName:    first_name,
		LastName:     last_name,
		Username:     username,
		Password:     password,
		Role:         role,
		Email:        email,
		Phone:        phone_number,
		Address:      address,
	}, nil
}

func (rr UsersRepository) GetAllUsers() ([]*models.Users, *models.ResponseError) {
	query := `
	SELECT *
	FROM Users`

	rows, err := rr.dbHandler.Query(query)
	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	defer rows.Close()

	users := make([]*models.Users, 0)
	var user_id string
	var  first_name, last_name, username, password, role, email,phone_number,address string

	for rows.Next() {
		err := rows.Scan(&user_id, &first_name, &last_name, &username, &password, &role, &email, &phone_number,&address)
		if err != nil {
			return nil, &models.ResponseError{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
			}
		}
		id,err:=strconv.Atoi(user_id)
		user := &models.Users{
			ID:           id,
			FirstName:    first_name,
			LastName:     last_name,
			Username:     username,
			Password:     password,
			Role:         role,
			Email:        email,
			Phone:        phone_number,
			Address:      address,
		}

		users = append(users, user)
	}

	if rows.Err() != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return users, nil
}

