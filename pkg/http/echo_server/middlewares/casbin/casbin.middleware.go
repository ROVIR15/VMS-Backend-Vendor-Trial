package middleware

import (
	"fmt"
	"net/http"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	fileadapter "github.com/casbin/casbin/v2/persist/file-adapter"
	"github.com/labstack/echo/v4"
)

type Adapter struct {
	enforcer *casbin.Enforcer
}

func NewCasbinAdapter(modelPath, policyPath string) (*Adapter, error) {

	m, err := model.NewModelFromFile(modelPath)
	if err != nil {
		return nil, err
	}

	a := fileadapter.NewAdapter(policyPath)

	enforcer, err := casbin.NewEnforcer(m, a)
	if err != nil {
		return nil, err
	}

	return &Adapter{
		enforcer: enforcer,
	}, nil
}

func (a *Adapter) Enforcer(sub, obj, act string) (bool, error) {
	return a.enforcer.Enforce(sub, obj, act)
}

// CasbinMiddleware is a middleware function that enforces access control using Casbin
type errorResponse struct {
	Error       string `json:"error"`
	ErrorDetail string `json:"error_detail,omitempty"`
}

func CasbinMiddleware(ce *Adapter) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			subject := "user"
			object := c.Request().URL.Path
			action := c.Request().Method

			allowed, err := ce.Enforcer("user", "/api/vendor", "GET")
			if err != nil {
				return c.JSON(http.StatusInternalServerError, errorResponse{
					Error:       err.Error(),
					ErrorDetail: fmt.Sprintf("error while enforcing access control: %v", err),
				})
			}

			if !allowed {
				return c.JSON(http.StatusForbidden, errorResponse{
					Error:       http.StatusText(http.StatusForbidden),
					ErrorDetail: fmt.Sprintf("access control denied: %s %s %s", subject, object, action),
				})
			}

			return next(c)
		}
	}
}
