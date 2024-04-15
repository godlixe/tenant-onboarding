package httpx

import "fmt"

type HttpError struct {
	Message    string
	ErrorData  error
	StatusCode int
}

func NewError(
	Message string,
	Error error,
	StatusCode int,
) HttpError {
	return HttpError{
		Message:    Message,
		ErrorData:  Error,
		StatusCode: StatusCode,
	}
}

func (e HttpError) Error() string {
	return fmt.Sprintf("message: %s, error: %s", e.Message, e.ErrorData.Error())
}
