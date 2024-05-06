package error

import (
	"fmt"
	"net/http"
	"strings"
)

type ErrorMessage string
type errCode string

type errMsgCode struct {
	httpCode int
	code     errCode
}

const (
	// Err Message
	ErrDefault       ErrorMessage = "unexpected error occured"
	ErrDatabase      ErrorMessage = "database error"
	ErrBadRequest    ErrorMessage = "Bad Request"
	ErrUnauthorized  ErrorMessage = "Unauthorized"
	ErrForbidden     ErrorMessage = "Forbidden"
	ErrNotFound      ErrorMessage = "Not Found"
	ErrValidationMsg ErrorMessage = "validation error"

	// Err Message Auth
	ErrEmailExist        ErrorMessage = "Email already exist"
	ErrTokenAuthEmpty    ErrorMessage = "token authentication is empty"
	ErrTokenAuthInvalid  ErrorMessage = "token authentication is invalid"
	ErrTokenAuthDisabled ErrorMessage = "token authentication is disabled"
	ErrUserDeactivated   ErrorMessage = "user is deactivated"
	ErrUserNotFound      ErrorMessage = "user not found"

	// Generic Err Code
	codeDefault        errCode = "ERR-000"
	codeDatabase       errCode = "ERR-001"
	codeBadRequest     errCode = "ERR-400"
	codeUnauthorized   errCode = "ERR-401"
	codeForbidden      errCode = "ERR-403"
	codeNotFound       errCode = "ERR-404"
	codeInternalServer errCode = "ERR-500"

	// Err Code Auth
	codeEmailExist        errCode = "AUTH-001"
	codeTokenAuthEmpty    errCode = "AUTH-002"
	codeTokenAuthInvalid  errCode = "AUTH-003"
	codeTokenAuthDisabled errCode = "AUTH-004"
	codeUserDeactivated   errCode = "AUTH-005"
	codeUserNotFound      errCode = "AUTH-006"
)

var errorMapCodeMsg = map[ErrorMessage]errMsgCode{
	ErrDefault:           {httpCode: http.StatusInternalServerError, code: codeDefault},
	ErrDatabase:          {httpCode: http.StatusInternalServerError, code: codeDatabase},
	ErrBadRequest:        {httpCode: http.StatusBadRequest, code: codeBadRequest},
	ErrUnauthorized:      {httpCode: http.StatusUnauthorized, code: codeUnauthorized},
	ErrForbidden:         {httpCode: http.StatusForbidden, code: codeForbidden},
	ErrEmailExist:        {httpCode: http.StatusConflict, code: codeEmailExist},
	ErrTokenAuthEmpty:    {httpCode: http.StatusUnauthorized, code: codeTokenAuthEmpty},
	ErrTokenAuthInvalid:  {httpCode: http.StatusUnauthorized, code: codeTokenAuthInvalid},
	ErrTokenAuthDisabled: {httpCode: http.StatusUnauthorized, code: codeTokenAuthDisabled},
	ErrUserDeactivated:   {httpCode: http.StatusUnauthorized, code: codeUserDeactivated},
	ErrUserNotFound:      {httpCode: http.StatusNotFound, code: codeUserNotFound},
}

func (e ErrorMessage) Error() error {
	return e.ErrorWithMessage()
}

func (e ErrorMessage) ErrorWithMessage(customMsg ...string) error {
	msg := string(e)
	if len(customMsg) > 0 {
		// append custom message to error message if exist custom message with separator ": "
		msg = fmt.Sprintf("%s: %s", msg, strings.Join(customMsg, ": "))
	}
	if msgCode, ok := errorMapCodeMsg[e]; ok {
		return New(msgCode.httpCode, WithMessageCode(msg, string(msgCode.code)))
	}
	return New(http.StatusInternalServerError, WithMessageCode(string(ErrDefault), string(codeDefault)))
}
