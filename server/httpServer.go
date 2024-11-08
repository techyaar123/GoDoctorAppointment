package server

import (
	"database/sql"
	"log"
	"example.com/GoDoctor/controllers"
	"example.com/GoDoctor/repositories"
	"example.com/GoDoctor/services"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type HttpServer struct {
	config            *viper.Viper
	router            *gin.Engine
	usersController *controllers.UsersController
	doctorsController *controllers.DoctorsController
	appointmentsController *controllers.AppointmentsController
	adminController *controllers.AdminController
	
}

func InitHttpServer(config *viper.Viper, dbHandler *sql.DB) HttpServer {
	usersRepository := repositories.NewUsersRepository(dbHandler)
	usersService := services.NewUsersService(usersRepository)
	usersController := controllers.NewUsersController(usersService)

	appointmentsRepository := repositories.NewAppointsRepository(dbHandler)
	appointmentsService := services.NewAppointmentsService(appointmentsRepository)
	appointmentsController := controllers.NewAppointmentsController(appointmentsService)

	doctorsRepository := repositories.NewDoctorsRepository(dbHandler)
	doctorsService := services.NewDoctorsService(doctorsRepository)
	doctorsController := controllers.NewDoctorsController(doctorsService)

	adminRepo := repositories.NewAdminRepository(dbHandler)
	adminService := services.NewAdminService(adminRepo)
	adminController := controllers.NewAdminController(adminService)

	router := gin.Default()

	router.POST("/users", usersController.CreateUser)
	router.PUT("/users", usersController.UpdateUser)
	router.DELETE("/user/:id", usersController.DeleteUser)
	router.GET("/user/:id", usersController.GetUser)
	router.GET("/users", usersController.GetUsersBatch)

	router.POST("/doctors", doctorsController.CreateDoctor)
	router.PUT("/doctors", doctorsController.UpdateDoctor)
	router.DELETE("/doctor/:id", doctorsController.DeleteDoctor)
	router.GET("/doctor/:id", doctorsController.GetDoctor)
	router.GET("/doctors", doctorsController.GetDoctorsBatch)

	router.POST("/appointments", appointmentsController.CreateAppointment)
	router.PUT("/appointments", appointmentsController.UpdateAppointment)
	router.DELETE("/appointment/:id", appointmentsController.DeleteAppointment)
	router.GET("/appointment/:id", appointmentsController.GetAppointment)
	router.GET("/appointments", appointmentsController.GetAppointmentsBatch)


	adminGroup := router.Group("/admin")
    adminGroup.POST("/bookingss", adminController.Createbookings)
    adminGroup.DELETE("/bookings/:id", adminController.Deletebookings)
    adminGroup.GET("/bookings", adminController.Getbookings)
    adminGroup.GET("/bookings/:id", adminController.GetbookingDetails)
    adminGroup.PUT("/bookings/:id", adminController.Updatebooking)
	return HttpServer{
		config:            config,
		router:            router,
		usersController: usersController,
		doctorsController: doctorsController,
		appointmentsController: appointmentsController,
		adminController: adminController,
	}
}

func (hs HttpServer) Start() {
	err := hs.router.Run(hs.config.GetString("http.server_address"))
	if err != nil {
		log.Fatalf("Error while starting HTTP server: %v", err)
	}
}
