package exception

type UnauthorizedError struct {
	Message string
}

func NewUnauthorizedError(err string) UnauthorizedError {
	return UnauthorizedError{Message: err}
}
