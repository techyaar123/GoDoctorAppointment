package repositories

import (
    "example.com/GoDoctor/models"
    "database/sql"
    "log"
)


type AdminRepository struct {
    db *sql.DB
}


func NewAdminRepository(db *sql.DB) *AdminRepository {
    return &AdminRepository{db: db}
}


func (repo *AdminRepository) Createbookings(user_id int, doctor_id int, appointment_date string) error {
    _, err := repo.db.Exec(`INSERT INTO Appointments (user_id,doctor_id,appointment_date) VALUES (?,?,?)`, user_id, doctor_id, appointment_date)
    if err != nil {
        log.Println("Error creating bookings:", err)
    }
    return err
}


func (repo *AdminRepository) Deletebookings(id string) error {
    _, err := repo.db.Exec(`DELETE FROM Appointments WHERE appointment_id = ?`, id)
    if err != nil {
        log.Println("Error deleting bookings:", err)
    }
    return err
}


func (repo *AdminRepository) Getbookings(filter string) ([]models.Bookings, error) {
    var bookingss []models.bookings
    rows, err := repo.db.Query(`SELECT user_id,doctor_id,appointment_date FROM Appointments WHERE appointment_Date LIKE ?`, "%"+filter+"%")
    if err != nil {
        log.Println("Error fetching bookingss:", err)
        return bookingss, err
    }
    defer rows.Close()

    for rows.Next() {
        var bookings models.bookings
        if err := rows.Scan(&bookings.ID, &bookings.UserID, &bookings.DoctorID, &bookings.AppointmentDate); err != nil {
            log.Println("Error scanning bookings:", err)
            return bookingss, err
        }
        bookingss = append(bookingss, bookings)
    }

    return bookingss, nil
}


func (repo *AdminRepository) GetbookingsDetails(id string) (models.bookings, error) {
    var bookings models.bookings
    err := repo.db.QueryRow(`SELECT user_id,doctor_id ,appointment_date from Appointments WHERE appointment_id = ?`, id).
        Scan(&bookings.ID, &bookings.UserID, &bookings.DoctorID, &bookings.AppointmentDate)
    if err != nil {
        log.Println("Error fetching bookings details:", err)
    }
    return bookings, err
}


func (repo *AdminRepository) Updatebookings(appointment_id int,user_id int, doctor_id int, appointment_date string) error {
    _, err := repo.db.Exec(`UPDATE Appointments SET user_id = ?, doctor_id = ?, availabilty_date= ? WHERE Appointments_appointment_id = ?`, user_id, doctor_id, appointment_date, id)
    if err != nil {
        log.Println("Error updating bookings:", err)
    }
    return err
}
