package handler

import "github.com/eltaljohn/echo-api/model"

type Storage interface {
	Create(*model.Person) error
	Update(int, *model.Person) error
	Delete(int) error
	GetByID(int) (model.Person, error)
	GetAll() (model.Persons, error)
}
