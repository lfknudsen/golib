package network

type IP6Header struct {
	HeaderWord1   Uint32
	PayloadLength Uint16
	NextHeader    Uint8
	HopLimit      Uint8
}

type IP6Address [8]uint16
