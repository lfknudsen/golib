package network

type IP4 struct {
	host Host
	port Port
}

func NewIP4(h Host, p Port) IP4 {
	return IP4{port: p, host: h}
}

func IP4FromString(h string, p string) IP4 {
	return IP4{host: NewHost(h), port: NewPort(p)}
}
