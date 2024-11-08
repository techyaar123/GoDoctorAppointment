package controllers

import (
	"encoding/json"
	"io"
	"log"
	"fmt"
	"net/http"
	"example.com/GoDoctor/models"
	"example.com/GoDoctor/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AppointmentsController struct {
	appointmentsService *services.AppointmentsService
}

func NewAppointmentsController(appointmentsService *services.AppointmentsService) *AppointmentsController {
	return &AppointmentsController{
		appointmentsService: appointmentsService,
	}
}

func (rh AppointmentsController) CreateAppointment(ctx *gin.Context) {
	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println("Error while reading create runner request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var appointment models.Appointments
	err = json.Unmarshal(body, &appointment)
	if err != nil {
		log.Println("Error while unmarshaling create runner request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	response, responseErr := rh.appointmentsService.CreateAppointment(&appointment)
	if responseErr != nil {
		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (rh AppointmentsController) UpdateAppointment(ctx *gin.Context) {
	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println("Error while reading update runner request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var appointment models.Appointments
	err = json.Unmarshal(body, &appointment)
	if err != nil {
		log.Println("Error while unmarshaling update runner request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	responseErr := rh.appointmentsService.UpdateAppointment(&appointment)
	if responseErr != nil {
		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
		return
	}

	ctx.Status(http.StatusNoContent)
}

func (rh AppointmentsController) DeleteAppointment(ctx *gin.Context) {
	userId := ctx.Param("id")
	id,err:=strconv.Atoi(userId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	responseErr := rh.appointmentsService.DeleteAppointment(id)
	if responseErr != nil {
		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
		return
	}

	ctx.Status(http.StatusNoContent)
}

func (rh AppointmentsController) GetAppointment(ctx *gin.Context) {
	userId := ctx.Param("id")

	id,err:=strconv.Atoi(userId)
	fmt.Println("id======= ",id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	response, responseErr := rh.appointmentsService.GetAppointment(id)
	if responseErr != nil {
		ctx.JSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (rh AppointmentsController) GetAppointmentsBatch(ctx *gin.Context) {

	response, responseErr := rh.appointmentsService.GetAppointsBatch()
	if responseErr != nil {
		ctx.JSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)
}


