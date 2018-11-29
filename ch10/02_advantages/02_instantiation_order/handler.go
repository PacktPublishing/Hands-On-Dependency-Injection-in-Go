package main

import (
	"net/http"
)

func NewGetPersonHandler(model *GetPersonModel) *GetPersonHandler {
	return &GetPersonHandler{
		model: model,
	}
}

type GetPersonHandler struct {
	model *GetPersonModel
}

func (g *GetPersonHandler) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	response.WriteHeader(http.StatusInternalServerError)
	response.Write([]byte(`not implemented yet`))
}
