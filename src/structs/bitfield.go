package structs

import (
	"math/big"
	"strconv"
	"strings"
)

type BitReadDirection bool

const (
	FromLeft  BitReadDirection = false // index 0 = most significant bit
	FromRight BitReadDirection = true  // index 0 = least significant bit
)

func (d BitReadDirection) String() string {
	if d {
		return "from the right"
	}
	return "from the left"
}

var _bitFieldReadingDirection BitReadDirection = FromRight

// SetReadDirection determines which direction bits are read from when
// inputting the index for a bit.
func SetReadDirection(newDirection BitReadDirection) {
	_bitFieldReadingDirection = newDirection
}

func GetReadDirection() BitReadDirection {
	return _bitFieldReadingDirection
}

type Bitfield8 uint8
type Bitfield16 = Uint16
type Bitfield32 = Uint32
type Bitfield64 = Uint64

type Bitfield128 struct {
	left  Bitfield64
	right Bitfield64
}

type Bitfield256 struct {
	left  Bitfield128
	right Bitfield128
}

type BitfieldBig big.Int

type BitArray []Bit
type BitArray8 [8]Bit
type BitArray16 [16]Bit
type BitArray32 [32]Bit
type BitArray64 [64]Bit
type BitArray128 [128]Bit
type BitArray256 [256]Bit

// Put sets the value of the bit at the given 0-indexed position from the left.
// Returns the resulting Bitfield8.
func (f Bitfield8) Put(index Bitfield8, value Bool) (Bitfield8, error) {
	if index < 0 || index > 7 {
		return f,
			IndexOutOfRangeError{Attempted: Int(index), MaxSafeIndex: 7}
	}
	if value {
		return f | (0b1000_0000 >> index), nil
	}
	val, _ := FlipBitL(f, index)
	return f & val, nil
}

// PutR sets the value of the bit at the given 0-indexed position from the right.
// Returns this Bitfield8.
func (f Bitfield8) PutR(index Bitfield8, value Bool) (Bitfield8, error) {
	if index < 0 || index > 7 {
		return f, IndexOutOfRangeError{Attempted: Int(index), MaxSafeIndex: 7}
	}
	if value {
		return f | (0b0000_0001 << index), nil
	}
	val, _ := FlipBitL(f, index)
	return f & val, nil
}

func BitAt(bitfield, index Bitfield8) (Bool, error) {
	if index < 0 || index > 7 {
		return false, IndexOutOfRangeError{Attempted: Int(index), MaxSafeIndex: 7}
	}
	return (bitfield & (0b1000_0000 >> index)) == 0b1000_0000>>index, nil
}

func BitAtInt(bitfield, index Bitfield8) (Int, error) {
	bit, err := BitAt(bitfield, index)
	if bit {
		return 1, err
	}
	return 0, err
}

func BitAtR(bitfield, index Bitfield8) (Bool, error) {
	if index < 0 || index > 7 {
		return false, IndexOutOfRangeError{Attempted: Int(index), MaxSafeIndex: 7}
	}
	return (bitfield & (0b0000_0001 << index)) == 0b0000_0001<<index, nil
}

func BitAtRInt(bitfield, index Bitfield8) (Int, error) {
	bit, err := BitAtR(bitfield, index)
	if bit {
		return 1, err
	}
	return 0, err
}

// FlipBitL flips the bit at the 0-indexed position from the left in value.
func FlipBitL(value, index Bitfield8) (Bitfield8, error) {
	if index < 0 || index > 7 {
		return value, IndexOutOfRangeError{Attempted: Int(index), MaxSafeIndex: 7}
	}
	newVal, err := BitAt(value, index)
	var isolated Bitfield8 = 0b1000_0000 >> index
	if newVal {
		return value & (0b1111_1111 ^ isolated), err
	}
	return value | isolated, err
}

// FlipBitR flips the bit at the 0-indexed position from the right in value.
func FlipBitR(value, index Bitfield8) (Bitfield8, error) {
	if index < 0 || index > 7 {
		return value, IndexOutOfRangeError{Attempted: Int(index), MaxSafeIndex: 7}
	}
	newVal, err := BitAtR(value, index)
	var isolated Bitfield8 = 0b0000_0001 << index
	if newVal {
		return value & (0b1111_1111 ^ isolated), err
	}
	return value | isolated, err
}

func (f Bitfield8) String() string {
	str := strconv.FormatUint(uint64(f), 2)      // binary formatting
	return strings.Repeat("0", 8-len(str)) + str // prepend leading 0s
}
