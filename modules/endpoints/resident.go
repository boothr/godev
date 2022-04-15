package endpoints

import (
	"fmt"
	"net/http"

	"github.com/boothr/godev/modules/shared"

	"github.com/gin-gonic/gin"
)

func GetById(c *gin.Context) {
	id := c.Param("id")
	resident := shared.GetResidentById(id)
	fmt.Println(resident)
	c.IndentedJSON(http.StatusOK, resident)
}

func Get(c *gin.Context) {
	residents := shared.GetResidents()
	c.IndentedJSON(http.StatusOK, residents)
}
