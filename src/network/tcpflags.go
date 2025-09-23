package network

import (
	"errors"
	"math"
	"reflect"
	"strconv"
	"strings"

	"github.com/lfknudsen/golib/src/logging"
)

type TCPFlags uint8

func (f TCPFlags) Int() int {
	return int(f)
}

func (f TCPFlags) Uint() uint {
	return uint(f)
}

func (f TCPFlags) Uint8() uint8 {
	return uint8(f)
}

func (f TCPFlags) String() string {
	return strconv.FormatUint(uint64(f), 2)
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

func (f TCPFlags) AndI(operand uint8) TCPFlags {
	return TCPFlags(f.Uint8() & operand)
}

func (f TCPFlags) OrI(operand uint8) TCPFlags {
	return TCPFlags(f.Uint8() | operand)
}

func (f TCPFlags) XorI(operand uint8) TCPFlags {
	return TCPFlags(f.Uint8() ^ operand)
}

func (f TCPFlags) And(operand TCPFlags) TCPFlags {
	return TCPFlags(f.Uint8() & operand.Uint8())
}

func (f TCPFlags) Or(operand TCPFlags) TCPFlags {
	return TCPFlags(f.Uint8() | operand.Uint8())
}

func (f TCPFlags) Xor(operand TCPFlags) TCPFlags {
	return TCPFlags(f.Uint8() ^ operand.Uint8())
}

// At returns the value of the bit at the given 0-indexed position from the left.
func (f TCPFlags) At(index uint8) (bool, error) {
	return BitAt(f.Uint8(), index)
}

// AtR returns the value of the bit at the given 0-indexed position from the right.
func (f TCPFlags) AtR(index uint8) (bool, error) {
	return BitAtR(f.Uint8(), index)
}

func (f TCPFlags) PutFlag(flag TCPFlag, value bool) (TCPFlags, error) {
	return f.Put(flag.Uint8(), value)
}

func (f TCPFlags) SetFlag(flag TCPFlag) (TCPFlags, error) {
	return f.Put(flag.Uint8(), true)
}

func (f TCPFlags) UnsetFlag(flag TCPFlag) (TCPFlags, error) {
	return f.Put(flag.Uint8(), false)
}

// Put sets the value of the bit at the given 0-indexed position from the left.
// Returns the resulting TCPFlags.
func (f TCPFlags) Put(index uint8, value bool) (TCPFlags, error) {
	if index < 0 || index > 7 {
		return f, logging.IndexOutOfRange(index, 0, 7)
	}
	if value {
		return TCPFlags(f.Uint8() | (0b1000_0000 >> index)), nil
	}
	val, _ := FlipBit(f.Uint8(), index)
	return TCPFlags(f.Uint8() & val), nil
}

// PutR sets the value of the bit at the given 0-indexed position from the right.
// Returns this TCPFlags.
func (f TCPFlags) PutR(index uint8, value bool) (TCPFlags, error) {
	if index < 0 || index > 7 {
		return f, logging.IndexOutOfRange(index, 0, 7)
	}
	if value {
		return TCPFlags(f.Uint8() | (0b0000_0001 << index)), nil
	}
	val, _ := FlipBit(f.Uint8(), index)
	return TCPFlags(f.Uint8() & val), nil
}

func BitAt(bitfield, index uint8) (bool, error) {
	if index < 0 || index > 7 {
		return false, logging.IndexOutOfRange(index, 0, 7)
	}
	return (bitfield & (0b1000_0000 >> index)) == 0b1000_0000>>index, nil
}

func BitAtInt(bitfield, index uint8) (int, error) {
	bit, err := BitAt(bitfield, index)
	if bit {
		return 1, err
	}
	return 0, err
}

func BitAtR(bitfield, index uint8) (bool, error) {
	if index < 0 || index > 7 {
		return false, logging.IndexOutOfRange(index, 0, 7)
	}
	return (bitfield & (0b0000_0001 << index)) == 0b0000_0001<<index, nil
}

func BitAtRInt(bitfield, index uint8) (int, error) {
	bit, err := BitAtR(bitfield, index)
	if bit {
		return 1, err
	}
	return 0, err
}

// FlipBit flips the bit at the 0-indexed position from the left in value.
func FlipBit(value uint8, index uint8) (uint8, error) {
	if index < 0 || index > 7 {
		return value, logging.IndexOutOfRange(index, 0, 7)
	}
	newVal, err := BitAt(value, index)
	var isolated uint8 = 0b1000_0000 >> index
	if newVal {
		return value & (0b1111_1111 ^ isolated), err
	}
	return value | isolated, err

}

// FlipBitR flips the bit at the 0-indexed position from the right in value.
func FlipBitR(value uint8, index uint8) (uint8, error) {
	if index < 0 || index > 7 {
		return value, logging.IndexOutOfRange(index, 0, 7)
	}
	newVal, err := BitAtR(value, index)
	var isolated uint8 = 0b0000_0001 << index
	if newVal {
		return value & (0b1111_1111 ^ isolated), err
	}
	return value | isolated, err
}

func (f TCPFlags) CWR() bool {
	return f&0b1000_0000 == 0b1000_0000
}

func (f TCPFlags) SetCWR(val bool) {
	if val {
		f |= 0b1000_0000
	} else {
		f &= 0b0111_1111
	}
}

func (f TCPFlags) FlipCWR() {
	f ^= 0b1000_0000
}

func (f TCPFlags) ECE() bool {
	return f&0b0100_0000 == 0b0100_0000
}

func (f TCPFlags) SetECE(val bool) {
	if val {
		f |= 0b0100_0000
	} else {
		f &= 0b1011_1111
	}
}

func (f TCPFlags) FlipECE() {
	f ^= 0b0100_0000
}

func (f TCPFlags) URG() bool {
	return f&0b0010_0000 == 0b0010_0000
}

func (f TCPFlags) SetURG(val bool) {
	if val {
		f |= 0b0010_0000
	} else {
		f &= 0b1101_1111
	}
}

func (f TCPFlags) FlipURG() {
	f ^= 0b0010_0000
}

func (f TCPFlags) ACK() bool {
	return f&0b0001_0000 == 0b0001_0000
}

func (f TCPFlags) SetACK(val bool) {
	if val {
		f |= 0b0001_0000
	} else {
		f &= 0b1110_1111
	}
}

func (f TCPFlags) FlipACK() {
	f ^= 0b0001_0000
}

func (f TCPFlags) PSH() bool {
	return f&0b0000_1000 == 0b0000_1000
}

func (f TCPFlags) SetPSH(val bool) {
	if val {
		f |= 0b0000_1000
	} else {
		f &= 0b1111_0111
	}
}

func (f TCPFlags) FlipPSH() {
	f ^= 0b0000_1000
}

func (f TCPFlags) RST() bool {
	return f&0b0000_0100 == 0b0000_0100
}

func (f TCPFlags) SetRST(val bool) {
	if val {
		f |= 0b0000_0100
	} else {
		f &= 0b1111_1011
	}
}

func (f TCPFlags) FlipRST() {
	f ^= 0b0000_0100
}

func (f TCPFlags) SYN() bool {
	return f&0b0000_0010 == 0b0000_0010
}

func (f TCPFlags) SetSYN(val bool) {
	if val {
		f |= 0b0000_0010
	} else {
		f &= 0b1111_1101
	}
}

func (f TCPFlags) FlipSYN() {
	f ^= 0b0000_0010
}

func (f TCPFlags) FIN() bool {
	return f&0b0000_0001 == 0b0000_0001
}

func (f TCPFlags) SetFIN(val bool) {
	if val {
		f |= 0b0000_0001
	} else {
		f &= 0b1111_1110
	}
}

func (f TCPFlags) FlipFIN() {
	f ^= 0b0000_0001
}

type TCPFlag uint8

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
	ERROR TCPFlag = math.MaxUint8
)

func (f TCPFlag) String() string {
	switch f {
	case CWR:
		return "CWR"
	case ECE:
		return "ECE"
	case URG:
		return "URG"
	case ACK:
		return "ACK"
	case PSH:
		return "PSH"
	case RST:
		return "RST"
	case SYN:
		return "SYN"
	case FIN:
		return "FIN"
	default:
		return ""
	}
}

func StringToTCPFlag(str string) (TCPFlag, error) {
	switch strings.ToUpper(str) {
	case "CWR":
		return CWR, nil
	case "ECE":
		return ECE, nil
	case "URG":
		return URG, nil
	case "ACK":
		return ACK, nil
	case "PSH":
		return PSH, nil
	case "RST":
		return RST, nil
	case "SYN":
		return SYN, nil
	case "FIN":
		return FIN, nil
	}
	return ERROR, logging.ConversionError{str, TCPFlag(0)}
}

func (f TCPFlag) Uint8() uint8 { return uint8(f) }
