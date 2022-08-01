package main

import (
	"encoding/json"
	"excercise4"
	"excercise4/cmd/httpserver/internal/models"
	"fmt"
	"net/http"
)

type TemperatureHandler struct {
	service excercise4.TempConverter
}

func (t TemperatureHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	var tRequest models.TemperatureRequest
	decoder := json.NewDecoder(request.Body)

	err := decoder.Decode(&tRequest)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
	}

	if tRequest.ConvertTo == "f" {
		f, err := t.service.ConvertToF(tRequest.Value)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
		}

		response := models.TemperatureResponse{
			Type:  "f",
			Value: f,
		}

		encoder := json.NewEncoder(writer)

		err = encoder.Encode(response)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
		}
	}
	http.Error(writer, fmt.Errorf("not implemented").Error(), http.StatusNotImplemented)
}

func NewTemperatureHandler(service excercise4.TempConverter) *TemperatureHandler {
	return &TemperatureHandler{service: service}
}
