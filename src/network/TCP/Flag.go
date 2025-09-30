package TCP

import (
	"reflect"

	"github.com/lfknudsen/golib/src/structs"
)

type Flag structs.Uint8

const FlagMAX uint8 = 7

const FlagMIN uint8 = 0

func IsValidInt(val int) bool {
	return int(FlagMIN) <= val && val <= int(FlagMAX)
}

const (
	CWR Flag = iota
	ECE
	URG
	ACK
	PSH
	RST
	SYN
	FIN
)

var FlagToString = map[Flag]string{
	CWR: "CWR",
	ECE: "ECE",
	URG: "URG",
	ACK: "ACK",
	PSH: "PSH",
	RST: "RST",
	SYN: "SYN",
	FIN: "FIN",
}

var StringToFlag = map[string]Flag{
	"CWR": CWR,
	"ECE": ECE,
	"URG": URG,
	"ACK": ACK,
	"PSH": PSH,
	"RST": RST,
	"SYN": SYN,
	"FIN": FIN,
}

func (f Flag) String() string {
	return FlagToString[f]
}

func StringToTCPFlag(str string) Flag {
	return StringToFlag[str]
}

func (f Flag) Underlying() reflect.Type {
	return reflect.TypeOf(Flag(0))
}
