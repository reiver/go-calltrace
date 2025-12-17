# go-calltrace

Package **calltrace** provides tools to get a **call-trace**, for the Go programming language.

## Online Documention

Online documentation, which includes examples, can be found at: http://godoc.org/codeberg.org/reiver/go-calltrace

[![GoDoc](https://godoc.org/codeberg.org/reiver/go-calltrace?status.svg)](https://godoc.org/codeberg.org/reiver/go-calltrace)

## Example

Here is the simplest way to get a call-trace:

```golang
var trace string = calltrace.String()
```

This package has 2 other more advanced ways of gettnig the call-trace.

## Import

To import package **calltrace** use `import` code like the following:

```
import "codeberg.org/reiver/go-calltrace"
```

## Installation

To install package **calltrace** do the following:

```
GOPROXY=direct go get codeberg.org/reiver/go-calltrace
```

## Author

Package **calltrace** was written by [Charles Iliya Krempeaux](http://reiver.link)
