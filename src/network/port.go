package network

import (
	"math/big"
	"net"
	"strconv"

	"github.com/lfknudsen/golib/src/logging/v2"
)

type Port uint16

func (p Port) String() string {
	return strconv.FormatUint(uint64(p), 10)
}

func FromString(s string) Port {
	val, err := strconv.ParseUint(s, 10, 16)
	ErrorCheck(err)
	return Port(val)
}

func (p Port) TCPAddr() net.TCPAddr {
	return net.TCPAddr{Port: p.Int()}
}

func FromTCPAddr(addr net.TCPAddr) Port {
	return Port(addr.Port)
}

func FromAddr(addr net.Addr) Port {
	str := addr.String()
	_, p, err := net.SplitHostPort(str)
	ErrorCheck(err)
	return FromString(p)
}

func (p Port) Dial(network string) (net.Conn, error) {
	return net.Dial(network, p.String())
}

func (p Port) DialTCP(network string, remoteAddr *net.TCPAddr) (net.Conn, error) {
	localAddr := p.TCPAddr()
	return net.DialTCP(network, &localAddr, remoteAddr)
}

func (p Port) Listen(network string) (net.Listener, error) {
	return net.Listen(network, p.String())
}

func (p Port) ListenTCP(network string) (net.Listener, error) {
	localAddr := p.TCPAddr()
	return net.ListenTCP(network, &localAddr)
}

func (p Port) Int() int {
	return int(p)
}

func (p Port) Uint() uint {
	return uint(p)
}

func (p Port) Uint16() uint16 {
	return uint16(p)
}

func (p Port) Int32() int32 {
	return int32(p)
}

func (p Port) Uint32() uint16 {
	return uint16(p)
}

func (p Port) Int64() int64 {
	return int64(p)
}

func (p Port) Uint64() uint64 {
	return uint64(p)
}

func (p Port) BigInt() big.Int {
	bi := big.Int{}
	bi.SetUint64(uint64(p))
	return bi
}

func FromBigInt(bi big.Int) Port {
	return Port(bi.Uint64())
}
