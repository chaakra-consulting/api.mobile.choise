package handlers

import (
	"errors"
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

	err := configs.DB.Where("subtes = ? AND type_soal = 'Ujian'", sub_test).Find(&cfitQuestions).Error
	if err != nil {
		helpers.ErrorResponse(c, 400, err)
		return
	}
	helpers.SuccessResponse(c, 200, cfitQuestions)
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

	//declare struct to bind JSON request
	var cfitTestRequest requests.CFITTestRequest
	if err := c.ShouldBindBodyWith(&cfitTestRequest, binding.JSON); err != nil {
		helpers.ValidationErrorResponse(c, err)
		return
	}

	if subtes == 2 {
		// Create a new CFITTestAnswerExam record
		cfitAnswer := requests.CFITTestRequest{
			IDPelamar:     cfitTestRequest.IDPelamar,
			IDLowongan:    cfitTestRequest.IDLowongan,
			IDUjian:       cfitTestRequest.IDUjian,
			Subtes:        cfitTestRequest.Subtes,
			NomorSoal:     cfitTestRequest.NomorSoal,
			Jawaban:       cfitTestRequest.Jawaban,
			KunciJawaban:  cfitTestRequest.KunciJawaban,
			KunciJawaban2: cfitTestRequest.KunciJawaban2,
		}

		// err := configs.DB.Create(&cfitAnswer).Error
		// if err != nil {
		// 	helpers.ErrorResponse(c, 400, err)
		// 	return
		// }
		helpers.SuccessResponse(c, 200, cfitAnswer)
	} else {
		// Create a new CFITTestAnswer record
		cfitAnswer := requests.CFITTestRequest{
			IDPelamar:    cfitTestRequest.IDPelamar,
			IDLowongan:   cfitTestRequest.IDLowongan,
			IDUjian:      cfitTestRequest.IDUjian,
			Subtes:       cfitTestRequest.Subtes,
			NomorSoal:    cfitTestRequest.NomorSoal,
			Jawaban:      cfitTestRequest.Jawaban,
			KunciJawaban: cfitTestRequest.KunciJawaban,
		}

		// err := configs.DB.Create(&cfitAnswer).Error
		// if err != nil {
		// 	helpers.ErrorResponse(c, 400, err)
		// 	return
		// }
		helpers.SuccessResponse(c, 200, cfitAnswer)
	}

}
