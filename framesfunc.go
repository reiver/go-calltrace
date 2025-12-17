package calltrace

import (
	"runtime"
	"strings"
)

// FramesFunc iterates through the call-trace.
//
// Example usage:
//
//	var trace []Frame
//	calltrace.FramesFunc(func(frame Frame){
//		trace = append(trace, frame)
//	})
//
// See also: [Frame], [Frames], [FramesFunc], [String], [Strings].
func FramesFunc(fn func(frame Frame)) {
	if nil == fn {
		return
	}

	var frames *runtime.Frames
	{
		const max = 13
		var buffer [max]uintptr

		var pc []uintptr = buffer[:]

		n := runtime.Callers(0, pc)
		if n <= 0 {
			return
		}
		pc = pc[:n]
		frames = runtime.CallersFrames(pc)
	}

	{
		for {
			frame, more := frames.Next()

			if "runtime.Callers" == frame.Function {
				continue
			}
			if strings.HasPrefix(frame.Function, "github.com/reiver/go-calltrace.") {
				continue
			}

			var frm Frame = Frame{
				FilePath:              frame.File,
				QualifiedFunctionName: frame.Function,
				LineNumber:            frame.Line,
			}

			fn(frm)

			if !more {
				break
			}
		}
	}
}
