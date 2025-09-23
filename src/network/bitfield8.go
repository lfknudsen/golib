package network

type TCPFlags uint8

func (b TCPFlags) ACK() bool {
	return b&0b0001_0000 == 0b0001_0000
}

func (b TCPFlags) SetACK(val bool) {
	if val {
		b |= 0b0001_0000
	} else {
		b &= 0b1110_1111
	}
}

func (b TCPFlags) ToggleACK() {
	b ^= 0b0001_0000
}

func (b TCPFlags) SYN() bool {
	return b&0b0000_0010 == 0b0000_0010
}

func (b TCPFlags) SetSYN(val bool) {
	if val {
		b |= 0b0000_0010
	} else {
		b &= 0b1111_1101
	}
}

func (b TCPFlags) ToggleSYN() {
	b ^= 0b0000_0010
}

func (b TCPFlags) FIN() bool {
	return b&0b0000_0001 == 0b0000_0001
}

func (b TCPFlags) SetFIN(val bool) {
	if val {
		b |= 0b0000_0001
	} else {
		b &= 0b1111_1110
	}
}

func (b TCPFlags) ToggleFIN() {
	b ^= 0b0000_0001
}
