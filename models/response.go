package models

type Response struct {
	Code     string      `json:"code"`
	Message  string      `json:"message"`
	Response interface{} `json:"response"`
}

type ResponsePage struct {
	Code       string      `json:"code"`
	Message    string      `json:"message"`
	Response   interface{} `json:"response"`
	Pagination interface{} `json:"pagination,omitempty"`
}
