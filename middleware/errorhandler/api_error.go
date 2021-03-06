package errorhandler

type APIError struct {
	ErrorType    ErrorType `json:"error_type"`
	ErrorMessage string    `json:"error_message"`
}

func (e *APIError) Error() string {
	return e.ErrorMessage
}

func NewAPIError(errorType ErrorType, errorMessage string) *APIError {
	return &APIError{
		ErrorType:    errorType,
		ErrorMessage: errorMessage,
	}
}
