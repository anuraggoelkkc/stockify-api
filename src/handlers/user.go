package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserHandler struct{}

func (u UserHandler) AddUser(c *gin.Context) {
	fmt.Fprintf(c.Writer, "Welcome to the AddUser!")
	fmt.Println("Endpoint Hit: AddUser")
	c.String(http.StatusOK, "Working!")
}
