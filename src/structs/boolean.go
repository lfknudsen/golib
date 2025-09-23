package structs

// Bool is a type that exists to be able to treat booleans as integers and vice versa;
// one aspect of C that I really miss in other languages.
// Integer conversions work on the concept that non-zero values are true.
type Bool bool

func (b Bool) Int() int   { return BoolToInt(bool(b)) }
func (b Bool) Bool() bool { return bool(b) }
func (b Bool) Not() Bool  { return !b }
func (b Bool) Integer() Integer {
	if b {
		return Integer(1)
	}
	return Integer(0)
}

func IntToBool(i int) bool {
	return i != 0
}

func BoolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}
