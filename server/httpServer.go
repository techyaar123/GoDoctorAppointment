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
	
}

func InitHttpServer(config *viper.Viper, dbHandler *sql.DB) HttpServer {
	usersRepository := repositories.NewUsersRepository(dbHandler)
	usersService := services.NewUsersService(usersRepository)
	usersController := controllers.NewUsersController(usersService)

	router := gin.Default()

	router.POST("/users", usersController.CreateUser)
	router.PUT("/users", usersController.UpdateUser)
	router.DELETE("/user/:id", usersController.DeleteUser)
	router.GET("/user/:id", usersController.GetUser)
	router.GET("/users", usersController.GetUsersBatch)


	return HttpServer{
		config:            config,
		router:            router,
		usersController: usersController,
	}
}

func (hs HttpServer) Start() {
	err := hs.router.Run(hs.config.GetString("http.server_address"))
	if err != nil {
		log.Fatalf("Error while starting HTTP server: %v", err)
	}
}
