package handlers

import (
	"api.choise/configs"
	"api.choise/internals/helpers"
	"api.choise/internals/models"
	"api.choise/internals/requests"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
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

func DISCPostAnswer(c *gin.Context) {
	var discTestRequest requests.DISCTestRequest
	if err := c.ShouldBindBodyWith(&discTestRequest, binding.JSON); err != nil {
		helpers.ValidationErrorResponse(c, err)
		return
	}

	discAnswer := models.DISCTestAnswer{
		IDPelamar:  discTestRequest.IDPelamar,
		IDLowongan: discTestRequest.IDLowongan,
		IDUjian:    discTestRequest.IDUjian,
		NoSoal:     discTestRequest.NoSoal,
		Jawaban:    discTestRequest.Jawaban,
		Jawaban2:   discTestRequest.Jawaban2,
	}
	err := configs.DB.Where(
		models.DISCTestAnswer{
			IDPelamar:  discAnswer.IDPelamar,
			IDLowongan: discAnswer.IDLowongan,
			IDUjian:    discAnswer.IDUjian,
			NoSoal:     discAnswer.NoSoal,
		}).Assign(
		models.DISCTestAnswer{Jawaban: discAnswer.Jawaban, Jawaban2: discAnswer.Jawaban2},
	).FirstOrCreate(&discAnswer).Error
	if err != nil {
		helpers.ErrorResponse(c, 400, c.Error(err))
		return
	}
	helpers.SuccessResponse(c, 200, "Data berhasil dimasukkan")
}

func DISCAnswerByExamNumber(c *gin.Context) {
	var discAnswer models.DISCTestAnswer
	nomor_soal := c.Query("nomor_soal")
	id_pelamar := c.Query("id_pelamar")
	err := configs.DB.Where("no_soal = ? AND id_pelamar = ?", nomor_soal, id_pelamar).First(&discAnswer).Error
	if err != nil {
		helpers.ErrorResponse(c, 400, err)
		return
	}
	helpers.SuccessResponse(c, 200, discAnswer)
}
