package models

import "time"

type Appointment struct {
    ID             int       `json:"appointment_id"`
    UserID         int       `json:"user_id"`
    DoctorID       int       `json:"doctor_id"`
    AppointmentDate time.Time `json:"appointment_date"`
}
