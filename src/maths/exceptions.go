package maths

import (
	"reflect"

	"github.com/lfknudsen/golib/src/text"
)

type ExNonRealNumber struct {
	Input reflect.Kind
}

func (err ExNonRealNumber) Acceptable() []reflect.Kind {
	return RealNumbers
}

func (err ExNonRealNumber) Error() string {
	acceptable, _ := text.ConcatBetween(RealNumbers, "\n")
	return "non-real numerical type " + err.Input.String() +
		"\nAcceptable options are:" + acceptable
}
