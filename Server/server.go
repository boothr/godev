package main

import (
	"github.com/boothr/godev/modules/endpoints"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gi.Default()
	router.GET("/residents" endpoints.Get)
	router.GET("/residents/:id", endpoints.GetById)

  router.Run("localhost:8080")
}
