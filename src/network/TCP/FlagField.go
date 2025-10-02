package TCP

import (
	"errors"
	"strconv"
	"strings"

	. "github.com/lfknudsen/golib/src/structs"
)

type FlagField uint8

func (f FlagField) Bool() bool {
	return f != 0
}

func (f FlagField) Byte() byte {
	if f.Bool() {
		return 1
	}
	return 0
}

func (f FlagField) Int() int {
	return int(f)
}

func (f FlagField) Uint() uint {
	return uint(f)
}

func (f FlagField) Uint8() uint8 {
	return uint8(f)
}

func TCPFlagsFromString(value string) (FlagField, error) {
	if len(value) != 8 {
		return FlagField(0),
			errors.New(
				"to set all values of the bitfield at once," +
					" the input string must be 8 bytes long")
	}
	val, err := strconv.ParseUint(value, 2, 8)
	if err != nil {
		return FlagField(0), err
	}
	return FlagField(val), nil
}

func (f FlagField) String() string {
	str := strconv.FormatUint(uint64(f), 2)      // binary formatting
	return strings.Repeat("0", 8-len(str)) + str // prepend leading 0s
}

// Put sets the value of the bit at the given 0-indexed position from the left.
// Returns the resulting FlagField.
func (f FlagField) Put(index Flag, value Bool) (FlagField, error) {
	if index < 0 || index > 7 {
		return f, errors.New("index out of range")
	}
	if value {
		return f | (0b1000_0000 >> index), nil
	}
	val, _ := FlipBit(f, Uint8(index))
	return f & FlagField(val), nil
}

// FlipBit flips the bit at the 0-indexed position from the left in value.
func FlipBit(value FlagField, index Uint8) (FlagField, error) {
	newVal, err := BitAt(Bitfield8(value), Bitfield8(index))
	if err != nil {
		return value, err
	}
	var isolated FlagField = 0b1000_0000 >> index
	if newVal {
		return value & (0b1111_1111 ^ isolated), err
	}
	return value | isolated, err
}

func (f FlagField) AndI(operand Uint8) FlagField {
	return f & FlagField(operand)
}

func (f FlagField) OrI(operand Uint8) FlagField {
	return f | FlagField(operand)
}

func (f FlagField) XorI(operand Uint8) FlagField {
	return f ^ FlagField(operand)
}

func (f FlagField) And(operand FlagField) FlagField {
	return f & operand
}

func (f FlagField) Or(operand FlagField) FlagField {
	return f | operand
}

func (f FlagField) Xor(operand FlagField) FlagField {
	return f ^ operand
}

// At returns the value of the bit at the given 0-indexed position from the left.
func (f FlagField) At(index Flag) Bool {
	result, _ := BitAt(Bitfield8(f), Bitfield8(index))
	return result
}

func (f FlagField) Get(index uint8) (bool, error) {
	if index < FlagMIN || index > FlagMAX {
		return false, IndexOutOfRangeError{
			Attempted:    Int(index),
			MaxSafeIndex: Int(FlagMAX),
		}
	}
	return bool(f.At(Flag(index))), nil
}

// AtR returns the value of the bit at the given 0-indexed position from the right.
func (f FlagField) AtR(index Uint8) (Bool, error) {
	return BitAtR(Bitfield8(f), Bitfield8(index))
}

func (f FlagField) PutFlag(flag Flag, value Bool) (FlagField, error) {
	return f.Put(flag, value)
}

func (f FlagField) SetFlag(flag Flag) (FlagField, error) {
	return f.Put(flag, true)
}

func (f FlagField) UnsetFlag(flag Flag) (FlagField, error) {
	return f.Put(flag, false)
}

func (f FlagField) CWR() Bool {
	return f&0b1000_0000 == 0b1000_0000
}

func (f FlagField) SetCWR(val Bool) FlagField {
	if val {
		f |= 0b1000_0000
	} else {
		f &= 0b0111_1111
	}
	return f
}

func (f FlagField) FlipCWR() FlagField {
	f ^= 0b1000_0000
	return f
}

func (f FlagField) ECE() Bool {
	return f&0b0100_0000 == 0b0100_0000
}

func (f FlagField) SetECE(val Bool) FlagField {
	if val {
		f |= 0b0100_0000
	} else {
		f &= 0b1011_1111
	}
	return f
}

func (f FlagField) FlipECE() FlagField {
	f ^= 0b0100_0000
	return f
}

func (f FlagField) URG() Bool {
	return f&0b0010_0000 == 0b0010_0000
}

func (f FlagField) SetURG(val Bool) FlagField {
	if val {
		f |= 0b0010_0000
	} else {
		f &= 0b1101_1111
	}
	return f
}

func (f FlagField) FlipURG() FlagField {
	f ^= 0b0010_0000
	return f
}

func (f FlagField) ACK() Bool {
	return f&0b0001_0000 == 0b0001_0000
}

func (f FlagField) SetACK(val Bool) FlagField {
	if val {
		f |= 0b0001_0000
	} else {
		f &= 0b1110_1111
	}
	return f
}

func (f FlagField) FlipACK() FlagField {
	f ^= 0b0001_0000
	return f
}

func (f FlagField) PSH() Bool {
	return f&0b0000_1000 == 0b0000_1000
}

func (f FlagField) SetPSH(val Bool) FlagField {
	if val {
		f |= 0b0000_1000
	} else {
		f &= 0b1111_0111
	}
	return f
}

func (f FlagField) FlipPSH() FlagField {
	f ^= 0b0000_1000
	return f
}

func (f FlagField) RST() Bool {
	return f&0b0000_0100 == 0b0000_0100
}

func (f FlagField) SetRST(val Bool) FlagField {
	if val {
		f |= 0b0000_0100
	} else {
		f &= 0b1111_1011
	}
	return f
}

func (f FlagField) FlipRST() FlagField {
	f ^= 0b0000_0100
	return f
}

func (f FlagField) SYN() Bool {
	return f&0b0000_0010 == 0b0000_0010
}

func (f FlagField) SetSYN(val Bool) FlagField {
	if val {
		f |= 0b0000_0010
	} else {
		f &= 0b1111_1101
	}
	return f
}

func (f FlagField) FlipSYN() FlagField {
	f ^= 0b0000_0010
	return f
}

func (f FlagField) FIN() Bool {
	return f&0b0000_0001 == 0b0000_0001
}

func (f FlagField) SetFIN(val Bool) FlagField {
	if val {
		f |= 0b0000_0001
	} else {
		f &= 0b1111_1110
	}
	return f
}

func (f FlagField) FlipFIN() FlagField {
	f ^= 0b0000_0001
	return f
}

func (f FlagField) SetFlags(flag Flag, value Bit) (FlagField, error) {
	return f.Put(flag, value != 0)
}
