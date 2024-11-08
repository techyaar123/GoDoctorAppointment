package services

import (
	"example.com/GoDoctor/models"
	"example.com/GoDoctor/repositories"
    "database/sql"
    "log"
)


type AdminService struct {
    Repo *repositories.AdminRepository
}


func NewAdminService(repo *repositories.AdminRepository) *AdminService {
    return &AdminService{Repo: repo}
}


func (s *AdminService) Createbookings(db *sql.DB, name string, price float64, description string) error {
    err := s.Repo.Createbookings(name, price, description)
    if err != nil {
        log.Println("Error in service layer while creating bookings:", err)
    }
    return err
}


func (s *AdminService) Deletebookings(db *sql.DB, id string) error {
    err := s.Repo.Deletebookings(id)
    if err != nil {
        log.Println("Error in service layer while deleting bookings:", err)
    }
    return err
}


func (s *AdminService) Getbookings(db *sql.DB, filter string) ([]models.bookings, error) {
    bookingss, err := s.Repo.Getbookingss(filter)
    if err != nil {
        log.Println("Error in service layer while retrieving bookingss:", err)
    }
    return bookingss, err
}

func (s *AdminService) GetbookingsDetails(db *sql.DB, id string) (models.bookings, error) {
    bookings, err := s.Repo.GetbookingsDetails(id)
    if err != nil {
        log.Println("Error in service layer while retrieving bookings details:", err)
    }
    return bookings, err
}

func (s *AdminService) Updatebookings(id string, name string, price float64, description string) error {
    err := s.Repo.Updatebookings(id, name, price, description)
    if err != nil {
        log.Println("Error in service layer while updating bookings:", err)
    }
    return err
}
