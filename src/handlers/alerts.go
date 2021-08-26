package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	_struct "stockify-api/src/struct"
	"stockify-api/src/support_packs/zerodha"
)

type AlertHandler struct{
	z *zerodha.Zerodha
}

func (a *AlertHandler) AddAlert(c *gin.Context) {
	fmt.Println("Endpoint Hit: AddAlert")

	var alert _struct.Alert
	if c.ShouldBind(&alert) == nil || len(alert.UserId) > 0 || len(alert.Id) <= 0 {
		log.Println(alert.InstrumentId)
		log.Println(alert.UserId)

		err := a.z.AddAlert(alert)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusOK, gin.H{"message": "Alert successfully added"})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid alert object passed"})
	}
}

func (a *AlertHandler) RemoveAlert(c *gin.Context) {
	fmt.Println("Endpoint Hit: RemoveAlert")

	alertId := c.Param("alertID")
	if len(alertId) > 0 {
		err := a.z.RemoveAlert(alertId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusOK, gin.H{"message": "Alert successfully removed"})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid alert object passed"})
	}
}

func (a *AlertHandler) AlertList(c *gin.Context) {
	fmt.Fprintf(c.Writer, "Welcome to the AlertList!")
	fmt.Println("Endpoint Hit: AlertList")
	userID := c.Param("userID")
	list, err := a.z.FetchAlerts(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"alerts": list})
	}
}

func NewAlertHandler() *AlertHandler {
	return &AlertHandler{
		z : zerodha.NewZerodha(),
	}
}