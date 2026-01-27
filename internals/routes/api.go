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
	r.POST("/cfit/send", handlers.CFITPostAnswer)

	//DISC Test Routes
	r.GET("/disc-questions", handlers.DISCQuestion)

	//MSDT Test Routes
	r.GET("/msdt-questions", handlers.MSDTQuestion)

	//Cepat Teliti Test Routes
	r.GET("/cepat-teliti-questions", handlers.CepatTelitiQuestion)

	//Exam List Route
	r.GET("/exam-list", handlers.ExamListHandler)
}
