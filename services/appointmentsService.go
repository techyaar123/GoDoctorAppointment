package services

import (
	"net/http"
	"example.com/GoDoctor/models"
	"example.com/GoDoctor/repositories"
)

type AppointmentsService struct {
	appointmentsRepository *repositories.AppointmentsRepository
}

func NewAppointmentsService(appointmentsRepository *repositories.AppointmentsRepository) *AppointmentsService {
	return &AppointmentsService{
		appointmentsRepository: appointmentsRepository,
	}
}

func (rs AppointmentsService) CreateAppointment(appointment *models.Appointments) (*models.Appointments, *models.ResponseError) {
	responseErr := validateAppointment(appointment)
	if responseErr != nil {
		return nil, responseErr
	}
	return rs.appointmentsRepository.CreateAppointment(appointment)
}

func (rs AppointmentsService) UpdateAppointment(appointment *models.Appointments) *models.ResponseError {
	

	return rs.appointmentsRepository.UpdateAppointment(appointment)
}

func (rs AppointmentsService) DeleteAppointment(appointmentId int) *models.ResponseError {
	

	return rs.appointmentsRepository.DeleteAppointment(appointmentId)
}

func (rs AppointmentsService) GetAppointment(appointmentId int) (*models.Appointments, *models.ResponseError) {
	responseErr := validateAppointmentId(appointmentId)
	if responseErr != nil {
		return nil, responseErr
	}

	appointment, responseErr := rs.appointmentsRepository.GetAppointment(appointmentId)
	if responseErr != nil {
		return nil, responseErr
	}

	return appointment, nil
}

func (rs AppointmentsService) GetAppointmentsBatch() ([]*models.Appointments, *models.ResponseError) {

	return rs.appointmentsRepository.GetAllAppointments()
}


func validateAppointment(appointment *models.Appointments) *models.ResponseError {
	if appointment.UserID == 0 {
		return &models.ResponseError{
			Message: "Invalid user id",
			Status:  http.StatusBadRequest,
		}
	}

	if appointment.AppointmentID == 0 {
		return &models.ResponseError{
			Message: "Invalid appointment id",
			Status:  http.StatusBadRequest,
		}
	}

	return nil
}

func validateAppointmentId(appointmentId int) *models.ResponseError {
	if appointmentId == 0 {
		return &models.ResponseError{
			Message: "Invalid appointment ID",
			Status:  http.StatusBadRequest,
		}
	}

	return nil
}

