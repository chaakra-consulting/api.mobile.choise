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
	var msdtAnswer []models.MSDTTestAnswer
	id_lowongan := c.Query("id_lowongan")
	id_pelamar := c.Query("id_pelamar")
	err := configs.DB.Where("id_lowongan = ? AND id_pelamar = ?", id_lowongan, id_pelamar).Find(&msdtAnswer).Error
	if err != nil {
		helpers.ErrorResponse(c, 400, err)
		return
	}
	helpers.SuccessResponse(c, 200, msdtAnswer)
}

func MSDTPostAnswer(c *gin.Context) {
	var msdtTestRequest requests.MSDTTestRequest
	if err := c.ShouldBindBodyWith(&msdtTestRequest, binding.JSON); err != nil {
		helpers.ValidationErrorResponse(c, c.Error(err))
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
