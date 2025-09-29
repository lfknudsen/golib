package maths

import (
	"reflect"

	"github.com/lfknudsen/golib/src/logging"
)

var IsRealNumber = map[reflect.Kind]bool{
	reflect.Bool:    true,
	reflect.Int:     true,
	reflect.Int8:    true,
	reflect.Int16:   true,
	reflect.Int32:   true,
	reflect.Int64:   true,
	reflect.Uint:    true,
	reflect.Uint8:   true,
	reflect.Uint16:  true,
	reflect.Uint32:  true,
	reflect.Uint64:  true,
	reflect.Float32: true,
	reflect.Float64: true,
	_:               false,
}

var RealNumbers = []reflect.Kind{
	reflect.Bool,
	reflect.Int,
	reflect.Int8,
	reflect.Int16,
	reflect.Int32,
	reflect.Int64,
	reflect.Uint,
	reflect.Uint8,
	reflect.Uint16,
	reflect.Uint32,
	reflect.Uint64,
	reflect.Float32,
	reflect.Float64,
}

func IsReal(n any) bool {
	return IsRealNumber[reflect.TypeOf(n).Kind()]
}

func IsRealType(n reflect.Type) bool {
	return IsRealNumber[n.Kind()]
}

func IsRealKind(n reflect.Kind) bool {
	return IsRealNumber[n]
}

func IsComplex(n any) bool {
	kind := reflect.TypeOf(n).Kind()
	return kind == reflect.Complex64 || kind == reflect.Complex128
}

var Integers = []reflect.Kind{
	reflect.Bool,
	reflect.Int,
	reflect.Int8,
	reflect.Int16,
	reflect.Int32,
	reflect.Int64,
	reflect.Uint,
	reflect.Uint8,
	reflect.Uint16,
	reflect.Uint32,
	reflect.Uint64,
}

func IsUnsigned(val any) bool {
	switch val.(type) {
	case uint, uint8, uint16, uint32, uint64:
		return true
	default:
		return false
	}
}

func IsSigned(val any) (bool, error) {
	switch val.(type) {
	case int, int8, int16, int32, int64:
		return true, nil
	case uint, uint8, uint16, uint32, uint64:
		return false, nil
	default:
		return false,
			logging.ExUnexpectedType{Expected: "an integer", Input: val}
	}
}
