package logging

import (
	"log"
	"os"
	"reflect"
	"strconv"
)

func ErrorCheck(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func PanicCheck(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func WarningCheck(err error) {
	if err != nil {
		log.Println(err)
	}
}

func ErrorCheckMsg(err error, messages ...string) {
	if err != nil {
		if messages == nil {
			log.Fatal(err)
		}
		log.Println(err)
		for _, v := range messages {
			log.Println(v)
		}
		os.Exit(1)
	}
}

type IndexOutOfRangeError struct {
	Attempted    int
	MinSafeIndex int
	MaxSafeIndex int
}

func IndexOutOfRange(attempted, min, max int) string {
	return IndexOutOfRangeError{attempted, min, max}.Error()
}

func (err IndexOutOfRangeError) Error() string {
	attempted := strconv.FormatInt(int64(err.Attempted), 10)
	minIndex := strconv.FormatInt(int64(err.MinSafeIndex), 10)
	maxIndex := strconv.FormatInt(int64(err.MaxSafeIndex), 10)
	return "Index " + attempted + " out of range [" + minIndex + "," + maxIndex + "]."
}

type UnimplementedFunctionError struct {
	function reflect.Method
}

func UnimplementedFunction(function reflect.Method) string {
	return UnimplementedFunctionError{function: function}.Error()
}

func (err UnimplementedFunctionError) Error() string {
	return "The function " + err.function.Name + " is not implemented yet."
}

type UnexpectedNilValue struct {
	identifier string
	refKind    reflect.Kind
	refType    reflect.Type
	refValue   reflect.Value
}

func (err UnexpectedNilValue) Error() string {
	return "the value of " + err.identifier + " is nil.\n" +
		"Kind is " + err.refKind.String() + ".\n" +
		"Type is " + err.refType.String() + ".\n" +
		"Value is " + err.refValue.String() + ".\n"
}

func UnexpectedNilValueError(
	identifier string, refKind reflect.Kind, refType reflect.Type, refValue reflect.Value) string {
	return UnexpectedNilValue{
		identifier, refKind, refType, refValue}.Error()
}

type ConversionError struct {
	From any
	To   any
}

func (err ConversionError) Error() string {
	return "cannot convert " + reflect.TypeOf(err.From).String() +
		" to " + reflect.TypeOf(err.To).String()
}
