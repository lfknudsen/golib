package text

import (
	"container/list"
	"fmt"
	"sort"
	"strings"

	"github.com/lfknudsen/golib/src/collections"
)

type String string

func (s String) Append(prefix String) String {
	return Join(s, prefix)
}

func (s String) String() string {
	return string(s)
}

func Join(str ...String) String {
	length := 0
	for _, part := range str {
		length += len(part)
	}
	builder := Builder{}
	builder.GrowTo(length)
	for _, part := range str {
		builder.WriteString(part)
	}
	return builder.String()
}

type Builder struct {
	length int
	cap    int // Use this since cap(v) returns a different result for arrays and slices.
	runes  []rune
}

func NewBuilder(capacity int) *Builder {
	return &Builder{
		length: 0,
		runes:  make([]rune, capacity),
	}
}

func (b *Builder) WriteString(str String) int {
	b.Append(str)
	return len(str)
}

func (b *Builder) Length() int {
	return b.length
}

func (b *Builder) Cap() int {
	return b.cap
}

func (b *Builder) Runes() []rune {
	return b.runes
}

func (b *Builder) String() String {
	return String(b.runes)
}

func (b *Builder) Append(strings ...String) {
	lengthNecessary := b.Length()
	for _, str := range strings {
		lengthNecessary += len(str)
	}
	b.GrowTo(lengthNecessary)
}

// ChangeCapTo changes the capacity of the underlying array to the amount
// specified. If the new capacity is larger than the old one, then it will
// grow, and if it is smaller, then it will shrink.
// See also ChangeCapBy, ShrinkBy, ShrinkTo, GrowBy, and GrowTo.
func (b *Builder) ChangeCapTo(newCapacity int) *Builder {
	if newCapacity < b.cap {
		return b.ShrinkTo(newCapacity)
	}
	return b.GrowTo(newCapacity)
}

// ChangeCapBy changes the capacity of the underlying array by the amount
// specified. A negative delta causes it to shrink, and a positive one to grow.
//
// See also ChangeCapTo, ShrinkBy, ShrinkTo, GrowBy, and GrowTo.
func (b *Builder) ChangeCapBy(delta int) *Builder {
	if delta < 0 {
		return b.ShrinkBy(-delta)
	} else if delta > 0 {
		return b.GrowBy(delta)
	}
	return b
}

// ShrinkTo reduces the underlying string capacity to the specified amount.
// Returns this Builder for chaining.
// If the new capacity is greater than the current capacity, then nothing
// is changed; the underlying string will not grow.
// Negative inputs are handled as if they were 0.
//
// See also ShrinkBy, GrowBy, GrowTo, ChangeCapBy, and ChangeCapTo.
func (b *Builder) ShrinkTo(newCapacity int) *Builder {
	if newCapacity <= 0 {
		b.cap = 0
		b.length = 0
		b.runes = make([]rune, 0)
	} else if b.cap >= newCapacity {
		if b.length >= newCapacity {
			b.length = newCapacity
			b.runes = b.runes[:newCapacity]
		} else if b.length < newCapacity {
			newRuneArray := make([]rune, newCapacity)
			copy(newRuneArray, b.runes)
			b.runes = newRuneArray
		}
		b.cap = newCapacity
	}
	return b
}

// ShrinkBy reduces the underlying string capacity by the specified amount.
// Returns this Builder for chaining.
// A negative number will cause nothing to be changed; the underlying string
// will not grow.
//
// See also ShrinkTo, GrowBy, GrowTo, ChangeCapBy, and ChangeCapTo.
func (b *Builder) ShrinkBy(capacityChange int) *Builder {
	if capacityChange <= 0 {
		return b
	}
	newCapacity := b.cap - capacityChange
	return b.ShrinkTo(newCapacity)
}

// GrowTo increases the underlying string capacity to the specified amount.
// Returns this Builder for chaining.
// If the input length is smaller than the current capacity, then nothing
// is changed; the underlying string will not shrink.
//
// See also GrowBy, ShrinkBy, ShrinkTo, ChangeCapBy, and ChangeCapTo.
func (b *Builder) GrowTo(newCapacity int) *Builder {
	if b.cap >= newCapacity {
		return b
	}
	runes := make([]rune, newCapacity)
	b.runes = append(runes, b.runes...)
	b.cap = b.cap + newCapacity
	return b
}

// GrowBy increases the underlying string capacity by the specified amount.
// Returns this Builder for chaining.
// A negative number will cause nothing to be changed; the underlying string
// will not shrink.
//
// See also GrowTo, ShrinkBy, ShrinkTo, ChangeCapBy, and ChangeCapTo.
func (b *Builder) GrowBy(capacityDelta int) *Builder {
	if capacityDelta <= 0 {
		return b
	}
	return b.GrowTo(b.cap + capacityDelta)
}

func Concat(array []any) string {
	builder := strings.Builder{}
	var capNeeded int
	for _, value := range array {
		switch value := value.(type) {
		case fmt.Stringer:
			capNeeded += len(value.String())
		case error:
			capNeeded += len(value.Error())
		default:
			capNeeded += len(value.(string))
		}
	}
	builder.Grow(capNeeded)
	for _, value := range array {
		builder.WriteString(fmt.Sprintf("%v", value))
	}
	return builder.String()
}

// ConcatBetween assembles a string out of the supplied array, separating each
// element with the delim string. An empty array returns an empty string.
func ConcatBetween(array []any, delim string) string {
	if len(array) == 0 {
		return ""
	}
	if delim == "" {
		return Concat(array)
	}
	if len(array) == 1 {
		return array[0].(string)
	}

	builder := strings.Builder{}
	var capNeeded = len(array) - 1
	for _, value := range array {
		switch value := value.(type) {
		case fmt.Stringer:
			capNeeded += len(value.String())
		case error:
			capNeeded += len(value.Error())
		default:
			capNeeded += len(value.(string))
		}
	}
	builder.Grow(capNeeded)
	for _, value := range array[:len(array)-1] {
		builder.WriteString(fmt.Sprintf("%v%s", value, delim))
	}
	builder.WriteString(fmt.Sprintf("%v", array[len(array)-1]))
	return builder.String()
}

func RemoveAllSubstrings(src string, needles ...string) string {
	indices := collections.IdentityArray(len(src) + 1)
	indices[len(src)] = -1
	il := collections.IndexList{Indices: indices}

	sort.Strings(needles)

	//needleArray := make([][]rune, len(needles))
	for _, needle := range needles {
		var idx = strings.Index(src, needle)
		for idx != -1 {
			il.Indices[idx] = idx + len(needle)
		}
		//needleArray[i] = []rune(needle)
	}
	//return String(RemoveAllRunes(&runes, &needleArray))
}

func RemoveAllRunes(src *[]rune, needles *[][]rune) []rune {
	for i := 0; i < len(*needles); i++ {

	}
}
