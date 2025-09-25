package structs

// Bool is a type that exists to be able to treat booleans as integers and vice versa;
// one aspect of C that I really miss in other languages.
// Int conversions work on the concept that non-zero values are true.
type Bool bool
type Bit = int8

type IBooleanValue interface {
	IsTrue() Bool
	IsFalse() Bool
	Bool() Bool
}

const (
	FALSE Bit = iota
	TRUE
)

func (b Bool) Byte() Byte {
	if b {
		return 1
	}
	return 0
}

func (b Bool) Not() Bool { return !b }
func (b Bool) Int() Int {
	return BoolToInt(b)
}

func IntToBool(i Int) Bool {
	return i != 0
}

func BoolToInt(b Bool) Int {
	if b {
		return 1
	}
	return 0
}

func (b Bool) Equals(other Bool) Bool {
	return b == other
}

func (b Bool) CompareToInt(o Int) Int {
	return CompareInts(b.Int(), o)
}
