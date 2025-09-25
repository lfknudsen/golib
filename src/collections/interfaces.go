package collections

import "github.com/lfknudsen/golib/src/structs"

type IImmutableList[T any] interface {
	Get(index structs.Int) (T, error)
	First() T
	Last() T
	Length() structs.Int64
}

type IStack[T any] interface {
	PopBack() T
	PushBack(v T) bool
	PeekBack() T
}

type IQueue[T any] interface {
	PopFront() T
	PushBack(v T) bool
	PeekFront() T
}

type IDeque[T any] interface {
	IQueue[T]
	PopBack() T
	PushFront(v T) bool
	PeekBack() T
}

type IGenericList[T any] interface {
	IImmutableList[T]
	Insert(index structs.Int, value T)
}

type ISequence[T any] interface {
	IGenericList[T]
	IDoubleIterator[T]
}

type IForwardSequence[T any] interface {
	IGenericList[T]
	IForwardIterator[T]
}

type IBackwardSequence[T any] interface {
	IGenericList[T]
	IForwardIterator[T]
}

type IDoubleIterator[T any] interface {
	IForwardIterator[T]
	IBackwardIterator[T]
}

type IBackwardIterator[T any] interface {
	Prev() T
	HasPrev() bool
}

type IForwardIterator[T any] interface {
	Next() T
	HasNext() bool
}

type IMapper[T any] interface {
	Map(func(T) (T, error)) error
}

type IAccumulator[T any, I any] interface {
	Acc(func(val T, acc I) (T, I)) (I, error)
}

type ICollector[T any, C any] interface {
	Collect(func(val T, coll C) (T, C)) (C, error)
}

type IReducer[T any, R any] interface {
	Reduce(func(element T, reduction R) (R, error)) (R, error)
}
