package collections

import (
	"fmt"
	"math"
)

// Split64 allows for conveniently storing and retrieving two uint32 values in a single variable.
type Split64 uint64

func (s Split64) String() string {
	left := s.Left()
	right := s.Right()
	return fmt.Sprintf("< %d | %d >", left, right)
}

// Left retrieves index stored in the 32 most significant bits, and returns them as an int.
func (s Split64) Left() int32 {
	return int32(s >> 32)
}

// Right retrieves index stored in the 32 least significant bits, and returns them as an int.
func (s Split64) Right() int32 {
	return int32(s & math.MaxUint32)
}

func (s Split64) SetLeft(val int32) Split64 {
	fmt.Printf("Input: %x\n", val)
	fmt.Printf("Split: %x\n", uint64(s))
	left := uint64(val) << 32
	right := uint64(s.Right()) << 32 >> 32
	concat := left | right
	fmt.Printf("Left: %x; Right: %x\n", left, right)
	fmt.Printf("Result: %x\n", concat)
	return Split64(concat)
}

func (s Split64) SetRight(val int32) Split64 {
	fmt.Printf("Input Right: %x\n", val)
	fmt.Printf("Input Right (uint): %x\n", uint32(val))
	right := uint64(val) << 32 >> 32
	fmt.Printf("Split: %x\n", uint64(s))
	left := uint64(s.Left()) << 32
	fmt.Printf("Left: %x; ", left)
	fmt.Printf("Right: %x\n ", right)
	result := left | right
	fmt.Printf("Result: %x\n", result)
	return Split64(result)
}
