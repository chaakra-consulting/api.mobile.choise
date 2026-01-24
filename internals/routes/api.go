package routes

import (
	"api.choise/internals/handlers"
	"github.com/gin-gonic/gin"
)

func ApiRoutes(r *gin.Engine) {
	r.GET("/ping", handlers.PingHandler)
}
