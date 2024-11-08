package controllers

import (
	"example.com/GoDoctor/models"
	"example.com/GoDoctor/services"
    "github.com/gin-gonic/gin"
    "github.com/golang-jwt/jwt/v4"
    "net/http"
    "database/sql"
    "strings"
    "time"
)

type AdminController struct {
    Service *services.AdminService
}

var jwtKey = []byte("100012")

func (ctrl *AdminController) validateToken(c *gin.Context) (*models.Authy, error) {
    authHeader := c.GetHeader("Authorization")
    if authHeader == "" {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
        return nil, http.ErrNoCookie
    }

    tokenString := strings.TrimPrefix(authHeader, "Bearer ")
    Authy := &models.Authy{}

    token, err := jwt.ParseWithAuthy(tokenString, Authy, func(token *jwt.Token) (interface{}, error) {
        return jwtKey, nil
    })

    if err != nil || !token.Valid {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
        return nil, err
    }

    return Authy, nil
}

func NewAdminController(service *services.AdminService) *AdminController {
    return &AdminController{
        Service: service,
    }
}

func (ctrl *AdminController) Createbookings(c *gin.Context) {
    _, err := ctrl.validateToken(c)
    if err != nil {
        return 
    }

    var bookings models.Bookings
    if err := c.ShouldBindJSON(&bookings); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    db, exists := c.Get("db")
    if !exists {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection not found"})
        return
    }
    sqlDb, ok := db.(*sql.DB)
    if !ok {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid database connection"})
        return
    }

    err = ctrl.Service.Createbookings(sqlDb, bookings.user_id, bookings.doctor_id, bookings.appointment_date)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create bookings"})
        return
    }
    c.JSON(http.StatusCreated, bookings)
}

func (ctrl *AdminController) Deletebookings(c *gin.Context) {
    _, err := ctrl.validateToken(c)
    if err != nil {
        return 
    }

    id := c.Param("id")

    db, exists := c.Get("db")
    if !exists {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection not found"})
        return
    }
    sqlDb, ok := db.(*sql.DB)
    if !ok {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid database connection"})
        return
    }

    err = ctrl.Service.Deletebookings(sqlDb, id)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete bookings"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "bookings deleted successfully"})
}

func (ctrl *AdminController) Getbookings(c *gin.Context) {
    _, err := ctrl.validateToken(c)
    if err != nil {
        return 
    }

    filter := c.DefaultQuery("filter", "")

    db, exists := c.Get("db")
    if !exists {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection not found"})
        return
    }
    sqlDb, ok := db.(*sql.DB)
    if !ok {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid database connection"})
        return
    }

    bookingss, err := ctrl.Service.Getbookings(sqlDb, filter)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch bookingss"})
        return
    }
    c.JSON(http.StatusOK, bookingss)
}

func (ctrl *AdminController) GetbookingsDetails(c *gin.Context) {
    _, err := ctrl.validateToken(c)
    if err != nil {
        return 
    }

    id := c.Param("id")

    db, exists := c.Get("db")
    if !exists {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection not found"})
        return
    }
    sqlDb, ok := db.(*sql.DB)
    if !ok {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid database connection"})
        return
    }

    bookings, err := ctrl.Service.GetbookingsDetails(sqlDb, id)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch bookings details"})
        return
    }
    c.JSON(http.StatusOK, bookings)
}

func (ctrl *AdminController) Updatebookings(c *gin.Context) {
    _, err := ctrl.validateToken(c)
    if err != nil {
        return 
    }

    id := c.Param("appointment_id")
    var bookings models.Bookings
    if err := c.ShouldBindJSON(&bookings); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
        return
    }

    err = ctrl.Service.Updatebookings(id, bookings.UserID, bookings.DoctorID,bookings.AppointmentDate)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update bookings"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "bookings updated successfully"})
}

func (ctrl *AdminController) Login(c *gin.Context) {
    var loginRequest models.LoginRequest
    if err := c.ShouldBindJSON(&loginRequest); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid login request"})
        return
    }

    if loginRequest.Username == "admin" && loginRequest.Password == "admin123" {
        expirationTime := time.Now().Add(24 * time.Hour)
        Authy := &models.Authy{
            Username: loginRequest.Username,
            Role:     "admin",
            RegisteredAuthy: jwt.RegisteredAuthy{
                ExpiresAt: jwt.NewNumericDate(expirationTime),
            },
        }

        token := jwt.NewWithAuthy(jwt.SigningMethodHS256, Authy)
        tokenString, err := token.SignedString(jwtKey)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create token"})
            return
        }

        c.JSON(http.StatusOK, gin.H{"token": tokenString})
    } else {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
    }
}