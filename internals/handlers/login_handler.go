package handlers

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"log"
	"net/http"
	"reflect"
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

	// Bind the JSON request to the input struct
	if err := c.ShouldBindBodyWith(&input, binding.JSON); err != nil {
		helpers.ValidationErrorResponse(c, err)
		return
	}
	hash := md5.Sum([]byte(input.Password))
	passwordField := hex.EncodeToString(hash[:])

	log.Println(passwordField)

	// Find the user
	err := configs.DB.Unscoped().Preload("UserData").Where("username = ? AND password = ?", input.Username, passwordField).Find(&applicantUser).Error
	if err != nil {
		helpers.ErrorResponse(c, http.StatusUnauthorized, c.Error(errors.New("invalid username or password")))
		return
	}

	// check if data is exist
	if reflect.ValueOf(applicantUser).IsZero() {
		helpers.ErrorResponse(c, http.StatusNotFound, c.Error(errors.New("data not found")))
		return
	}

	// Generate JWT Token within user
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

func FindEmail(c *gin.Context) {
	var input requests.FindEmailRequest
	var applicantUser models.ApplicantUser

	// Bind the JSON request to the input struct
	if err := c.ShouldBindBodyWith(&input, binding.JSON); err != nil {
		helpers.ValidationErrorResponse(c, err)
		return
	}

	// Find the user by email
	err := configs.DB.Unscoped().Where("email = ?", input.Email).Find(&applicantUser).Error
	if err != nil {
		helpers.ErrorResponse(c, http.StatusUnauthorized, c.Error(err))

		return
	}

	// Check if the user was found
	if reflect.ValueOf(applicantUser).IsZero() {
		helpers.ErrorResponse(c, http.StatusNotFound, c.Error(errors.New("data not found")))
		return
	}
	log.Println(applicantUser)

	helpers.SuccessResponse(c, http.StatusOK, "Data berhasil ditemukan")

}

func ResetPassword(c *gin.Context) {
	var input requests.ResetPasswordRequest

	// Bind the JSON request to the input struct
	if err := c.ShouldBindBodyWith(&input, binding.JSON); err != nil {
		helpers.ValidationErrorResponse(c, err)
		return
	}

	// Check if password and confirm password match
	if input.Password != input.ConfirmPassword {
		helpers.BadRequestResponse(c, c.Error(errors.New("Password dan konfirmasi password tidak sama")))
		return
	}

	// Hash password
	new_password := md5.Sum([]byte(input.Password))
	passwordField := hex.EncodeToString(new_password[:])

	// Update password
	err := configs.DB.Table("tb_pelamar").Where(models.ApplicantUser{
		Email: input.Email,
	}).Updates(map[string]interface{}{"password": passwordField, "confirm_password": input.Password}).Error
	if err != nil {
		helpers.ErrorResponse(c, http.StatusInternalServerError, c.Error(err))
		return
	}

	helpers.SuccessResponse(c, http.StatusOK, "Password berhasil direset")

}
