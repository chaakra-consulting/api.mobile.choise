package handlers

import (
	"api.choise/configs"
	"api.choise/internals/helpers"
	"api.choise/internals/models"
	"github.com/gin-gonic/gin"
)

func MSDTQuestion(c *gin.Context) {
	var msdtQuestion []models.MSDTTestQuestion
	err := configs.DB.Find(&msdtQuestion).Error
	if err != nil {
		helpers.ErrorResponse(c, 400, err)
		return
	}
	helpers.SuccessResponse(c, 200, msdtQuestion)
}
