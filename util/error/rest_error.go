package exception

type Exception struct {
	Message string `json:message`
	Status  uint   `json:status`
	Error   string `json:error`
}

func NewException(message string, status uint, error string) Exception {
	return Exception{
		Message: message,
		Status: status,
		Error: error,
	}
}
