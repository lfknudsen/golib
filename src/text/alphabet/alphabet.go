package alphabet

import (
	"strings"

	"github.com/lfknudsen/golib/src/text"
)

type Alphabet interface {
	All() text.String
	Upper() text.String
	Lower() text.String
	Valid(text.String) bool
	Contains(text.String) bool
}

const (
	Empty       string = ""
	LatinLower         = "abcdefghijklmnopqrstuvwxyz"
	LatinUpper         = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	Danish             = "abcdefghijklmnopqrstuvwxyzæøåABCDEFGHIJKLMNOPQRSTUVWXYZÆØÅ"
	DanishLower        = "abcdefghijklmnopqrstuvwxyzæøå"
	DanishUpper        = "ABCDEFGHIJKLMNOPQRSTUVWXYZÆØÅ"
	Symbols            = " ,.;:-_'^¨~|`=?|}][{€$£@!\"#¤%&/()€<>\\+*"
	Digits             = "0123456789"
)

type DanishAlphabet text.String

func (d DanishAlphabet) Upper() text.String {
	return DanishUpper
}

func (d DanishAlphabet) Lower() text.String {
	return DanishLower
}

func (d DanishAlphabet) Valid(s text.String) bool {
	for _, ch := range s {
		if strings.IndexRune(Danish, ch) < 0 &&
			strings.IndexRune(Symbols, ch) < 0 &&
			strings.IndexRune(Digits, ch) < 0 {
			return false
		}
	}
	return true
}
