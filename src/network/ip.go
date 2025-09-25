package network

type IPAddress interface {
	Address() string
}

type IPConnection struct {
	localAddr  IPAddress
	remoteAddr IPAddress
}
