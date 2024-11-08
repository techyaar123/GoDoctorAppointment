package services

import (
	"net/http"
	"example.com/GoDoctor/models"
	"example.com/GoDoctor/repositories"
)

type DoctorsService struct {
	doctorsRepository *repositories.DoctorsRepository
}

func NewDoctorsService(doctorsRepository *repositories.DoctorsRepository) *DoctorsService {
	return &DoctorsService{
		doctorsRepository: doctorsRepository,
	}
}

func (rs DoctorsService) CreateDoctor(doctor *models.Doctors) (*models.Doctors, *models.ResponseError) {
	responseErr := validateDoctor(doctor)
	if responseErr != nil {
		return nil, responseErr
	}
	return rs.doctorsRepository.CreateDoctor(doctor)
}

func (rs DoctorsService) UpdateDoctor(doctor *models.Doctors) *models.ResponseError {
	

	return rs.doctorsRepository.UpdateDoctor(doctor)
}

func (rs DoctorsService) DeleteDoctor(doctorId int) *models.ResponseError {
	

	return rs.doctorsRepository.DeleteDoctor(doctorId)
}

func (rs DoctorsService) GetDoctor(doctorId int) (*models.Doctors, *models.ResponseError) {
	responseErr := validateDoctorId(doctorId)
	if responseErr != nil {
		return nil, responseErr
	}

	doctor, responseErr := rs.doctorsRepository.GetDoctor(doctorId)
	if responseErr != nil {
		return nil, responseErr
	}

	return doctor, nil
}

func (rs DoctorsService) GetDoctorsBatch() ([]*models.Doctors, *models.ResponseError) {

	return rs.doctorsRepository.GetAllDoctors()
}


func validateDoctor(doctor *models.Doctors) *models.ResponseError {
	if doctor.Name == "" {
		return &models.ResponseError{
			Message: "Invalid Doctor Name",
			Status:  http.StatusBadRequest,
		}
	}

	return nil
}

func validateDoctorId(doctorId int) *models.ResponseError {
	if doctorId == 0 {
		return &models.ResponseError{
			Message: "Invalid doctor ID",
			Status:  http.StatusBadRequest,
		}
	}

	return nil
}

