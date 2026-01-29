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
	err := configs.DB.Find(&cepatTelitiQuestions).Error
	if err != nil {
		helpers.ErrorResponse(c, 400, err)
		return
	}
	helpers.SuccessResponse(c, 200, cepatTelitiQuestions)
}

func CepatTelitiAnswerByExamNumber(c *gin.Context) {
	var cepatTelitiAnswer models.CepatTelitiTestAnswer
	nomor_soal := c.Query("nomor_soal")
	id_pelamar := c.Query("id_pelamar")
	err := configs.DB.Where("no_soal = ? AND id_pelamar = ?", nomor_soal, id_pelamar).First(&cepatTelitiAnswer).Error
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
