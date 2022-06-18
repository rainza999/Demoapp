package controllers

import (
	"net/http"
	"resume/models"
	"resume/utils/token"
	"strconv"

	"github.com/gin-gonic/gin"
)

// AccountGetHandler godoc
// @Description list users
// @Tags Account
// @Id AccountGetHandler
// @Produce json
// @Param st query string  false  "name search by create start date (yyyy-mm-dd)"  Format(datetime)
// @Param et query string  false  "name search by create end date (yyyy-mm-dd)"  Format(datetime)
// @Param sort query string  false  "name sort asc or desc"  Format(string)
// @Param page query int  false  "name search by page"  Format(int)
// @Param limit query int  false  "name search by limit"  Format(int)
// @Success	200 {object} models.ResponsePage "OK"
// @Failure 400 {object} models.ResponsePage "Bad Request"
// @Failure 401 {object} models.ResponsePage "Unauthorized"
// @Security Bearer
// @Router  /api/v1/demo/account [get]
func AccountGetHandler(c *gin.Context) {
	_, err := token.ExtractTokenMetadataApi(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, models.ResponsePage{Code: "401", Message: err.Error()})
		return
	}

	limit := 0
	startdate := c.Query("st")
	enddate := c.Query("et")

	sort := c.Query("sort")
	if sort == "" {
		sort = "desc"
	}

	lmt, _ := strconv.Atoi(c.Query("limit"))
	if lmt == 0 {
		limit = limit + 10
	} else {
		limit = limit + lmt
	}

	pg, _ := strconv.Atoi(c.Query("page"))

	scheme := "http"
	if c.Request.TLS != nil {
		scheme = "https"
	}
	url := scheme + "://" + c.Request.Host + c.FullPath()

	data, pagination, err := models.GetListUsers(url, startdate, enddate, sort, pg, limit)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponsePage{Code: "400", Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, models.ResponsePage{Code: "200", Message: "success", Response: data, Pagination: pagination})
}
