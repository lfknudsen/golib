package structs

import (
	"errors"
	"math"
	"math/big"
	"reflect"
	"strconv"
)

// Int is a type that exists to be able to treat booleans as integers and vice versa;
// one aspect of C that I really miss in other languages.
// Bool conversions work on the concept that 0 = false, and non-zero values true.
type Int int

type IInt interface {
	Int() Int
}

func (i Int) String() string { return strconv.Itoa(int(i)) }
func (i Int) Bool() Bool     { return i != 0 }

// =============================================================================
// Int Mathematics
// =============================================================================

func (i Int) AbsI() Int { return max(i, -i) }

func (i Int) Abs() Int {
	return max(i, -i)
}

func AbsI(i Int) Int {
	return max(i, -i)
}

func MaxInt(a, b Int) Int {
	if a > b {
		return a
	}
	return b
}

func MinInt(a, b Int) Int {
	if a > b {
		return b
	}
	return a
}

// =============================================================================
// Bool Mathematics
// =============================================================================

func (i Int) AddB(o Bool) Int {
	return i + o.Int()
}

func (i Int) SubB(o Bool) Int {
	return i - o.Int()
}

func (i Int) MulB(o Bool) Int {
	return i * o.Int()
}

// =============================================================================
// Int Comparison Functions
// =============================================================================

func (i Int) Compare(other Int) Int {
	return CompareInts(i, other)
}

func CompareInts(a, b Int) Int {
	if a < b {
		return -1
	}
	if a > b {
		return 1
	}
	return 0
}

func (i Int) CompareB(other Bool) Int {
	return CompareInts(i, other.Int())
}

// =============================================================================
// Known-safe Int Type Casting
// =============================================================================

func (i Int) Int() int        { return int(i) }
func (i Int) Uint() uint      { return uint(i) }
func (i Int) Int32() int32    { return int32(i) }
func (i Int) Uint32() uint32  { return uint32(i) }
func (i Int) Long() int64     { return int64(i) }
func (i Int) Ulong() uint64   { return uint64(i) }
func (i Int) Int64() int64    { return int64(i) }
func (i Int) Uint64() uint64  { return uint64(i) }
func (i Int) Float() float32  { return float32(i) }
func (i Int) Double() float64 { return float64(i) }

func SumOverflowInt(a, b int) bool {
	bigA := big.NewInt(int64(a))
	bigB := big.NewInt(int64(b))
	bigMax := big.NewInt(math.MaxInt)
	bigA.Add(bigA, bigB)
	return bigA.Cmp(bigMax) > 1
}

type MaxVal struct {
	maximums map[any]uint64
}

func MaxValue(a any) (uint64, error) {
	t := reflect.TypeOf(a)
	switch t.Kind() {
	case reflect.Int8:
		return math.MaxInt8, nil
	case reflect.Uint8:
		return math.MaxUint8, nil
	case reflect.Int16:
		return math.MaxInt16, nil
	case reflect.Uint16:
		return math.MaxUint16, nil
	case reflect.Int:
		return math.MaxInt, nil
	case reflect.Uint:
		return math.MaxUint, nil
	case reflect.Int32:
		return math.MaxInt32, nil
	case reflect.Uint32:
		return math.MaxUint32, nil
	case reflect.Int64:
		return math.MaxInt64, nil
	case reflect.Uint64:
		return math.MaxUint64, nil
	case reflect.Bool:
		return 1, nil
	default:
		return 0,
			errors.New("Type has larger maximum than a uint64, or is non-numeric: " +
				t.Kind().String())
	}

}
