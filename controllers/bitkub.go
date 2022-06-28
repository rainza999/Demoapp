package controllers

import (
	"net/http"
	"resume/models"

	"github.com/gin-gonic/gin"
)

// BKTickerGetHandler godoc
// @Description	Get ticker information as Bitkub
// @Tags Bitkub
// @Id BKTickerGetHandler
// @Produce	json
// @Param sym query string false "sym string The symbol (optional)" Format(string)
// @Success	200 {object} models.Response "OK"
// @Router /api/v1/demo/thirdparty/bitkub/market/ticker [get]
func BKTickerGetHandler(c *gin.Context) {
	symbol := c.Query("sym")
	resp, err := models.TickerBitkub(symbol)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{Code: "400", Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, models.Response{Code: "200", Message: "success", Response: resp})
}
