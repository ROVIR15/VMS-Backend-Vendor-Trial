package utils

import "fmt"

// Port returns a string with the port number prefixed by a colon. For example, Port(8080) returns ":8080".
func ColonPort(port int64) string {
	return fmt.Sprintf(":%d", port)
}
