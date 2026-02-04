package handlers

import (
	"api.choise/configs"
	"api.choise/internals/helpers"
	"github.com/gin-gonic/gin"
)

func ExamListHandler(c *gin.Context) {

	//declare exam maps
	var cfitExam map[string]interface{}
	var hollandExam map[string]interface{}
	var discExam map[string]interface{}
	var msdtExam map[string]interface{}
	var cepatTelitiExam map[string]interface{}

	var data []interface{}

	//fetch exam data
	err := configs.DB.Table("tb_ujian").Find(&cfitExam).Error
	err = configs.DB.Table("tb_ujian_holland").Find(&hollandExam).Error
	err = configs.DB.Table("tb_ujian_disc").Find(&discExam).Error
	err = configs.DB.Table("tb_ujian_msdt").Find(&msdtExam).Error
	err = configs.DB.Table("tb_ujian_cepat").Find(&cepatTelitiExam).Error

	if err != nil {
		helpers.ErrorResponse(c, 400, err)
		return
	}

	//combine exam data
	// data = make([]map[string]interface{}, 0)
	// data = append(data, map[string]interface{}{

	// "cfit":         cfitExam,
	// "holland":      hollandExam,
	// "disc":         discExam,
	// "msdt":         msdtExam,
	// "cepat_teliti": cepatTelitiExam,
	data = append(data,
		cfitExam,
		hollandExam,
		cepatTelitiExam,
		discExam,
		msdtExam,
	)

	helpers.SuccessResponse(c, 200, data)
}
