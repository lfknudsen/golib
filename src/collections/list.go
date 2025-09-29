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

func (l *List[T]) Get(index structs.Int) (any, error) {
	if index < 0 || index >= l.length {
		return nil, logging.IndexOutOfRangeError{
			Attempted:    index,
			MinSafeIndex: 0,
			MaxSafeIndex: l.length - 1,
		}
	}
	return l.array[index], nil
}

func (l *List[T]) First() any {
	//TODO implement me
	panic("implement me")
}

func (l *List[T]) Last() any {
	//TODO implement me
	panic("implement me")
}

func (l *List[T]) Length() int64 {
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
	if len(array) == 0 {
		return 0, nil
	}
	sum, err := SumAll(array[1:])
	return array[0] + sum, err
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

type IntList List[structs.Int]

func (l *IntList) Aggregate(
	f func(element structs.Int, aggregate structs.Int64) (aggregated structs.Int64, err error)) (
	finalAggregate structs.Int64, finalErr error) {
	var collector structs.Int64 = 0
	var err error
	for _, i := range l.array {
		collector, err = f(i, collector)
		if err != nil {
			break
		}
	}
	return collector, err
}

func (l *IntList) Sum() structs.Int64 {
	sum := structs.Int64(0)
	for i := 0; i < len(l.array); i++ {
		sum += structs.Int64(l.array[i])
	}
	return sum
}
