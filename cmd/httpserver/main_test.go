package main

import (
	"bytes"
	"encoding/json"
	"excercise4"
	"excercise4/cmd/httpserver/internal/models"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestName(t *testing.T) {
	handler := NewTemperatureHandler(excercise4.Converter{})
	server := httptest.NewServer(handler)

	t.Cleanup(server.Close)

	excercise4.ConverterSpecification(t, Driver{
		server: server,
	})
}

type Driver struct {
	request  io.Reader
	response *io.Writer
	server   *httptest.Server
}

func (d Driver) ConvertToF(celsius float64) (float64, error) {
	marshalledRequest, err := json.Marshal(models.TemperatureRequest{
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

	var response models.TemperatureResponse

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
