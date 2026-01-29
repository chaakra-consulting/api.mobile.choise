package handlers

import (
	"api.choise/configs"
	"api.choise/internals/helpers"
	"api.choise/internals/models"
	"api.choise/internals/requests"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
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

func MSDTAnswerByExamNumber(c *gin.Context) {
	var msdtAnswer models.MSDTTestAnswer
	nomor_soal := c.Query("nomor_soal")
	id_pelamar := c.Query("id_pelamar")
	err := configs.DB.Where("no_soal = ? AND id_pelamar = ?", nomor_soal, id_pelamar).First(&msdtAnswer).Error
	if err != nil {
		helpers.ErrorResponse(c, 400, err)
		return
	}
	helpers.SuccessResponse(c, 200, msdtAnswer)
}

func MSDTPostAnswer(c *gin.Context) {
	var msdtTestRequest requests.MSDTTestRequest
	if err := c.ShouldBindBodyWith(&msdtTestRequest, binding.JSON); err != nil {
		helpers.ValidationErrorResponse(c, err)
		return
	}

	msdtAnswer := models.MSDTTestAnswer{
		IDPelamar:  msdtTestRequest.IDPelamar,
		IDLowongan: msdtTestRequest.IDLowongan,
		IDUjian:    msdtTestRequest.IDUjian,
		NoSoal:     msdtTestRequest.NoSoal,
		Jawaban:    msdtTestRequest.Jawaban,
	}
	err := configs.DB.Where(
		models.MSDTTestAnswer{
			IDPelamar:  msdtAnswer.IDPelamar,
			IDLowongan: msdtAnswer.IDLowongan,
			IDUjian:    msdtAnswer.IDUjian,
			NoSoal:     msdtAnswer.NoSoal,
		}).Assign(
		models.MSDTTestAnswer{Jawaban: msdtAnswer.Jawaban},
	).FirstOrCreate(&msdtAnswer).Error
	if err != nil {
		helpers.ErrorResponse(c, 400, c.Error(err))
		return
	}
	helpers.SuccessResponse(c, 200, "Data berhasil dimasukkan")

}
