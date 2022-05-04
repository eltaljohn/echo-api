package main

import (
	"github.com/eltaljohn/echo-api/authorization"
	"github.com/eltaljohn/echo-api/handler"
	"github.com/eltaljohn/echo-api/storage"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
)

func main() {
	err := authorization.LoadFiles("certificates/app.rsa", "certificates/app.rsa.pub")
	if err != nil {
		log.Fatalf("Error loading certificates: %v", err)
	}
	store := storage.NewMemory()

	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	handler.RoutePerson(e, &store)
	handler.RouteLogin(e, &store)

	log.Println("Server running in port 8080")
	err = e.Start(":8080")
	if err != nil {
		log.Printf("error to run serve %v", err)
	}
}
