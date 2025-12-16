package calltrace

import (
	"strings"
)

// String returns a call-trace as a string.
//
// Example usage:
//
//	var trace string = calltrace.String()
//
// See also: [Strings]
func String() string {
	trace := Strings()

	return strings.Join(trace, ", ")
}
