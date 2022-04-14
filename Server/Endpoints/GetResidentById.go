package endpoints

import (
	domain "Server/Domain"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetResidentById(c *gin.Context) {
	id := c.Param("id")
	resident := domain.GetResidentById(id)
	fmt.Println(resident)
	c.IndentedJSON(http.StatusOK, resident)
}
