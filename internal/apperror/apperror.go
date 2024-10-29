package apperror

import "net/http"

var (
	ErrInternalServerError     = newError(http.StatusInternalServerError, "Internal Server Error")
	ErrNotFound                = newError(http.StatusNotFound, "Not Found")
	ErrBadRequest              = newError(http.StatusBadRequest, "Bad Request")
	ErrUserNotFound            = newError(http.StatusNotFound, "Username not found")
	ErrUserHasRegister         = newError(http.StatusConflict, "Username has been registered")
	ErrInvalidUsernamePassword = newError(http.StatusUnauthorized, "Invalid username or password")
	ErrInvalidToken            = newError(http.StatusUnauthorized, "Invalid token")
	ErrUnauthorized            = newError(http.StatusUnauthorized, "Unauthorized")
	ErrInvalidPasswordFormat   = newError(http.StatusUnauthorized, "Password must include at least one lowercase letter, one uppercase letter, one digit, and one special character")
)

type Error struct {
	Code    int
	Message string
}

func (e *Error) Error() string {
	return e.Message
}

func newError(code int, message string) *Error {
	return &Error{
		Code:    code,
		Message: message,
	}
}
