package network

type TCPPacket struct {
	srcPort      Port
	dstPort      Port
	seq          uint32
	ack          uint32
	offsetAndRes uint8 // Unused
	flags        TCPFlags
	windowSize   uint16 // Unused
	checksum     uint16 // Unused
	urgentPtr    uint16 // Unused
	options      uint32 // Unused
}

func (p *TCPPacket) Init(localPort string) {
	p.srcPort = NewPort(localPort)
}
