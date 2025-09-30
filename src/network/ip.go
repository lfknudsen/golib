package network

type IPAddress interface {
	Address() string
}

type Connection struct {
	localAddr  IPAddress
	remoteAddr IPAddress
}
