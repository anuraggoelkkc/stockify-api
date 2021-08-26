package handlers

import (
	"fmt"
	"log"
	"github.com/gin-gonic/gin"
	"net/http"
	_struct "stockify-api/src/struct"
	"stockify-api/src/support_packs/firestore"
)

type AlertHandler struct{}

func (a AlertHandler) AddAlert(c *gin.Context) {
	fmt.Fprintf(c.Writer, "Welcome to the AddAlert!")
	fmt.Println("Endpoint Hit: AddAlert")

	var alert _struct.Alert
	if c.ShouldBind(&alert) == nil {
		log.Println(alert.Instrument_ID)
		log.Println(alert.User_ID)

		err := firestore.AddAlert(alert)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusOK, gin.H{"message": "Alert successfully added"})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid alert object passed"})
	}
}

func (a AlertHandler) RemoveAlert(c *gin.Context) {
	fmt.Fprintf(c.Writer, "Welcome to the RemoveAlert!")
	fmt.Println("Endpoint Hit: RemoveAlert")
	c.String(http.StatusOK, "Working!")
}

func (a AlertHandler) AlertList(c *gin.Context) {
	fmt.Fprintf(c.Writer, "Welcome to the AlertList!")
	fmt.Println("Endpoint Hit: AlertList")
	userID := c.Param("userID")
	list, err := firestore.FetchAlerts(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"alerts": list})
	}
}