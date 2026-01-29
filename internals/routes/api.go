package routes

import (
	"api.choise/internals/handlers"
	"api.choise/internals/middlewares"
	"github.com/gin-gonic/gin"
)

func ApiRoutes(r *gin.Engine) {
	r.GET("/ping", handlers.PingHandler)

	//Login Route
	r.POST("/login", handlers.Login)

	//Job Vacancy Routes
	r.GET("/job-vacancy", middlewares.AuthMiddleware(), handlers.GetAvailableJobVacancy)

	//Holland Test Routes
	r.GET("/holland-questions", handlers.HollandTestQuestion)
	r.POST("/holland/send", handlers.HollandPostAnswer)

	//CFIT Test Routes
	r.GET("/cfit-questions/:sub_tes", handlers.CFITQuestion)
	r.GET("/cfit-questions/by-exam-number/:sub_tes", handlers.CFITAnswerByExamNumber)
	r.GET("/cfit-questions/practice", handlers.CFITQuestionPractice)
	r.POST("/cfit/send", middlewares.AuthMiddleware(), handlers.CFITPostAnswer)

	//DISC Test Routes
	r.GET("/disc-questions", handlers.DISCQuestion)
	r.POST("/disc/send", middlewares.AuthMiddleware(), handlers.DISCPostAnswer)
	r.GET("/disc-questions/by-exam-number", handlers.DISCAnswerByExamNumber)

	//MSDT Test Routes
	r.GET("/msdt-questions", handlers.MSDTQuestion)
	r.GET("/msdt-questions/by-exam-number", handlers.MSDTAnswerByExamNumber)
	r.POST("/msdt/send", handlers.MSDTPostAnswer)

	//Cepat Teliti Test Routes
	r.GET("/cepat-teliti-questions", handlers.CepatTelitiQuestion)
	r.GET("/cepat-teliti-questions/by-exam-number", handlers.CepatTelitiAnswerByExamNumber)
	r.POST("/cepat-teliti/send", middlewares.AuthMiddleware(), handlers.CepatTelitiPostAnswer)

	//Exam List Route
	r.GET("/exam-list", handlers.ExamListHandler)
}
