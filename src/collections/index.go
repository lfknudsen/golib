package collections

import (
	"fmt"
)

// IndexList is an array-based singly-linked-list, where each element's value is the
// index of the next value. -1 represents the end of the list.
//
// Implements [IForwardIterator]
type IndexList struct {
	Indices    []int
	Cursor     int
	CursorNext int
	StartPos   int
}

func (l *IndexList) Init(list []int) IndexList {
	return IndexList{
		Indices:  list,
		Cursor:   -1,
		StartPos: 0,
	}
}

func (l *IndexList) Reset() {
	l.Cursor = l.StartPos - 1
}

func (l *IndexList) Peek() int {
	return l.Indices[l.Cursor+1]
}

func (l *IndexList) Length() int {
	backupCursor := l.Cursor
	l.Reset()
	sum := 0
	for l.HasNext() {
		sum++
		l.Next()
	}
	l.Reset()
	l.Cursor = backupCursor
	return sum
}

func (l *IndexList) HasNext() bool {
	return l.Cursor+1 < len(l.Indices) && l.Indices[l.Cursor+1] != -1
}

func (l *IndexList) Next() int {
	l.Cursor = l.Indices[l.Cursor+1]
	return l.Cursor
}

func (l *IndexList) String() string {
	return fmt.Sprintf("%v", l.Indices)
}

func (l *IndexList) IndexOf(indexed []any, needle any) int {
	if len(l.Indices) != len(indexed) || len(l.Indices) == 0 {
		return -1
	}
	if indexed[l.Cursor] == needle {
		return l.Cursor
	}
	for l.HasNext() {
		l.Next()
		if indexed[l.Cursor] == needle {
			return l.Cursor
		}
	}
	return -1
}

// IndexDList is an array-based doubly-linked-list, where each element's
// 32 most significant bits represent the index of the previous element, and
// their 32 least significant bits represent the index of the next element.
// A value of -1 represents the end of the list.
//
// Implements [IForwardIterator], [IBackwardIterator], and [IDoubleIterator]
type IndexDList struct {
	Indices []Split64
	Cursor  int32
}

func (l *IndexDList) Reset() {
	l.Cursor = 0
}

func (l *IndexDList) Peek() int32 {
	return l.Indices[l.Cursor].Right()
}

func (l *IndexDList) Length() int32 {
	backupCursor := l.Cursor
	l.Reset()
	var sum int32 = 0
	for l.HasNext() {
		sum++
		l.Next()
	}
	l.Reset()
	l.Cursor = backupCursor
	return sum
}

func (l *IndexDList) HasNext() bool {
	return (l.Cursor < int32(len(l.Indices))) &&
		(l.Indices[l.Cursor].Right() != -1)
}

func (l *IndexDList) Next() int32 {
	l.Cursor = l.Indices[l.Cursor].Right()
	return l.Cursor
}

func (l *IndexDList) HasPrev() bool {
	return (l.Cursor > 0) &&
		(l.Indices[l.Cursor].Left() != -1)
}

func (l *IndexDList) Prev() int32 {
	l.Cursor = l.Indices[l.Cursor].Left()
	return l.Cursor
}
