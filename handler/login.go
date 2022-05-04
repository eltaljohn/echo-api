package handler

import (
	"github.com/eltaljohn/echo-api/authorization"
	"github.com/eltaljohn/echo-api/model"
	"github.com/labstack/echo/v4"
	"net/http"
)

type login struct {
	storage Storage
}

func newLogin(s Storage) login {
	return login{s}
}

func (l *login) login(c echo.Context) error {

	data := model.Login{}
	err := c.Bind(&data)
	if err != nil {
		resp := newResponse(Error, "no valid structure", nil)
		return c.JSON(http.StatusBadRequest, resp)
	}
	if !isLoginValid(&data) {
		resp := newResponse(Error, "no valid credentials", nil)
		return c.JSON(http.StatusBadRequest, resp)
	}

	token, err := authorization.GenerateToken(&data)
	if err != nil {
		resp := newResponse(Error, "could not generate token", nil)
		return c.JSON(http.StatusInternalServerError, resp)
	}

	dataToken := map[string]string{"token": token}
	response := newResponse(Message, "Ok", dataToken)
	return c.JSON(http.StatusOK, response)
}

func isLoginValid(data *model.Login) bool {
	return data.Email == "contact@ed.team" && data.Password == "123456"
}
