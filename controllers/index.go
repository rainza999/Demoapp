package controllers

import (
	"net/http"
	"resume/models"

	"github.com/gin-gonic/gin"
)

type ContactInput struct {
	Name    string `json:"name" binding:"required"`
	Email   string `json:"email" binding:"required"`
	Subject string `json:"subject"`
	Msg     string `json:"msg"`
}

func IndexGetHandler(c *gin.Context) {
	scheme := "http"
	if c.Request.TLS != nil {
		scheme = "https"
	}

	host := scheme + "://" + c.Request.Host
	healthcheck := models.HealthCheck(host)
	skills := models.SkillsInfo()
	tools := models.ToolsInfo()
	personal := models.PersonalInfo(host)
	experiences := models.ExperienceInfo()
	education := models.EducationInfo()
	about := models.AboutInfo()
	interest := models.InterestInfo()

	data := gin.H{
		"powerby":     healthcheck,
		"skills":      skills,
		"tools":       tools,
		"personal":    personal,
		"experiences": experiences,
		"education":   education,
		"about":       about,
		"interest":    interest,
	}

	c.HTML(http.StatusOK, "index.html", data)
}

// PingGetHandler godoc
// @Description	Check API server status
// @Tags Ping
// @Id PingGetHandler
// @Produce	json
// @Success	200 {object} models.Response "OK"
// @Router  /api/v1 [get]
func PingGetHandler(c *gin.Context) {
	scheme := "http"
	if c.Request.TLS != nil {
		scheme = "https"
	}
	host := scheme + "://" + c.Request.Host
	res := make(map[string]interface{})

	healthcheck := models.HealthCheck(host)
	res["powerby"] = healthcheck

	c.JSON(http.StatusOK, models.Response{Code: "200", Message: "success", Response: res})
}

// AboutGetHandler godoc
// @Description	Get Personal, Interests and References
// @Tags About
// @Id AboutGetHandler
// @Produce	json
// @Success	200 {object} models.Response "OK"
// @Router  /api/v1/about [get]
func AboutGetHandler(c *gin.Context) {
	scheme := "http"
	if c.Request.TLS != nil {
		scheme = "https"
	}
	host := scheme + "://" + c.Request.Host

	res := make(map[string]interface{})

	about := models.AboutInfo()
	interest := models.InterestInfo()
	personal := models.PersonalInfo(host)

	res["personal"] = personal
	res["about"] = about
	res["interest"] = interest

	c.JSON(http.StatusOK, models.Response{Code: "200", Message: "success", Response: res})
}

// ExperiencesGetHandler godoc
// @Description	Get Education, Experience and Portfolio
// @Tags Experience
// @Id ExperiencesGetHandler
// @Produce	json
// @Success	200 {object} models.Response "OK"
// @Router  /api/v1/experience [get]
func ExperiencesGetHandler(c *gin.Context) {

	res := make(map[string]interface{})
	education := models.EducationInfo()
	experiences := models.ExperienceInfo()

	res["education"] = education
	res["experiences"] = experiences

	c.JSON(http.StatusOK, models.Response{Code: "200", Message: "success", Response: res})
}

// SkillGetHandler godoc
// @Description	Get Technical Skills and Tools
// @Tags Skills
// @Id SkillGetHandler
// @Produce	json
// @Success	200 {object} models.Response "OK"
// @Router  /api/v1/skill [get]
func SkillGetHandler(c *gin.Context) {
	res := make(map[string]interface{})
	skill := models.SkillsInfo()
	tool := models.ToolsInfo()

	res["skill"] = skill
	res["tool"] = tool

	c.JSON(http.StatusOK, models.Response{Code: "200", Message: "success", Response: res})
}

func ContactPostHandler(c *gin.Context) {

	name := c.PostForm("name")
	email := c.PostForm("email")
	subject := c.PostForm("subject")
	msg := c.PostForm("message")

	models.CreateContact(name, email, subject, msg)

	host := c.Request.Host
	healthcheck := models.HealthCheck(host)
	skills := models.SkillsInfo()
	tools := models.ToolsInfo()
	personal := models.PersonalInfo(host)
	experiences := models.ExperienceInfo()
	education := models.EducationInfo()
	about := models.AboutInfo()
	interest := models.InterestInfo()

	data := gin.H{
		"powerby":     healthcheck,
		"skills":      skills,
		"tools":       tools,
		"personal":    personal,
		"experiences": experiences,
		"education":   education,
		"about":       about,
		"interest":    interest,
	}

	c.HTML(http.StatusOK, "index.html", data)
}

// ContactPostHandlerApi godoc
// @Description	I will contact you soon
// @Tags Contact
// @Id ContactPostHandlerApi
// @Produce	json
// @Param contact body ContactInput true "Contact me"
// @Success	200 {object} models.Response "OK"
// @Failure 400 {object} models.Response "Bad Request"
// @Failure 422 {object} models.Response "Unprocessable Entity"
// @Router  /api/v1/contact [post]
func ContactPostHandlerApi(c *gin.Context) {
	var input ContactInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusUnprocessableEntity, models.Response{Code: "422", Message: err.Error()})
		return
	}
	err := models.CreateContact(input.Name, input.Email, input.Subject, input.Msg)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{Code: "400", Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, models.Response{Code: "200", Message: "success"})
}
