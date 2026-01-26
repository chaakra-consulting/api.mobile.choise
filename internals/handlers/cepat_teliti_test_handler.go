package handlers

import (
	"api.choise/configs"
	"api.choise/internals/helpers"
	"api.choise/internals/models"
	"github.com/gin-gonic/gin"
)

func CepatTelitiQuestion(c *gin.Context) {
	var cepatTelitiQuestions []models.CepatTelitiQuestion
	err := configs.DB.Find(&cepatTelitiQuestions).Error
	if err != nil {
		helpers.ErrorResponse(c, 400, err)
		return
	}
	helpers.SuccessResponse(c, 200, cepatTelitiQuestions)

}
