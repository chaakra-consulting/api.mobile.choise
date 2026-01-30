package main

import (
	"fmt"

	"api.choise/configs"
	"api.choise/internals/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	configs.InitDB()
	router := gin.Default()
	routes.ApiRoutes(router)

	err := router.Run(":8080")
	// err := router.Run("192.168.100.59:8080")
	if err != nil {
		fmt.Println("Failed to start server:", err)
	}
}
