package calltrace

import (
	"fmt"
	"strings"
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
		var filename string = frame.FilePath

		if index := strings.LastIndex(filename, "/"); 0 <= index {
			filename = filename[1+index:]
		}

		s := fmt.Sprintf("%s():%s:%d", frame.QualifiedFunctionName, frame.FilePath, frame.LineNumber)

		trace = append(trace, s)
	}

	FramesFunc(collectFrames)

	return trace
}
