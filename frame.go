package calltrace

import (
	"fmt"
	"strings"
)

// Frame represents a part of the call-trace.
//
// See also: [Frames], [FramesFunc].
type Frame struct {
	FilePath              string
	LineNumber            int
	QualifiedFunctionName string
}

func (receiver Frame) ShortFilePath() string {
	var filename string = receiver.FilePath

	if index1 := strings.LastIndex(filename, "/"); 0 <= index1 {
		if index2 := strings.LastIndex(filename[:index1], "/"); 0 <= index1 {
			filename = filename[1+index2:]
		}
	}

	return filename
}

func (receiver Frame) String() string {
	return fmt.Sprintf("%s():%s:%d", receiver.QualifiedFunctionName, receiver.FilePath, receiver.LineNumber)
}
