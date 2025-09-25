package collections

import (
	"container/list"
)

type Array []interface{}

// RemoveAll returns a new Array with all the elements of this
// Array, except for those which are equal to the needle parameter.
func (a Array) RemoveAll(needle interface{}) Array {
	out := make(Array, len(a))
	for _, element := range a {
		if element != needle {
			out = append(out, element)
		}
	}
	return out
}

// ListToArray constructs an Array based on the given list.List.
// Though passed by reference, the source list.List is not modified.
func ListToArray(source *list.List) Array {
	out := make(Array, source.Len())
	current := source.Front()
	for i := 0; i < source.Len() && current != nil; i++ {
		out[i] = current.Value
		current = current.Next()
	}
	return out
}

// ArrayToList returns a reference to a new list.List with
// all the elements of the source Array in the same order as
// the original.
func ArrayToList(source Array) *list.List {
	l := new(list.List).Init()
	for _, element := range source {
		l.PushBack(element)
	}
	return l
}

func IdentityArray(length int) []int {
	indices := make([]int, length)
	for i := 0; i < length; i++ {
		indices[i] = i
	}
	return indices
}

func IdentityArrayU(length uint) []uint {
	indices := make([]uint, length)
	var i uint = 0
	for ; i < length; i++ {
		indices[i] = i
	}
	return indices
}
