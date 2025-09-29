package logging

import (
	"reflect"
	"strconv"

	"github.com/lfknudsen/golib/src/collections"
	"github.com/lfknudsen/golib/src/maths"
	"github.com/lfknudsen/golib/src/structs"
)

type ErrorType interface {
	Error() string
	String() string
}

type IndexOutOfRangeError struct {
	Attempted    structs.Int
	MinSafeIndex structs.Int
	MaxSafeIndex structs.Int
}

func (err IndexOutOfRangeError) String() string {
	return err.Error()
}

func IndexOutOfRange(attempted, min, max structs.Int) string {
	return IndexOutOfRangeError{attempted, min, max}.Error()
}

func (err IndexOutOfRangeError) Error() string {
	attempted := strconv.FormatInt(int64(err.Attempted), 10)
	minIndex := strconv.FormatInt(int64(err.MinSafeIndex), 10)
	maxIndex := strconv.FormatInt(int64(err.MaxSafeIndex), 10)
	return "Index " + attempted + " out of range [" + minIndex + "," + maxIndex + "]."
}

type UnimplementedFunctionError struct {
	Function reflect.Method
}

func (err UnimplementedFunctionError) String() string {
	return err.Error()
}

func UnimplementedFunction(function reflect.Method) string {
	return UnimplementedFunctionError{Function: function}.Error()
}

func (err UnimplementedFunctionError) Error() string {
	return "The Function " + err.Function.Name + " is not implemented yet."
}

type ExUnexpectedNilValue struct {
	Identifier string
	RefKind    reflect.Kind
	RefType    reflect.Type
	RefValue   reflect.Value
}

func (err ExUnexpectedNilValue) String() string {
	return err.Error()
}

func (err ExUnexpectedNilValue) Error() string {
	return "the value of " + err.Identifier + " is nil.\n" +
		"Kind is " + err.RefKind.String() + ".\n" +
		"Type is " + err.RefType.String() + ".\n" +
		"Value is " + err.RefValue.String() + ".\n"
}

func UnexpectedNilValueError(
	identifier string, refKind reflect.Kind, refType reflect.Type, refValue reflect.Value) string {
	return ExUnexpectedNilValue{
		identifier, refKind, refType, refValue}.Error()
}

type ExConversion struct {
	From any
	To   any
}

func (err ExConversion) String() string {
	return err.Error()
}

func (err ExConversion) Error() string {
	return "cannot convert " + reflect.TypeOf(err.From).String() +
		" to " + reflect.TypeOf(err.To).String()
}

type IWrongType interface {
	error
	Acceptable() []reflect.Kind
}

type ExNonRealNumber struct {
	Input reflect.Kind
}

type ExUnexpectedType struct {
	Expected string
	Input    any
}

func (err ExUnexpectedType) Error() string {
	return "wrong type: " + reflect.TypeOf(err.Input).Kind().String() + ". Was expecting " + err.Expected
}

func (err ExNonRealNumber) Acceptable() []reflect.Kind {
	return maths.RealNumbers
}

func (err ExNonRealNumber) Error() string {
	return "non-real numerical type " + err.Input.String() +
		"\nAcceptable options are:" + collections.Concat(err.Acceptable(), "\n")
}
