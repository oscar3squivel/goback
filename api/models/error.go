package models

//Error: Represenation of http requests errors and MySQL Connection error
type ErrorHandler struct {
	OK      bool   `json:ok`
	Status  int64  `json:status`
	Message string `json:message`
}
