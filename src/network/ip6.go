package network

import "github.com/lfknudsen/golib/src/structs"

type HeaderIP6 struct {
	HeaderWord1   structs.Uint32
	PayloadLength structs.Uint16
	NextHeader    structs.Uint8
	HopLimit      structs.Uint8
}

type AddressIP6 [8]uint16
