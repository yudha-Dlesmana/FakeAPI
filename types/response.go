package types

type Response struct {
	// Status code (e.g. 200 or 400)
	Status int `json:"status"`
	// Descriptive message
	Message string `json:"message"`
	// Data Payload
	Data any `json:"data"`
}

type DefaultResponse struct {
	Message string `json:"message"`
	Info    string `json:"info"`
	Status  string `json:"status"`
	Version string `json:"version"`
}
