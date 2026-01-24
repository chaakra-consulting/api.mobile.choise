package handlers

import (
	"api.choise/internals/helpers"
	"github.com/gin-gonic/gin"
)

func PingHandler(c *gin.Context) {

	helpers.JSONResponse(c, 200, gin.H{
		"message": "pong",
	})
}
