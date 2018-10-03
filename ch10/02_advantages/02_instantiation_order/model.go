package main

import (
	"database/sql"
	"errors"
)

func NewGetPersonModel(db *sql.DB) *GetPersonModel {
	return &GetPersonModel{
		db: db,
	}
}

type GetPersonModel struct {
	db *sql.DB
}

func (g *GetPersonModel) LoadByID(ID int) (*Person, error) {
	return nil, errors.New("not implemented yet")
}

type Person struct {
	Name string
}
