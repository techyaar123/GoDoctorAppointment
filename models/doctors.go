


package models

type Doctor struct {
    ID                int    `json:"id"`
    Name              string `json:"name"`
    Specialty         string `json:"specialty"`
    AvailabilityTiming string `json:"availability_timing"` // e.g., "09:00 - 17:00"
}
