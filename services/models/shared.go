package models

// JSONResponse for response api
type JSONResponse struct {
	Data       interface{} `json:"data,omitempty"`
	Message    string      `json:"message,omitempty"`
	Status     string      `json:"status,omitempty"`
	TotalData  int         `json:"total_data,omitempty"`
	TotalPage  int         `json:"total_page,omitempty"`
	Validation interface{} `json:"validation,omitempty"`
}
