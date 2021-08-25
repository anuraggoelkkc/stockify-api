package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AlertHandler struct{}

func (a AlertHandler) AddAlert(c *gin.Context) {
	fmt.Fprintf(c.Writer, "Welcome to the AddAlert!")
	fmt.Println("Endpoint Hit: AddAlert")
	c.String(http.StatusOK, "Working!")
}

func (a AlertHandler) RemoveAlert(c *gin.Context) {
	fmt.Fprintf(c.Writer, "Welcome to the RemoveAlert!")
	fmt.Println("Endpoint Hit: RemoveAlert")
	c.String(http.StatusOK, "Working!")
}

func (a AlertHandler) AlertList(c *gin.Context) {
	fmt.Fprintf(c.Writer, "Welcome to the AlertList!")
	fmt.Println("Endpoint Hit: AlertList")
	c.String(http.StatusOK, "Working!")
}