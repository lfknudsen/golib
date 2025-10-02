package structs

import (
	"strconv"
)

type IndexOutOfRangeError struct {
	Attempted    Int
	MinSafeIndex Int
	MaxSafeIndex Int
}

func (err IndexOutOfRangeError) String() string {
	return err.Error()
}

func IndexOutOfRange(attempted, min, max Int) string {
	return IndexOutOfRangeError{attempted, min, max}.Error()
}

func (err IndexOutOfRangeError) Error() string {
	attempted := strconv.FormatInt(int64(err.Attempted), 10)
	minIndex := strconv.FormatInt(int64(err.MinSafeIndex), 10)
	maxIndex := strconv.FormatInt(int64(err.MaxSafeIndex), 10)
	return "Index " + attempted + " out of range [" + minIndex + "," + maxIndex + "]."
}
