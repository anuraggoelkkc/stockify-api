package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	_struct "stockify-api/src/struct"
	"stockify-api/src/support_packs/firestore"
)

type UserHandler struct{
	f *firestore.FireStore
}

func (u *UserHandler) AddUser(c *gin.Context) {
	fmt.Println("Endpoint Hit: AddUser")

	var user _struct.User
	if c.ShouldBind(&user) == nil && len(user.Id) > 0 {
		log.Println(user.Id)
		log.Println(user.DeviceID)
		err := u.f.AddOrUpdateUser(user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusOK, gin.H{"message": "User successfully added"})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid user object passed"})
	}
}

func NewUserHandler() *UserHandler {
	return &UserHandler{
		f : firestore.NewFireStore("","","","",""),
	}
}
