package main

import (
	"bytes"
	"encoding/json"
	"excercise4"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

type Driver struct {
	request  io.Reader
	response *io.Writer
	server   *httptest.Server
}

func (d Driver) ConvertToF(celsius float64) (float64, error) {
	marshalledRequest, err := json.Marshal(TemperatureRequest{
		ConvertTo: "f",
		Value:     celsius,
	})
	if err != nil {
		return 0, err
	}

	buffer := bytes.NewBuffer(marshalledRequest)
	request, err := http.NewRequest(http.MethodPost, d.server.URL, buffer)

	client := http.DefaultClient

	r, err := client.Do(request)
	if err != nil {
		return 0, err
	}

	var response TemperatureResponse

	decoder := json.NewDecoder(r.Body)

	err = decoder.Decode(&response)
	if err != nil {
		return 0, err
	}

	return response.Value, nil
}

func (d Driver) ConvertToC(fah float64) (float64, error) {
	//TODO implement me
	panic("implement me")
}

type TemperatureRequest struct {
	ConvertTo string  `json:"convert_to,omitempty"`
	Value     float64 `json:"value,omitempty"`
}

type TemperatureResponse struct {
	Type  string  `json:"type,omitempty"`
	Value float64 `json:"value,omitempty"`
}

type TemperatureHandler struct {
	service excercise4.TempConverter
}

func (t TemperatureHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	var tRequest TemperatureRequest
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

		response := TemperatureResponse{
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

func TestName(t *testing.T) {
	handler := NewTemperatureHandler(excercise4.Converter{})
	server := httptest.NewServer(handler)

	t.Cleanup(server.Close)

	excercise4.ConverterSpecification(t, Driver{
		server: server,
	})
}
