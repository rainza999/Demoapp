package main

import (
	_ "resume/docs"
	"resume/models"
	"resume/routes"

	log "github.com/sirupsen/logrus"
)

// @title Demoapp API
// @version 1.0
// @description Demoapp develop a RESTful API with Go 1.18 (Gin framework)

// @contact.name Mr.Navamin Sawasdee
// @contact.email navaminsawasdee@gmail.com

// @license.name Copyright © 2022 By Navamin Sawasdee. All rights reserved.

// @schemes http https

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
func main() {
	models.ConnectMysql()
	models.ConnectRedis()
	models.ConnectMailer()

	r := routes.SetupRouter()

	if err := r.Run(":8080"); err != nil {
		log.Fatal("Unable to start:", err)
	}
}
