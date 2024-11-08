package repositories

import (
	"database/sql"
	"fmt"
	"net/http"
	"example.com/GoDoctor/models"
	"strconv"
)

type DoctorsRepository struct {
	dbHandler   *sql.DB
}

func NewDoctorsRepository(dbHandler *sql.DB) *DoctorsRepository {
	return &DoctorsRepository{
		dbHandler: dbHandler,
	}
}

func (rr DoctorsRepository) CreateDoctor(doctor *models.Doctors) (*models.Doctors, *models.ResponseError) {
	query := `
		INSERT INTO Doctors (name, specialty, availability_timing)
		VALUES (?, ?, ?, ?)`

	res, err := rr.dbHandler.Exec(query, doctor.Name,doctor.Specialty,doctor.AvailabilityTiming)
	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	doctorsId, err := res.LastInsertId()
	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	return &models.Doctors{
		ID: int(doctorId),
		Specialty: doctor.Specialty,
		AvailabilityTiming: doctor.AvailabilityTiming
	}, nil
}

func (rr DoctorsRepository) UpdateDoctor(doctor *models.Doctors) *models.ResponseError {
	query := `
		UPDATE Doctors
		SET
			name = ? ,
			specialty = ?,
			availability_timing = ?

		WHERE doctor_id = ?`

	res, err := rr.dbHandler.Exec(query, doctor.Name,doctor.Specialty,doctor.AvailabilityTiming)
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
			Message: "Doctor not found",
			Status:  http.StatusNotFound,
		}
	}

	return nil
}



func (rr DoctorsRepository) DeleteDoctor(doctorId int) *models.ResponseError {
	query := `DELETE FROM Doctors WHERE doctor_id = ?`
	res, err := rr.dbHandler.Exec(query, doctorId)
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
			Message: "Doctor not found",
			Status:  http.StatusNotFound,
		}
	}

	return nil
}

func (rr DoctorsRepository) GetDoctor(doctorId int) (*models.Doctors, *models.ResponseError) {
	fmt.Println(doctorId)
	
	query := `
		SELECT *
	FROM Doctors WHERE doctor_id=?`

	
	rows, err := rr.dbHandler.Query(query, doctorId)
	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	defer rows.Close()
	var doctor_id string
	var  name,specialty,availability_timing string
	for rows.Next() {
		err := rows.Scan(&doctor_id, &name,&specialty,&availability_timing)
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
	return &models.Doctors{
		ID:           doctorId,
		Name:    name,
		Specialty: specialty
		AvailabilityTiming: availability_timing 
	}, nil
}

func (rr DoctorsRepository) GetAllDoctors() ([]*models.Doctors, *models.ResponseError) {
	query := `
	SELECT *
	FROM Doctors`

	rows, err := rr.dbHandler.Query(query)
	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	defer rows.Close()

	doctors := make([]*models.Doctors, 0)
	var doctor_id string
	var  name,specialty,availability_timing string

	for rows.Next() {
		err := rows.Scan(&doctor_id, &name,&specialty,&availability_timing)
		if err != nil {
			return nil, &models.ResponseError{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
			}
		}
		id,err:=strconv.Atoi(doctor_id)
		doctor := &models.Doctors{
			ID:           id,
			Name:    name,
			Specialty: specialty,
			AvailabilityTiming: availability_timing
		}

		doctors = append(doctors, doctor)
	}

	if rows.Err() != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return doctors, nil
}

