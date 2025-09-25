package network

import (
	"errors"
	"fmt"
	"math"
	"reflect"
	"strconv"
	"strings"

	"github.com/lfknudsen/golib/src/logging"
)

type Bool bool
type IBool interface {
	Int() int
	Boolean() bool
	Bool() Bool
}

type Bit = Bool

type Boolean Bool

func (b *Boolean) Boolean(val int) {
	if val == 0 {
		*b = false
	} else {
		*b = true
	}
}

func (b Bool) Int() int {
	if b {
		return 1
	}
	return 0
}

type Uint8 uint8
type Bitfield8 = Uint8
type TCPFlag Uint8
type TCPFlags = Bitfield8

func (f TCPFlags) Bool() bool {
	return f != 0
}

func (f TCPFlags) Byte() byte {
	if f.Bool() {
		return 1
	}
	return 0
}

func (f TCPFlags) Int() int {
	return int(f)
}

func (f TCPFlags) Uint() uint {
	return uint(f)
}

func (f TCPFlags) Uint8() uint8 {
	return uint8(f)
}

func TCPFlagsFromString(value string) (TCPFlags, error) {
	if len(value) != 8 {
		return TCPFlags(0),
			errors.New(
				"to set all values of the bitfield at once," +
					" the input string must be 8 bytes long")
	}
	val, err := strconv.ParseUint(value, 2, 8)
	if err != nil {
		return TCPFlags(0), err
	}
	return TCPFlags(val), nil
}

func (f Bitfield8) String() string {
	str := strconv.FormatUint(uint64(f), 2)      // binary formatting
	return strings.Repeat("0", 8-len(str)) + str // prepend leading 0s
}

// Put sets the value of the bit at the given 0-indexed position from the left.
// Returns the resulting Bitfield8.
func (f Bitfield8) Put(index Uint8, value Bool) (Bitfield8, error) {
	if index < 0 || index > 7 {
		return f, errors.New("index out of range")
	}
	if value {
		return f | (0b1000_0000 >> index), nil
	}
	val, _ := FlipBit(f, index)
	return f & val, nil
}

// FlipBit flips the bit at the 0-indexed position from the left in value.
func FlipBit(value Bitfield8, index Uint8) (Bitfield8, error) {
	newVal, err := BitAt(value, index)
	if err != nil {
		return value, err
	}
	var isolated Bitfield8 = 0b1000_0000 >> index
	if newVal {
		return value & (0b1111_1111 ^ isolated), err
	}
	return value | isolated, err
}

func BitAt(bitfield Bitfield8, index TCPFlag) (Bool, error) {
	if index < 0 || index > 7 {
		return false, logging.IndexOutOfRange{index, 0, 7}
	}
	return (bitfield & (0b1000_0000 >> index)) == 0b1000_0000>>index, nil
}

func (f TCPFlags) AndI(operand Uint8) TCPFlags {
	return f & operand
}

func (f TCPFlags) OrI(operand Uint8) TCPFlags {
	return f | operand
}

func (f TCPFlags) XorI(operand Uint8) TCPFlags {
	return f ^ operand
}

func (f TCPFlags) And(operand TCPFlags) TCPFlags {
	return f & operand
}

func (f TCPFlags) Or(operand TCPFlags) TCPFlags {
	return f | operand
}

func (f TCPFlags) Xor(operand TCPFlags) TCPFlags {
	return f ^ operand
}

// At returns the value of the bit at the given 0-indexed position from the left.
func (f Uint8) At(index TCPFlag) (Bool, error) {
	return BitAt(f, index)
}

func (f TCPFlags) Get(index uint8) bool {

}

// AtR returns the value of the bit at the given 0-indexed position from the right.
func (f TCPFlags) AtR(index Uint8) (Bool, error) {
	return BitAtR(f, index)
}

func (f TCPFlags) PutFlag(flag Uint8, value Bool) (TCPFlags, error) {
	return f.Put(flag, value)
}

func (f TCPFlags) SetFlag(flag Uint8) (TCPFlags, error) {
	return f.Put(flag, true)
}

func (f TCPFlags) UnsetFlag(flag Uint8) (TCPFlags, error) {
	return f.Put(flag, false)
}

func (f TCPFlags) CWR() Bool {
	return f&0b1000_0000 == 0b1000_0000
}

