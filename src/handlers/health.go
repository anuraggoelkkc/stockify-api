package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type HealthHandler struct{}

func (h HealthHandler) Status(c *gin.Context) {
	c.String(http.StatusOK, "Working!")
}
