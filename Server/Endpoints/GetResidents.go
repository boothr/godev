package Endpoints

import (
	"godev/Domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListResidents(c *gin.Context) {
	residents := Domain.GetResidents()
	c.IndentedJSON(http.StatusOK, residents)
}
