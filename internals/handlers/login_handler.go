package handlers

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"net/http"
	"strconv"

	"api.choise/configs"
	"api.choise/internals/helpers"
	"api.choise/internals/models"
	"api.choise/internals/requests"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func Login(c *gin.Context) {
	var input requests.LoginRequest
	var applicantUser models.ApplicantUser

	if err := c.ShouldBindBodyWith(&input, binding.JSON); err != nil {
		helpers.ValidationErrorResponse(c, err)
		return
	}
	hash := md5.Sum([]byte(input.Password))
	passwordField := hex.EncodeToString(hash[:])

	err := configs.DB.Unscoped().Where("username = ? AND password = ?", input.Username, passwordField).Find(&applicantUser).Error
	if err != nil {
		helpers.ErrorResponse(c, http.StatusUnauthorized, c.Error(errors.New("invalid username or password")))
		return
	}

	token, err := helpers.GenerateJWT(strconv.Itoa(int(applicantUser.IDPelamar)))
	if err != nil {
		helpers.ErrorResponse(c, http.StatusInternalServerError, c.Error(err))
		return
	}

	data := map[string]interface{}{
		"user":  applicantUser,
		"token": token,
	}

	helpers.SuccessResponse(c, 200, data)
}
