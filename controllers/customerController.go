
package controllers

import (
	"example.com/GoDoctor/models"
	"example.com/GoDoctor/services"
    "github.com/gin-gonic/gin"
    "github.com/golang-jwt/jwt/v4"
    "net/http"
    "strconv"
    "database/sql"
    "strings"
    "time"
    "fmt"
)

type CustomerController struct {
    CustomerService *services.CustomerService
    PaymentService *services.PaymentService 
}

var jwtKey = []byte("your_secret_key")

func getDbConnection(c *gin.Context) (*sql.DB, error) {
    db, exists := c.Get("db")
    if !exists {
        return nil, fmt.Errorf("Database connection not found")
    }
    sqlDb, ok := db.(*sql.DB)
    if !ok {
        return nil, fmt.Errorf("Invalid database connection")
    }
    return sqlDb, nil
}


func (ctrl *CustomerController) validateToken(c *gin.Context) (*models.Claims, error) {
    authHeader := c.GetHeader("Authorization")
    if authHeader == "" {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
        return nil, http.ErrNoCookie
    }

    tokenString := strings.TrimPrefix(authHeader, "Bearer ")
    claims := &models.Claims{}

    token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
        return jwtKey, nil
    })

    if err != nil {
        if err == jwt.ErrSignatureInvalid {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token signature"})
        } else {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Token parsing failed"})
        }
        return nil, err
    }

    if claims.ExpiresAt != nil && claims.ExpiresAt.Before(time.Now()) {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Token has expired"})
        return nil, nil
    }

    return claims, nil
}

func NewCustomerController(customerService *services.CustomerService, paymentService *services.PaymentService) *CustomerController {
    return &CustomerController{CustomerService: customerService, PaymentService: paymentService}
}


func (ctrl *CustomerController) Getbookingss(c *gin.Context) {
    _, err := ctrl.validateToken(c)
    if err != nil {
        return
    }

    filter := c.DefaultQuery("filter", "")
    db, err := getDbConnection(c)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    bookingss, err := ctrl.CustomerService.Getbookingss(db, filter)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch bookingss", "details": err.Error()})
        return
    }
    c.JSON(http.StatusOK, bookingss)
}

func (ctrl *CustomerController) GetAppointsDetails(c *gin.Context) {
    _, err := ctrl.validateToken(c)
    if err != nil {
        return
    }

    id := c.Param("id")
    db, err := getDbConnection(c)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    bookingss, err := ctrl.CustomerService.GetbookingsDetails(db, id)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch bookingss details"})
        return
    }
    c.JSON(http.StatusOK, bookingss)
}

func (ctrl *CustomerController) AddTobookings(c *gin.Context) {
    _, err := ctrl.validateToken(c)
    if err != nil {
        return
    }

    var bookingItem models.bookings
    if err := c.ShouldBindJSON(&bookingItem); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    db, err := getDbConnection(c)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    err = ctrl.CustomerService.AddTobookings(db, bookingItem.UserID, bookingItem.DoctorID, bookingItem.bookingsDate)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add  to booking"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "bookingss added to booking"})
}

func (ctrl *CustomerController) Getbookings(c *gin.Context) {
    _, err := ctrl.validateToken(c)
    if err != nil {
        return
    }

    customerID, err := strconv.Atoi(c.Param("customer_id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid customer ID"})
        return
    }

    db, err := getDbConnection(c)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    booking, err := ctrl.CustomerService.Getbookings(db, customerID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve booking"})
        return
    }
    c.JSON(http.StatusOK, booking)
}

func (ctrl *CustomerController) Clearbookings(c *gin.Context) {
    _, err := ctrl.validateToken(c)
    if err != nil {
        return
    }

    customerID, err := strconv.Atoi(c.Param("customer_id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid customer ID"})
        return
    }

    db, err := getDbConnection(c)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    err = ctrl.CustomerService.Clearbookings(db, customerID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to clear booking"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "bookings cleared"})
}


func (ctrl *CustomerController) ProcessPayment(c *gin.Context) {
    _, err := ctrl.validateToken(c)
    if err != nil {
        return
    }

    customerID, err := strconv.Atoi(c.Param("customer_id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid customer ID"})
        return
    }

    db, err := getDbConnection(c)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    bookingItems, err := ctrl.CustomerService.Getbookings(db, customerID)
    if err != nil || len(bookingItems) == 0 {
        c.JSON(http.StatusBadRequest, gin.H{"error": "bookings is empty or failed to retrieve items"})
        return
    }

    var totalAmount float64
    for _, item := range bookingItems {
        totalAmount += item.Price * float64(item.Quantity)
    }

    paymentStatus, err := ctrl.PaymentService.ProcessPayment(db, customerID, totalAmount)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Payment processing failed", "details": err.Error()})
        return
    }

    if paymentStatus == "success" {
        err = ctrl.CustomerService.Clearbookings(db, customerID)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to clear booking after payment"})
            return
        }
        c.JSON(http.StatusOK, gin.H{"message": "Payment successful, booking cleared"})
    } else {
        c.JSON(http.StatusPaymentRequired, gin.H{"error": "Payment failed"})
    }
}

func (ctrl *CustomerController) Login(c *gin.Context) {
    var loginRequest models.LoginRequest
    if err := c.ShouldBindJSON(&loginRequest); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid login request"})
        return
    }

    if loginRequest.Username == "customer" && loginRequest.Password == "customer123" {
        expirationTime := time.Now().Add(24 * time.Hour)
        claims := &models.Claims{
            Customername: loginRequest.Username,
            Role:     "user",
            
            RegisteredClaims: jwt.RegisteredClaims{
                ExpiresAt: jwt.NewNumericDate(expirationTime),
            },
        }

        token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
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


