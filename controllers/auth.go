package controllers

import (
	"fmt"
	"net/http"
	"resume/models"
	"resume/utils/token"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type RegisterInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Name     string `json:"name" binding:"required"`
}

type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type ChangePasswordInput struct {
	Current   string `json:"current" binding:"required"`
	New       string `json:"new" binding:"required"`
	RetypeNew string `json:"retype_new" binding:"required"`
}

// RegisterPostHandler godoc
// @Description Register for Authentication as Demo App
// @Tags Register
// @Id RegisterPostHandler
// @Produce json
// @Param user body RegisterInput true "Request username, password and name"
// @Success	201 {object} models.Response "OK"
// @Failure 400 {object} models.Response "Bad Request"
// @Failure 422 {object} models.Response "Unprocessable Entity"
// @Router  /api/v1/demo/register [post]
func RegisterPostHandler(c *gin.Context) {

	var input RegisterInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusUnprocessableEntity, models.Response{Code: "422", Message: err.Error()})
		return
	}

	data, err := models.CreateUser(input.Username, input.Password, input.Name)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{Code: "400", Message: err.Error()})
		return
	}
	c.JSON(http.StatusCreated, models.Response{Code: "201", Message: "success", Response: data})
}

// LoginPostHandler godoc
// @Description Login as Demo App
// @Tags Authentication
// @Id LoginPostHandler
// @Produce json
// @Param authen body LoginInput true "Request username, password and return access token(JWT)"
// @Success	200 {object} models.Response "OK"
// @Failure 401 {object} models.Response "Unauthorized"
// @Failure 422 {object} models.Response "Unprocessable Entity"
// @Router  /api/v1/demo/login [post]
func LoginPostHandler(c *gin.Context) {

	var input LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusUnprocessableEntity, models.Response{Code: "422", Message: err.Error()})
		return
	}

	token, err := models.Login(input.Username, input.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, models.Response{Code: "401", Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, models.Response{Code: "200", Message: "welcome to demo app", Response: token})
}

// ProfileGetHandler godoc
// @Description My Profile
// @Tags Authentication
// @Id ProfileGetHandler
// @Produce json
// @Success	200 {object} models.Response "OK"
// @Failure 400 {object} models.Response "Bad Request"
// @Failure 401 {object} models.Response "Unauthorized"
// @Security Bearer
// @Router  /api/v1/demo/user/profile [get]
func ProfileGetHandler(c *gin.Context) {
	tokenAuth, err := token.ExtractTokenMetadataApi(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, models.Response{Code: "401", Message: err.Error()})
		return
	}

	user_id, err := models.FetchAuth(tokenAuth)
	if err != nil {
		c.JSON(http.StatusUnauthorized, models.Response{Code: "401", Message: err.Error()})
		return
	}

	data, err := models.Profile(user_id)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{Code: "400", Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, models.Response{Code: "200", Message: "success", Response: data})
}

// PasswordPatchHandler godoc
// @Description reset password as Demo App
// @Tags Authentication
// @Id PasswordPatchHandler
// @Produce json
// @Param password body ChangePasswordInput true "Reset password"
// @Success	200 {object} models.Response "Ok"
// @Failure 400 {object} models.Response "Bad Request"
// @Failure 401 {object} models.Response "Unauthorized"
// @Failure 422 {object} models.Response "Unprocessable Entity"
// @Security Bearer
// @Router  /api/v1/demo/user/profile [patch]
func PasswordPatchHandler(c *gin.Context) {
	var input ChangePasswordInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusUnprocessableEntity, models.Response{Code: "422", Message: err.Error()})
		return
	}

	if input.New != input.RetypeNew {
		err := fmt.Errorf("passwords do not match")
		log.Error(err)
		c.JSON(http.StatusBadRequest, models.Response{Code: "400", Message: err.Error()})
		return
	}

	tokenAuth, err := token.ExtractTokenMetadataApi(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, models.Response{Code: "401", Message: err.Error()})
		return
	}

	user_id, err := models.FetchAuth(tokenAuth)
	if err != nil {
		c.JSON(http.StatusUnauthorized, models.Response{Code: "401", Message: err.Error()})
		return
	}

	data, err := models.ChangePassword(user_id, input.Current, input.New)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{Code: "400", Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, models.Response{Code: "200", Message: "success", Response: data})
}

// LogoutDelHandler godoc
// @Description Logout and delete access token
// @Tags Authentication
// @Id LogoutDelHandler
// @Produce json
// @Success	200 {object} models.Response "OK"
// @Failure 400 {object} models.Response "Bad Request"
// @Failure 401 {object} models.Response "Unauthorized"
// @Security Bearer
// @Router  /api/v1/demo/user/logout [delete]
func LogoutDelHandler(c *gin.Context) {
	tokenAuth, err := token.ExtractTokenMetadataApi(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, models.Response{Code: "401", Message: err.Error()})
		return
	}
	_, err = models.DeleteToken(tokenAuth)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{Code: "400", Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, models.Response{Code: "200", Message: "sucess"})
}
