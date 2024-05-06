package response

import "github.com/labstack/echo/v4"

type AppSuccessOption func(*AppSuccess)

type AppSuccess struct {
	Message *string      `json:"message,omitempty"`
	Data    *interface{} `json:"data,omitempty"`
}

func SuccessData(data interface{}) AppSuccessOption {
	return func(h *AppSuccess) {
		h.Data = &data
	}
}

func SuccessMessage(message string) AppSuccessOption {
	return func(h *AppSuccess) {
		h.Message = &message
	}
}

func Success(c echo.Context, status int, option ...AppSuccessOption) error {
	appSuccess := new(AppSuccess)
	// Loop through each option
	for _, opt := range option {
		// Call the option giving the instantiated
		opt(appSuccess)
	}
	return c.JSON(status, *appSuccess)
}
