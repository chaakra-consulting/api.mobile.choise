package handlers

import (
	"api.choise/configs"
	"api.choise/internals/helpers"
	"api.choise/internals/models"
	"api.choise/internals/requests"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func CepatTelitiQuestion(c *gin.Context) {
	var cepatTelitiQuestions []models.CepatTelitiQuestion
	err := configs.DB.Find(&cepatTelitiQuestions).Order("id_soal ASC").Error
	if err != nil {
		helpers.ErrorResponse(c, 400, err)
		return
	}
	helpers.SuccessResponse(c, 200, cepatTelitiQuestions)
}

func CepatTelitiAnswerByExamNumber(c *gin.Context) {
	var cepatTelitiAnswer []models.CepatTelitiTestAnswer
	id_lowongan := c.Query("id_lowongan")
	id_pelamar := c.Query("id_pelamar")
	err := configs.DB.Where("id_lowongan = ? AND id_pelamar = ?", id_lowongan, id_pelamar).Find(&cepatTelitiAnswer).Error
	if err != nil {
		helpers.ErrorResponse(c, 400, err)
		return
	}
	helpers.SuccessResponse(c, 200, cepatTelitiAnswer)
}

func CepatTelitiPostAnswer(c *gin.Context) {
	var cepatTelitiTestRequest requests.CepatTelitiRequest
	if err := c.ShouldBindBodyWith(&cepatTelitiTestRequest, binding.JSON); err != nil {
		helpers.ValidationErrorResponse(c, err)
		return
	}

	cepatTelitiAnswer := models.CepatTelitiTestAnswer{
		IDPelamar:  cepatTelitiTestRequest.IDPelamar,
		IDLowongan: cepatTelitiTestRequest.IDLowongan,
		IDUjian:    cepatTelitiTestRequest.IDUjian,
		NoSoal:     cepatTelitiTestRequest.NoSoal,
		Jawaban:    cepatTelitiTestRequest.Jawaban,
	}
	err := configs.DB.Where(
		models.CepatTelitiTestAnswer{
			IDPelamar:  cepatTelitiAnswer.IDPelamar,
			IDLowongan: cepatTelitiAnswer.IDLowongan,
			IDUjian:    cepatTelitiAnswer.IDUjian,
			NoSoal:     cepatTelitiAnswer.NoSoal,
		}).Assign(
		models.CepatTelitiTestAnswer{Jawaban: cepatTelitiAnswer.Jawaban},
	).FirstOrCreate(&cepatTelitiAnswer).Error
	if err != nil {
		helpers.ErrorResponse(c, 400, c.Error(err))
		return
	}
	helpers.SuccessResponse(c, 200, "Data berhasil dimasukkan")
}
