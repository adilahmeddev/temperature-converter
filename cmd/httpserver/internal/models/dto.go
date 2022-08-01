package models

type TemperatureRequest struct {
	ConvertTo string  `json:"convert_to,omitempty"`
	Value     float64 `json:"value,omitempty"`
}

type TemperatureResponse struct {
	Type  string  `json:"type,omitempty"`
	Value float64 `json:"value,omitempty"`
}
