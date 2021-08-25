package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"stockify-api/src/support_packs/zerodha"
)

type MetadataHandler struct{}

func (m MetadataHandler) ReloadInstruments(c *gin.Context) {
	fmt.Fprintf(c.Writer, "Welcome to the ReloadInstruments!")
	fmt.Println("Endpoint Hit: ReloadInstruments")

	err := zerodha.ReloadInstrumentsInFirebase()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Instrument list is successfully updated"})
	}
}

func (m MetadataHandler) TrendingInstruments(c *gin.Context) {
	fmt.Fprintf(c.Writer, "Welcome to the TrendingList!")
	fmt.Println("Endpoint Hit: TrendingList")
	c.String(http.StatusOK, "Working!")
}
