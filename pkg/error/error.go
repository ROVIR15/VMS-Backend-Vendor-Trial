package error

import (
	// validation "github.com/go-ozzo/ozzo-validation/v4"

	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type AppErrorOption func(*AppError)

// APPError is the default error struct containing detailed information about the error
type AppError struct {
	// HTTP Status code to be set in response
	Status int `json:"-"`
	// Code From Frontend
	Code string `json:"code"`
	// Message is the error message that may be displayed to end users
	Message *string `json:"message,omitempty"`
}

// New generates an application error
func New(status int, opts ...AppErrorOption) *AppError {
	err := new(AppError)
	// Loop through each option
	for _, opt := range opts {
		// Call the option giving the instantiated
		opt(err)
	}
	err.Status = status
	return err
}

// Error returns the error message.
func (e AppError) Error() string {
	return *e.Message
}

func WithMessage(message string) AppErrorOption {
	return func(h *AppError) {
		h.Message = &message
	}
}

func WithMessageCode(message string, code string) AppErrorOption {
	return func(h *AppError) {
		h.Message = &message
		h.Code = code
	}
}

// Response writes an error response to client
func Response(c echo.Context, err error) error {
	switch e := err.(type) {
	case *AppError:
		return c.JSON(e.Status, err)
	case validator.ValidationErrors:
		return c.JSON(echo.ErrBadRequest.Code, newErrValidation(e))
	default:
		return c.JSON(http.StatusInternalServerError, err)
	}
}
