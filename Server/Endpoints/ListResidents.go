package endpoints

import (
	domain "Server/Domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListResidents(c *gin.Context) {
	residents := domain.GetResidents()
	c.IndentedJSON(http.StatusOK, residents)
}
