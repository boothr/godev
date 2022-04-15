package main

import (
	endpoints "Server/Endpoints"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/residents", endpoints.ListResidents)
	router.GET("/residents/:id", endpoints.GetResidentById)

	router.Run("localhost:8080")
}
