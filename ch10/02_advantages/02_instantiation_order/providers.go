package main

import (
	"database/sql"

	"github.com/google/go-cloud/wire"
)

func ProvideHandler(model *GetPersonModel) *GetPersonHandler {
	return &GetPersonHandler{
		model: model,
	}
}

func ProvideModel(db *sql.DB) *GetPersonModel {
	return &GetPersonModel{
		db: db,
	}
}

func ProvideDatabase() *sql.DB {
	return &sql.DB{}
}

var wireSet = wire.NewSet(
	ProvideHandler,
	ProvideModel,
	ProvideDatabase,
)
