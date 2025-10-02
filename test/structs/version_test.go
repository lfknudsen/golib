package structs

import (
	"reflect"
	"strconv"
	"testing"

	. "github.com/lfknudsen/golib/src/structs"
)

func TestVersion_Bytes(t *testing.T) {
	type test struct {
		name string
		v    Version
		want []byte
	}
	var tests []test

	numbers := []byte{0, 1, 2, 255}
	for _, iNumber := range numbers {
		for _, jNumber := range numbers {
			for _, kNumber := range numbers {
				tests = append(tests, test{
					name: strconv.FormatUint(uint64(iNumber), 10) + "." +
						strconv.FormatUint(uint64(jNumber), 10) + "." +
						strconv.FormatUint(uint64(kNumber), 10),
					v:    Version{Major: iNumber, Minor: jNumber, Patch: kNumber},
					want: []byte{iNumber, jNumber, kNumber},
				})
			}
		}
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v.Bytes(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Bytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVersion_String(t *testing.T) {
	type test struct {
		name string
		v    Version
		want string
	}
	var tests []test

	numbers := []byte{0, 1, 2, 255}
	for _, iNumber := range numbers {
		for _, jNumber := range numbers {
			for _, kNumber := range numbers {
				name := strconv.FormatUint(uint64(iNumber), 10) + "." +
					strconv.FormatUint(uint64(jNumber), 10) + "." +
					strconv.FormatUint(uint64(kNumber), 10)
				tests = append(tests, test{
					name: name,
					v:    Version{Major: iNumber, Minor: jNumber, Patch: kNumber},
					want: name,
				})
			}
		}
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}
