package models

type Users struct {
    ID        int    `json:"user_id"`
    Username  string `json:"username"`
    Password  string `json:"password"`
    Role      string `json:"role"` // e.g., "patient", "admin"
    FirstName string `json:"first_name"`
    LastName  string `json:"last_name"`
    Email     string `json:"email"`
    Phone     string `json:"phone_number"`
    Address   string `json:"address"`
}
