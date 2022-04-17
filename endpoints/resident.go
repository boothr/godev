package endpoints

import (
	"fmt"
	"net/http"

	"github.com/boothr/godev/residents/shared"

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

func Post(c *gin.Context) {
	var newResident = shared.GetNewResident()
	if err := c.BindJSON(&newResident); err != nil {
		return
	}

	if shared.AddResident(newResident) {
		c.IndentedJSON(http.StatusCreated, newResident)
	}

	// Call BindJSON to bind the received JSON to
	// newAlbum.

}
