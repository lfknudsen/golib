package network

type AddressIP4 struct {
	host Host
	port Port
}

type v4Packet struct {
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
	Options                *OptionsIP4
	Data                   *DataIP4
}

type OptionsIP4 []byte
type DataIP4 []byte

func NewIP4(h Host, p Port) AddressIP4 {
	return AddressIP4{port: p, host: h}
}

func IP4FromString(h string, p string) AddressIP4 {
	return AddressIP4{host: NewHost(h), port: NewPort(p)}
}
