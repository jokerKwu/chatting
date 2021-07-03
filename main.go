package main

import (
	"Chatting/api"
	"Chatting/config"
	"Chatting/routes"
	"Chatting/utils"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

var userController *api.UserController


// @title Golang CHATTING REST API
// @version 1.0
// @description This is a sample server server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1
func main() {
	e := echo.New()
	e.HTTPErrorHandler = api.ErrorHandler

	//middleware
	e.Validator = utils.NewValidationUtil()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	config.CORSConfig(e)


	routes.GetUserApiRoutes(e, userController)
	routes.GetSwaggerRoutes(e)


	e.Logger.Fatal(	e.Start(fmt.Sprintf(":%s",config.ServerPort.(string))))
}

// HealthCheck godoc
// @Summary Show the status of server.
// @Description get the status of server.
// @Tags root
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router / [get]
func HealthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": "Server is up and running",
	})
}
