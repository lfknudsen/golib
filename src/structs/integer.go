package structs

import "strconv"

// Integer is a type that exists to be able to treat booleans as integers and vice versa;
// one aspect of C that I really miss in other languages.
// Bool conversions work on the concept that 0 = false, and non-zero values true.
type Integer int

func (i Integer) String() string { return strconv.Itoa(int(i)) }
func (i Integer) Bool() bool     { return i != 0 }
func (i Integer) Boolean() Bool  { return i != 0 }

// =============================================================================
// Integer Mathematics
// =============================================================================

func (i Integer) AbsI() int { return int(max(i, -i)) }

func (i Integer) Abs() Integer {
	return max(i, -i)
}

func AbsI(i int) int {
	return max(i, -i)
}

func MaxInt(a, b Integer) Integer {
	if a > b {
		return a
	}
	return b
}

func MinInt(a, b Integer) Integer {
	if a > b {
		return b
	}
	return a
}

// =============================================================================
// Integer Comparison Functions
// =============================================================================

func (i Integer) Compare(other Integer) Integer {
	return Compare(i, other)
}

func (i Integer) CompareI(other Integer) int {
	return CompareI(i, other)
}

func Compare(a, b Integer) Integer {
	return Integer(CompareI(a, b))
}

func CompareI(a, b Integer) int {
	if a < b {
		return -1
	}
	if a > b {
		return 1
	}
	return 0
}

// =============================================================================
// Known-safe Integer Type Casting
// =============================================================================

func (i Integer) Int() int        { return int(i) }
func (i Integer) Uint() uint      { return uint(i) }
func (i Integer) Int32() int32    { return int32(i) }
func (i Integer) Uint32() uint32  { return uint32(i) }
func (i Integer) Long() int64     { return int64(i) }
func (i Integer) Ulong() uint64   { return uint64(i) }
func (i Integer) Int64() int64    { return int64(i) }
func (i Integer) Uint64() uint64  { return uint64(i) }
func (i Integer) Float() float32  { return float32(i) }
func (i Integer) Double() float64 { return float64(i) }
