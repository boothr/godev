package server

import (
	"github.com/boothr/godev/residents/endpoints"
	"github.com/gin-gonic/gin"
)

func Start() {
	router := gin.Default()
	router.GET("/residents", endpoints.Get)
	router.GET("/residents/:id", endpoints.GetById)
	router.POST("/residents", endpoints.Post)

	router.Run("localhost:8080")
}
