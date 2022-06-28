package controllers

import (
	"net/http"
	"resume/models"

	"github.com/gin-gonic/gin"
)

// GeckoGetHandler godoc
// @Description	Get coins market data as CoinGecko API
// @Tags CoinsGecko
// @Id GeckoGetHandler
// @Produce	json
// @Success	200 {object} models.Response "OK"
// @Router /api/v1/demo/thirdparty/coingecko/coins/markets [get]
func GeckoGetHandler(c *gin.Context) {
	markets, err := models.CoinsMarket()
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{Code: "400", Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, models.Response{Code: "200", Message: "success", Response: markets})
}
