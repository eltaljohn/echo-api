package handler

import (
	"errors"
	"github.com/eltaljohn/echo-api/model"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type person struct {
	storage Storage
}

func newPerson(storage Storage) person {
	return person{storage}
}

func (p *person) create(c echo.Context) error {
	data := model.Person{}
	err := c.Bind(&data)
	if err != nil {
		response := newResponse(Error, "body malformed", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	err = p.storage.Create(&data)
	if err != nil {
		response := newResponse(Error, "error creating person", nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := newResponse(Message, "person created successfully", nil)
	return c.JSON(http.StatusCreated, response)
}

func (p *person) getAll(c echo.Context) error {
	data, err := p.storage.GetAll()
	if err != nil {
		response := newResponse(Error, "error getting all people", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	response := newResponse(Message, "Ok", data)
	return c.JSON(http.StatusOK, response)
}

func (p *person) getByID(c echo.Context) error {
	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response := newResponse(Error, "the id must be int", nil)
		return c.JSON(http.StatusBadRequest, response)
	}
	data, err := p.storage.GetByID(ID)
	if err != nil {
		response := newResponse(Error, "error getting the person", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	response := newResponse(Message, "Ok", data)
	return c.JSON(http.StatusOK, response)
}

func (p *person) update(c echo.Context) error {
	id := c.Param("id")
	ID, err := strconv.Atoi(id)
	if err != nil {
		response := newResponse(Error, "the id must be int", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	data := model.Person{}
	err = c.Bind(&data)
	if err != nil {
		response := newResponse(Error, "body malformed", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	err = p.storage.Update(ID, &data)
	if err != nil {
		response := newResponse(Error, "error updating person", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	response := newResponse(Message, "person updated successfully", nil)
	return c.JSON(http.StatusOK, response)
}

func (p *person) delete(c echo.Context) error {
	id := c.Param("id")
	ID, err := strconv.Atoi(id)
	if err != nil {
		response := newResponse(Error, "the id must be int", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	err = p.storage.Delete(ID)
	if errors.Is(err, model.ErrIDPersonDoesNotExist) {
		response := newResponse(Error, "user id does not exist", nil)
		return c.JSON(http.StatusBadRequest, response)
	}
	if err != nil {
		response := newResponse(Error, "error deleting person", nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := newResponse(Message, "Ok", nil)
	return c.JSON(http.StatusOK, response)
}
