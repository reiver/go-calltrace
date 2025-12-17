package calltrace_test

import (
	"fmt"

	"github.com/reiver/go-calltrace"
)

func ExampleString() {

	trace := calltrace.String()

	fmt.Println(trace)

	// Output:
	// github.com/reiver/go-calltrace_test.ExampleString():go-calltrace/string_examples_test.go:11, testing.runExample():testing/run_example.go:63, testing.runExamples():testing/example.go:41, testing.(*M).Run():testing/testing.go:2339, main.main():_testmain.go:47, runtime.main():runtime/proc.go:285, runtime.goexit():runtime/asm_amd64.s:1693
}
