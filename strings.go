package calltrace

import (
	"fmt"
)


// Strings returns a call-trace as a []string.
//
// Example usage:
//
//	var trace []string = calltrace.Strings()
//
// See also: [String]
func Strings() []string {

	var trace []string

	collectFrames := func(frame Frame) {
		var filename string = frame.ShortFilePath()

		s := fmt.Sprintf("%s():%s:%d", frame.QualifiedFunctionName, filename, frame.LineNumber)

		trace = append(trace, s)
	}

	FramesFunc(collectFrames)

	return trace
}
