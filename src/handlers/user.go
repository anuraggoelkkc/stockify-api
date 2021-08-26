package handlers

import (
	"fmt"
	"log"
	"github.com/gin-gonic/gin"
	"net/http"
	_struct "stockify-api/src/struct"
	"stockify-api/src/support_packs/firestore"
)

type UserHandler struct{}

func (u UserHandler) AddUser(c *gin.Context) {
	fmt.Fprintf(c.Writer, "Welcome to the AddUser!")
	fmt.Println("Endpoint Hit: AddUser")

	var user _struct.User
	if c.ShouldBind(&user) == nil && len(user.Id) > 0 {
		log.Println(user.Id)
		log.Println(user.DeviceID)
		err := firestore.AddOrUpdateUser(user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusOK, gin.H{"message": "User successfully added"})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid user object passed"})
	}
}
