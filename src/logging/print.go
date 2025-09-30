package logging

import (
	"fmt"
	"strings"
	"sync"
)

// logBuilder is an internal, thread-safe string builder re-used for each string
// operation.
var logBuilder strings.Builder = strings.Builder{}
var logBuilderLock sync.Mutex = sync.Mutex{}

type StringPrinter interface {
	String() string
}

// Out prints all the parameters, separated by a single space, to the standard output.
func Out(a ...interface{}) {
	OutDelim(" ", a)
}

// Concat joins together all the parameters into a single string.
func Concat(a ...interface{}) string {
	logBuilderLock.Lock()
	defer logBuilderLock.Unlock()
	logBuilder.Reset()
	for _, c := range a {
		logBuilder.WriteString(c.(string))
	}
	return logBuilder.String()
}

func OutDelim(delim string, a ...interface{}) {
	if len(a) == 0 {
		return
	}
	if delim == "" {
		fmt.Print(Concat(a))
	}

	for _, v := range a[:len(a)-1] {
		fmt.Print(v)
		fmt.Print(delim)
	}
	fmt.Print(a[len(a)-1])
}

func Outf(format string, a ...interface{}) {
	for _, v := range a {
		fmt.Print(v)
	}
}
