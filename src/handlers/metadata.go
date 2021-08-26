package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"stockify-api/src/support_packs/zerodha"
)

type MetadataHandler struct{
	z *zerodha.Zerodha
}

func (m *MetadataHandler) ReloadInstruments(c *gin.Context) {
	fmt.Fprintf(c.Writer, "Welcome to the ReloadInstruments!")
	fmt.Println("Endpoint Hit: ReloadInstruments")

	err := m.z.ReloadInstrumentsInFirebase()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Instrument list is successfully updated"})
	}
}

func (m *MetadataHandler) TrendingInstruments(c *gin.Context) {
	fmt.Fprintf(c.Writer, "Welcome to the TrendingList!")
	fmt.Println("Endpoint Hit: TrendingList")
	c.String(http.StatusOK, "Working!")
}

func (m *MetadataHandler) FetchInstrumentDetails(c *gin.Context) {
	fmt.Fprintf(c.Writer, "Welcome to the FetchInstrumentDetails!")
	fmt.Println("Endpoint Hit: FetchInstrumentDetails")
	exchange := c.Param("exchange")
	symbol := c.Param("symbol")
	val, err := m.z.FetchInstrumentDetails(exchange,symbol)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": val})
	}
}

func NewMetadataHandler() *MetadataHandler {
	return &MetadataHandler{
		z : zerodha.NewZerodha(),
	}
}
