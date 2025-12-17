package calltrace

// Frames returns a call-trace as a []Frame.
//
// Example usage:
//
//	var frames []calltrace.Frame = calltrace.Frames()
//
// See also: [Frame], [FrameFunc], [String], [Strings].
func Frames() []Frame {
	var frames []Frame

	collectFrames := func(frame Frame) {
		frames = append(frames, frame)
	}

	FramesFunc(collectFrames)

	return frames
}
