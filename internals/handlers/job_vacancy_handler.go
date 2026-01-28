package handlers

import (
	"net/http"

	"api.choise/configs"
	"api.choise/internals/helpers"
	"api.choise/internals/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetAvailableJobVacancy(c *gin.Context) {
	var jobVacancy []models.JobVacancy
	idPelamar := c.Query("id_pelamar")
	err := configs.DB.Unscoped().Preload("Company").Preload("Applicants", func(db *gorm.DB) *gorm.DB {
		return db.Where("id_pelamar = ?", idPelamar).Limit(1)
	}).Where("status = ?", "tersedia").Find(&jobVacancy).Error
	if err != nil {
		helpers.ErrorResponse(c, http.StatusInternalServerError, c.Error(err))
		return
	}
	helpers.SuccessResponse(c, http.StatusOK, jobVacancy)
}
