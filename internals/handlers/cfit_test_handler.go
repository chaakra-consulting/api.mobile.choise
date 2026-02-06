package handlers

import (
	"errors"
	"fmt"
	"strconv"

	"api.choise/configs"
	"api.choise/internals/helpers"
	"api.choise/internals/models"
	"api.choise/internals/requests"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func CFITQuestion(c *gin.Context) {
	var cfitQuestions []models.CFITTestQuestion
	// nomor_soal := c.Query("nomor_soal")
	sub_test := c.Param("sub_tes")
	if sub_test == "" {
		helpers.ErrorResponse(c, 400, errors.New("sub_test and nomor_soal query parameter is required"))
		return
	}

	err := configs.DB.Where("subtes = ? AND type_soal = 'Ujian'", sub_test).Order("nomor_soal ASC").Find(&cfitQuestions).Error
	if err != nil {
		helpers.ErrorResponse(c, 400, err)
		return
	}
	helpers.SuccessResponse(c, 200, cfitQuestions)
}
func CFITAnswerByExamNumber(c *gin.Context) {
	var cfitAnswer []models.CFITTestAnswer
	// nomor_soal := c.Query("nomor_soal")
	id_pelamar := c.Query("id_pelamar")
	id_lowongan := c.Query("id_lowongan")
	sub_test := c.Param("sub_tes")
	if sub_test == "" {
		helpers.ErrorResponse(c, 400, errors.New("sub_test and nomor_soal query parameter is required"))
		return
	}

	err := configs.DB.Where("subtes = ?  AND id_pelamar = ? and id_lowongan = ?", sub_test, id_pelamar, id_lowongan).Order("nomor_soal ASC").Find(&cfitAnswer).Error
	if err != nil {
		helpers.ErrorResponse(c, 400, c.Error(err))
		return
	}
	helpers.SuccessResponse(c, 200, cfitAnswer)
}

func CFITQuestionPractice(c *gin.Context) {
	var cfitQuestions []models.CFITTestQuestion
	sub_test := c.Query("sub_test")
	if sub_test == "" {
		helpers.ErrorResponse(c, 400, errors.New("sub_test query parameter is required"))
		return
	}
	type_soal := "Contoh"
	err := configs.DB.Where("subtes = ? AND type_soal = ?", sub_test, type_soal).Find(&cfitQuestions).Error
	if err != nil {
		helpers.ErrorResponse(c, 400, err)
		return
	}
	helpers.SuccessResponse(c, 200, cfitQuestions)
}

func CFITPostAnswer(c *gin.Context) {
	subtes, err := strconv.Atoi(c.Query("subtes")) // parse subtes to int
	if err != nil {
		helpers.ErrorResponse(c, 400, c.Error(errors.New("invalid subtes parameter")))
		return
	}
	fmt.Println(subtes)

	//declare struct to bind JSON request
	var cfitTestRequest requests.CFITTestRequest
	if err := c.ShouldBindBodyWith(&cfitTestRequest, binding.JSON); err != nil {
		helpers.ValidationErrorResponse(c, c.Error(err))
		return
	}

	cfitAnswer := models.CFITTestAnswer{
		IDPelamar:     uint(cfitTestRequest.IDPelamar),
		IDLowongan:    uint(cfitTestRequest.IDLowongan),
		IDUjian:       1,
		Subtes:        c.Query("subtes"),
		NomorSoal:     cfitTestRequest.NomorSoal,
		Jawaban:       cfitTestRequest.Jawaban,
		JawabanKunci:  cfitTestRequest.JawabanKunci,
		Jawaban2:      cfitTestRequest.Jawaban2,
		JawabanKunci2: cfitTestRequest.JawabanKunci2,
	}

	err = configs.DB.Where(
		models.CFITTestAnswer{
			IDPelamar:  uint(cfitAnswer.IDPelamar),
			IDLowongan: uint(cfitAnswer.IDLowongan),
			NomorSoal:  cfitAnswer.NomorSoal,
			IDUjian:    1,
			Subtes:     c.Query("subtes"),
		}).Assign(
		models.CFITTestAnswer{
			Jawaban:       cfitTestRequest.Jawaban,
			JawabanKunci:  cfitAnswer.JawabanKunci,
			Jawaban2:      cfitAnswer.Jawaban2,
			JawabanKunci2: cfitAnswer.JawabanKunci2},
	).FirstOrCreate(&cfitAnswer).Error
	if err != nil {
		helpers.ErrorResponse(c, 400, c.Error(err))
		return
	}
	helpers.SuccessResponse(c, 200, "Data berhasil dimasukkan")

}
