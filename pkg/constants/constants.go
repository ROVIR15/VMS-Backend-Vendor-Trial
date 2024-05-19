package constants

import "time"

// AppEnv is the environment variable for the app environment
// ProjectName is the name of the project
// AppRoot is the root path of the app
// ConfigPath is the path to the config file
// DefaultConfigPath is the default path to the config file
// Json is the json file type
// Name is the name of the app
// MaxHeaderBytes is the maximum number of bytes the server will read parsing the request header's keys and values, including the request line.
// BodyLimit is the maximum size of the request body
// ReadTimeout is the maximum duration for reading the entire request, including the body.
// WriteTimeout is the maximum duration before timing out writes of the response.
// GzipLevel is the gzip compression level
// WaitShutDownDuration is the duration to wait before shutting down the server
const (
	AppEnv      = "APP_ENV"
	ProjectName = "PROJECT_NAME"
	AppRoot     = "APP_ROOT"
	ConfigPath  = "CONFIG_PATH"
)

const (
	DefaultConfigPath    = "./cmd/app/config"
	Offline              = "offline"
	Json                 = "json"
	Yaml                 = "yaml"
	Name                 = "NAME"
	MaxHeaderBytes       = 1 << 20
	BodyLimit            = "2M"
	ReadTimeout          = 15 * time.Second
	WriteTimeout         = 15 * time.Second
	GzipLevel            = 5
	WaitShutDownDuration = 3 * time.Second
	Local                = "local"
	Dev                  = "development"
	Staging              = "staging"
	Production           = "production"
)

const (
	ErrBadRequestTitle          = "Bad Request"
	ErrConflictTitle            = "Conflict Error"
	ErrNotFoundTitle            = "Not Found"
	ErrUnauthorizedTitle        = "Unauthorized"
	ErrForbiddenTitle           = "Forbidden"
	ErrRequestTimeoutTitle      = "Request Timeout"
	ErrInternalServerErrorTitle = "Internal Server Error"
	ErrDomainTitle              = "Domain Model Error"
	ErrApplicationTitle         = "Application Service Error"
	ErrApiTitle                 = "Api Error"
)
