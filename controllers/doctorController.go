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

type DoctorsController struct {
	doctorsService *services.DoctorsService
}

func NewDoctorsController(doctorsService *services.DoctorsService) *DoctorsController {
	return &DoctorsController{
		doctorsService: doctorsService,
	}
}

func (rh DoctorsController) CreateDoctor(ctx *gin.Context) {
	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println("Error while reading create runner request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var doctor models.Doctors
	err = json.Unmarshal(body, &doctor)
	if err != nil {
		log.Println("Error while unmarshaling create runner request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	response, responseErr := rh.doctorsService.CreateDoctor(&doctor)
	if responseErr != nil {
		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (rh DoctorsController) UpdateDoctor(ctx *gin.Context) {
	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println("Error while reading update runner request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var doctor models.Doctors
	err = json.Unmarshal(body, &doctor)
	if err != nil {
		log.Println("Error while unmarshaling update runner request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	responseErr := rh.doctorsService.UpdateDoctor(&doctor)
	if responseErr != nil {
		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
		return
	}

	ctx.Status(http.StatusNoContent)
}

func (rh DoctorsController) DeleteDoctor(ctx *gin.Context) {
	userId := ctx.Param("id")
	id,err:=strconv.Atoi(userId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	responseErr := rh.doctorsService.DeleteDoctor(id)
	if responseErr != nil {
		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
		return
	}

	ctx.Status(http.StatusNoContent)
}

func (rh DoctorsController) GetDoctor(ctx *gin.Context) {
	userId := ctx.Param("id")

	id,err:=strconv.Atoi(userId)
	fmt.Println("id======= ",id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	response, responseErr := rh.doctorsService.GetDoctor(id)
	if responseErr != nil {
		ctx.JSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (rh DoctorsController) GetDoctorsBatch(ctx *gin.Context) {

	response, responseErr := rh.doctorsService.GetDoctorsBatch()
	if responseErr != nil {
		ctx.JSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)
}


