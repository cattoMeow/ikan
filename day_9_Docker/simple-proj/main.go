package main

import (
	// "efishery-day6-tugas-lagi/controller"
	"log"
	"os"
	"simple-proj/controller"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	file, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/customer", controller.GetAllUsers)
	e.POST("/customer", controller.CreateUser)
	e.GET("/customer/:id", controller.GetUser)
	e.PUT("/customer/:id", controller.UpdateUser)
	e.DELETE("/customer/:id", controller.DeleteUser)

	// Start server
	e.Logger.Fatal(e.Start(":5002"))
}
