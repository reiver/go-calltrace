package calltrace

import (
	"strings"
)

// String returns a call-trace as a string.
//
// String is the simplest way to get the call-trace.
//
// Example usage:
//
//	var trace string = calltrace.String()
//
// See also: [Frames], [FramesFunc], [Strings].
func String() string {
	trace := Strings()

	return strings.Join(trace, ", ")
}
