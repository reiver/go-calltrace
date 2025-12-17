package calltrace

import (
	"fmt"
)

// Frame represents a part of the call-trace.
//
// See also: [Frames], [FramesFunc].
type Frame struct {
	FilePath              string
	LineNumber            int
	QualifiedFunctionName string
}

func (receiver Frame) String() string {
	return fmt.Sprintf("%s():%s:%d", receiver.QualifiedFunctionName, receiver.FilePath, receiver.LineNumber)
}
