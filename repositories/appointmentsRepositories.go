package repositories

import (
	"database/sql"
	"fmt"
	"net/http"
	"example.com/GoDoctor/models"
	"strconv"
)

type AppointmentsRepository struct {
	dbHandler   *sql.DB
}

func NewAppointmentsRepository(dbHandler *sql.DB) *AppointmentsRepository {
	return &AppointmentsRepository{
		dbHandler: dbHandler,
	}
}

func (rr AppointmentsRepository) CreateAppointment(appointment *models.Appointments) (*models.Appointments, *models.ResponseError) {
	query := `
		INSERT INTO Appointments (user_id,doctor_id,appointment_date)
		VALUES (?, ?, ?)`

	res, err := rr.dbHandler.Exec(query, appointment.UserID,appointment.DoctorID,appointment.AppointmentDate)
	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	appointmentId, err := res.LastInsertId()
	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	return &models.Appointments{
		ID: int(appointmentId),
		DoctorID: appointment.DoctorID,
		UserID:appointment.UserID,
		AppointmentDate:appointment.AppointmentDate
	
	}, nil
}

func (rr AppointmentsRepository) UpdateAppointment(appointment *models.Appointments) *models.ResponseError {
	query := `
		UPDATE Appointments
		SET
			doctor_id = ? ,
			user_id = ? ,
			appointment_data = ?
			
		WHERE appointment_id = ?`

	res, err := rr.dbHandler.Exec(query, appointment.DoctorID,appointment.UserID,appointment.AppointmentDate,appointment.ID)
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
			Message: "Appointment not found",
			Status:  http.StatusNotFound,
		}
	}

	return nil
}



func (rr AppointmentsRepository) DeleteAppointment(appointmentId int) *models.ResponseError {
	query := `DELETE FROM Appointments WHERE appointment_id = ?`
	res, err := rr.dbHandler.Exec(query, appointmentId)
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
			Message: "Appointment not found",
			Status:  http.StatusNotFound,
		}
	}

	return nil
}

func (rr AppointmentsRepository) GetAppointment(appointmentId int) (*models.Appointments, *models.ResponseError) {
	fmt.Println(appointmentId)
	
	query := `
		SELECT *
	FROM Appointments WHERE appointment_id=?`

	
	rows, err := rr.dbHandler.Query(query, appointmentId)
	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	defer rows.Close()
	var appointment_id string
	var  user_id,doctor_id,appointment_data string
	for rows.Next() {
		err := rows.Scan(&appointment_id, &user_id,&doctor_id,&appointment_data)
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
	return &models.Appointments{
		ID:           appointmentId,
		UserID:    user_id,
		DoctorID:  doctor_id
		AppointmentDate: appointment_data
	}, nil
}

func (rr AppointmentsRepository) GetAllAppointments() ([]*models.Appointments, *models.ResponseError) {
	query := `
	SELECT *
	FROM Appointments`

	rows, err := rr.dbHandler.Query(query)
	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	defer rows.Close()

	appointments := make([]*models.Appointments, 0)
	var appointment_id string
	var user_id,doctor_id,appointment_data string

	for rows.Next() {
		err := rows.Scan(&appointment_id,&user_id,&doctor_id,&appointment_data)
		if err != nil {
			return nil, &models.ResponseError{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
			}
		}
		id,err:=strconv.Atoi(appointment_id)
		appointment := &models.Appointments{
			ID:           id,
			UserID:    user_id,
			DoctorID: doctor_id,
			AppointmentDate: appointment_data
		}

		appointments = append(appointments, appointment)
	}

	if rows.Err() != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return appointments, nil
}

