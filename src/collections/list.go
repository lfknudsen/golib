package collections

import (
	"reflect"

	"github.com/lfknudsen/golib/src/logging"
	"github.com/lfknudsen/golib/src/structs"
)

type List[T interface{}] struct {
	length   structs.Int
	capacity structs.Int
	elemSize structs.Int64
	array    []T
}

func (l List[T]) Get(index structs.Int) (any, error) {
	if index < 0 || index >= l.length {
		return nil, logging.IndexOutOfRangeError{
			Attempted:    index,
			MinSafeIndex: 0,
			MaxSafeIndex: l.length - 1}
	}
	return l.array[index], nil
}

func (l List[T]) First() any {
	//TODO implement me
	panic("implement me")
}

func (l List[T]) Last() any {
	//TODO implement me
	panic("implement me")
}

func (l List[T]) Length() int64 {
	//TODO implement me
	panic("implement me")
}

func SumAll(array []structs.Int) (structs.Int, error) {
	if array == nil {
		return 0, logging.ExUnexpectedNilValue{
			Identifier: "array",
			RefKind:    reflect.Array,
			RefType:    reflect.TypeOf([]structs.Int{}),
			RefValue:   reflect.ValueOf([]structs.Int{}),
		}
	}

	return array[0] + SumAll(array[1:]), nil
}

func CollectRecursive[T any, C any](
	array []T, accumulator C,
	f func([]T) (C, error)) (C, error) {

	accumulator, err := f(array)
	return accumulator, err
}

func CollectIterative[T any, C any](array []T, accumulator C, f func(T, C) (C, error)) (C, error) {
	var err error
	for _, element := range array {
		accumulator, err = f(element, accumulator)
		if err != nil {
			break
		}
	}
	return accumulator, err
}

type IntList []int

func (l IntList) Collect(f func(val int, coll int64) (int, int64)) (int64, error) {
	var collector int64 = 0
	for _, i := range l {
		_, collector = f(i, collector)
	}
}
