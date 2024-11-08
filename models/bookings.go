package models


type Booking struct {
    ID         int `json:"booking_id"`
    UserID     int `json:"user_id"`
    DoctorID int `json:"doctor_id"`
    AppointmentDate   int `json:"appointment_date"`
}
