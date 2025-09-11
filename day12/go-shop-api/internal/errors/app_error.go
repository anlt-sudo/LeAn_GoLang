package errors

import "time"

type AppError struct {
	Code      int               `json:"status"`
	Message   string            `json:"message"`
	Errors    map[string]string `json:"errors,omitempty"`
	Timestamp time.Time
}

func (e *AppError) Error() string {
	return e.Message
}

func New(code int, msg string) *AppError {
	return &AppError{Code: code, Message: msg}
}
func NewWithErrors(code int, msg string, errs map[string]string) *AppError {
	return &AppError{Code: code, Message: msg, Errors: errs, Timestamp: time.Now()}
}


