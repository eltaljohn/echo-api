package handler

import (
	"github.com/eltaljohn/echo-api/middleware"
	"github.com/labstack/echo/v4"
)

func RoutePerson(e *echo.Echo, storage Storage) {
	h := newPerson(storage)

	people := e.Group("/v1/persons")
	people.Use(middleware.Authentication)

	people.POST("", h.create)
	people.GET("", h.getAll)
	people.PUT("/:id", h.update)
	people.GET("/:id", h.getByID)
	people.DELETE("/:id", h.delete)
}

func RouteLogin(e *echo.Echo, storage Storage) {
	h := newLogin(storage)

	e.POST("/v1/login", h.login)
}
