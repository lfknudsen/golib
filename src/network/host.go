package network

type Host string

func (h *Host) String() string {
	return string(*h)
}

func (h *Host) Join(p Port) IP4Address {
	return NewIP4(*h, p)
}

func NewHost(str string) Host {
	return Host(str)
}
