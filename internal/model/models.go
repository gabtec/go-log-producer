package model

type ApiError struct {
	StatusCode int    `json:"statusCode"`
	StatusText string `json:"statusText"`
	Message    string `json:"message"`
}
