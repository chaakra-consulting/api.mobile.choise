package handlers

import (
	"api.choise/configs"
	"api.choise/internals/helpers"
	"api.choise/internals/models"
	"github.com/gin-gonic/gin"
)

func DISCQuestion(c *gin.Context) {
	var discQuestions []models.DISCTestQuestion
	err := configs.DB.Find(&discQuestions).Error
	if err != nil {
		helpers.ErrorResponse(c, 400, err)
		return
	}
	helpers.SuccessResponse(c, 200, discQuestions)
}
