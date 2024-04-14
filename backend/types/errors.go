package types

type RequestError struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}

func (m *RequestError) Error() string {
	return m.Message
}

type ErrorResponse struct {
	StatusCode int    `json:"statusCode"`
	Error      string `json:"error"`
}
