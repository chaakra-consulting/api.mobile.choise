package routes

import (
	"api.choise/internals/handlers"
	"github.com/gin-gonic/gin"
)

func ApiRoutes(r *gin.Engine) {
	r.GET("/ping", handlers.PingHandler)

	//Holland Test Routes
	r.GET("/holland-questions", handlers.HollandTestQuestion)
	r.POST("/holland/send", handlers.HollandPostAnswer)

	//CFIT Test Routes
	r.GET("/cfit-questions/:sub_tes", handlers.CFITQuestion)
	r.GET("/cfit-questions/practice", handlers.CFITQuestionPractice)
	r.POST("/cfit-questions/send", handlers.CFITPostAnswer)
}
