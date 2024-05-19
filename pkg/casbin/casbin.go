package casbin

import (
	"fmt"
	"net/http"
	"reflect"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	fileadapter "github.com/casbin/casbin/v2/persist/file-adapter"
	"github.com/labstack/echo/v4"
)

// Adapter is a casbin adapter
type Adapter struct {
	enforcer *casbin.Enforcer
}

// NewAdapter creates a new casbin adapter
func NewAdapter(modelPath, policyPath string) (*Adapter, error) {
	m, err := model.NewModelFromFile(modelPath)
	if err != nil {
		return nil, err
	}

	a := fileadapter.NewAdapter(policyPath)

	enforcer, err := casbin.NewEnforcer(m, a)
	if err != nil {
		return nil, err
	}

	adapter := &Adapter{
		enforcer: enforcer,
	}

	return adapter, nil
}

// Enforce enforces access control using casbin
func (a *Adapter) Enforce(subject, object, action string) (bool, error) {

	if subject != "" && !(reflect.TypeOf(subject).Kind() == reflect.String) {
		return false, fmt.Errorf("subject must be string, got %T", subject)
	}
	if object != "" && !(reflect.TypeOf(object).Kind() == reflect.String) {
		return false, fmt.Errorf("object must be string, got %T", object)
	}
	if action != "" && !(reflect.TypeOf(action).Kind() == reflect.String) {
		return false, fmt.Errorf("action must be string, got %T", action)
	}

	fmt.Printf(subject, object, action)
	allowed, err := a.enforcer.Enforce(subject, object, action)
	if err != nil {
		return false, err
	}
	return allowed, nil
}

// CasbinAuthorization is a middleware function that enforces access control using Casbin
func CasbinAuthorization(adapter *Adapter) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			subject := "user"
			object := c.Request().URL.Path
			action := c.Request().Method

			allowed, err := adapter.Enforce(subject, object, action)
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			}

			if !allowed {
				return echo.NewHTTPError(http.StatusForbidden, "User is not authorized")
			}

			// Call the next handler
			return next(c)
		}
	}
}
