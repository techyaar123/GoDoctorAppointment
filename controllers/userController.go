package controllers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"example.com/GoDoctor/models"
	"example.com/GoDoctor/services"

	"github.com/gin-gonic/gin"
)

type UsersController struct {
	usersService *services.UsersService
}

func NewUsersController(usersService *services.UsersService) *UsersController {
	return &UsersController{
		usersService: usersService,
	}
}

func (rh UsersController) CreateUser(ctx *gin.Context) {
	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println("Error while reading create runner request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var runner models.Runner
	err = json.Unmarshal(body, &runner)
	if err != nil {
		log.Println("Error while unmarshaling create runner request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	response, responseErr := rh.runnersService.CreateRunner(&runner)
	if responseErr != nil {
		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (rh UsersController) UpdateUser(ctx *gin.Context) {
	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println("Error while reading update runner request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var user models.Users
	err = json.Unmarshal(body, &user)
	if err != nil {
		log.Println("Error while unmarshaling update runner request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	responseErr := rh.usersService.UpdateUser(&user)
	if responseErr != nil {
		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
		return
	}

	ctx.Status(http.StatusNoContent)
}

func (rh UsersController) DeleteUser(ctx *gin.Context) {
	userId := ctx.Param("user_id")

	responseErr := rh.usersService.DeleteUser(userId)
	if responseErr != nil {
		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
		return
	}

	ctx.Status(http.StatusNoContent)
}

func (rh UsersController) GetUser(ctx *gin.Context) {
	userId := ctx.Param("user_id")

	response, responseErr := rh.usersService.GetUser(userId)
	if responseErr != nil {
		ctx.JSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (rh UsersController) GetUsersBatch(ctx *gin.Context) {
	params := ctx.Request.URL.Query()

	response, responseErr := rh.usersService.GetRunnersBatch()
	if responseErr != nil {
		ctx.JSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)
}