func (f TCPFlags) SetCWR(val Bool) TCPFlags {
	if val {
		f |= 0b1000_0000
	} else {
		f &= 0b0111_1111
	}
	return f
}

func (f TCPFlags) FlipCWR() TCPFlags {
	f ^= 0b1000_0000
	return f
}

func (f TCPFlags) ECE() Bool {
	return f&0b0100_0000 == 0b0100_0000
}

func (f TCPFlags) SetECE(val Bool) TCPFlags {
	if val {
		f |= 0b0100_0000
	} else {
		f &= 0b1011_1111
	}
	return f
}

func (f TCPFlags) FlipECE() TCPFlags {
	f ^= 0b0100_0000
	return f
}

func (f TCPFlags) URG() Bool {
	return f&0b0010_0000 == 0b0010_0000
}

func (f TCPFlags) SetURG(val Bool) TCPFlags {
	if val {
		f |= 0b0010_0000
	} else {
		f &= 0b1101_1111
	}
	return f
}

func (f TCPFlags) FlipURG() TCPFlags {
	f ^= 0b0010_0000
	return f
}

func (f TCPFlags) ACK() Bool {
	return f&0b0001_0000 == 0b0001_0000
}

func (f TCPFlags) SetACK(val Bool) TCPFlags {
	if val {
		f |= 0b0001_0000
	} else {
		f &= 0b1110_1111
	}
	return f
}

func (f TCPFlags) FlipACK() TCPFlags {
	f ^= 0b0001_0000
	return f
}

func (f TCPFlags) PSH() Bool {
	return f&0b0000_1000 == 0b0000_1000
}

func (f TCPFlags) SetPSH(val Bool) TCPFlags {
	if val {
		f |= 0b0000_1000
	} else {
		f &= 0b1111_0111
	}
	return f
}

func (f TCPFlags) FlipPSH() TCPFlags {
	f ^= 0b0000_1000
	return f
}

func (f TCPFlags) RST() Bool {
	return f&0b0000_0100 == 0b0000_0100
}

func (f TCPFlags) SetRST(val Bool) TCPFlags {
	if val {
		f |= 0b0000_0100
	} else {
		f &= 0b1111_1011
	}
	return f
}

func (f TCPFlags) FlipRST() TCPFlags {
	f ^= 0b0000_0100
	return f
}

func (f TCPFlags) SYN() Bool {
	return f&0b0000_0010 == 0b0000_0010
}

func (f TCPFlags) SetSYN(val Bool) TCPFlags {
	if val {
		f |= 0b0000_0010
	} else {
		f &= 0b1111_1101
	}
	return f
}

func (f TCPFlags) FlipSYN() TCPFlags {
	f ^= 0b0000_0010
	return f
}

func (f TCPFlags) FIN() Bool {
	return f&0b0000_0001 == 0b0000_0001
}

func (f TCPFlags) SetFIN(val Bool) TCPFlags {
	if val {
		f |= 0b0000_0001
	} else {
		f &= 0b1111_1110
	}
	return f
}

func (f TCPFlags) FlipFIN() TCPFlags {
	f ^= 0b0000_0001
	return f
}

func (f TCPFlags) SetFlags(flag TCPFlag, value Bit) (TCPFlags, error) {
	return f.Put(Bitfield8(flag), value)
}

func (f TCPFlag) Underlying() reflect.Type {
	return reflect.TypeOf(TCPFlag(0))
}

const (
	CWR TCPFlag = iota
	ECE
	URG
	ACK
	PSH
	RST
	SYN
	FIN
)

var FlagToString = map[TCPFlag]string{
	CWR: "CWR",
	ECE: "ECE",
	URG: "URG",
	ACK: "ACK",
	PSH: "PSH",
	RST: "RST",
	SYN: "SYN",
	FIN: "FIN",
}

var StringToFlag = map[string]TCPFlag{
	"CWR": CWR,
	"ECE": ECE,
	"URG": URG,
	"ACK": ACK,
	"PSH": PSH,
	"RST": RST,
	"SYN": SYN,
	"FIN": FIN,
}

func (f TCPFlag) String() string {
	return FlagToString[f]
}

func StringToTCPFlag(str string) TCPFlag {
	return StringToFlag[str]
}
