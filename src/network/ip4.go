package network

type IP4Address struct {
	host Host
	port Port
}

type IP4Packet struct {
	VersionAndIHL          uint8
	TOS                    uint8
	TotalLength            uint16
	Identification         uint16
	FlagsAndFragmentOffset uint16
	TTL                    uint8
	Protocol               uint8
	HeaderChecksum         uint16
	SourceAddress          uint32
	DestinationAddress     uint32
	Options                *IP4PacketOptions
	Data                   *IP4PacketData
}

type IP4PacketOptions []byte
type IP4PacketData []byte

func NewIP4(h Host, p Port) IP4Address {
	return IP4Address{port: p, host: h}
}

func IP4FromString(h string, p string) IP4Address {
	return IP4Address{host: NewHost(h), port: NewPort(p)}
}
