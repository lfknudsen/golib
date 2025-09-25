package logging

import (
	"fmt"
	"strings"
)

type StringPrinter interface {
	String() string
}

func Out(a ...interface{}) {
	OutDelim(" ", a)
}

func Concat(a ...interface{}) string {
	b := strings.Builder{}
	for _, c := range a {
		b.WriteString(c.(string))
	}
	return b.String()
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
