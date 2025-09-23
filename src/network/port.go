package network

import (
	"math/big"
	"net"
	"strconv"
	"strings"

	. "github.com/lfknudsen/golib/src/logging"
)

type Port uint16

func (p Port) String() string {
	str := strconv.FormatUint(uint64(p), 10)
	return str + strings.Repeat("0", 4-len(str))
}

func NewPort(s string) Port {
	val, err := strconv.ParseUint(s, 10, 16)
	ErrorCheck(err)
	return Port(val)
}

func (p Port) Join(h Host) IP4 {
	return NewIP4(h, p)
}

func (p Port) TCPAddr() net.TCPAddr {
	return net.TCPAddr{Port: p.Int()}
}

func PortFromTCP(addr net.TCPAddr) Port {
	return Port(addr.Port)
}

func PortFromAddr(addr net.Addr) Port {
	str := addr.String()
	_, p, err := net.SplitHostPort(str)
	ErrorCheck(err)
	return NewPort(p)
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

func (p Port) Bytes() []byte {
	return strconv.AppendUint([]byte{}, p.Uint64(), 10)
}

func (p Port) Digits() []uint8 {
	return strconv.AppendUint([]uint8{}, p.Uint64(), 10)
}

func (p Port) Digits32() []int {
	digits := p.Digits()
	return []int{int(digits[0]), int(digits[1]), int(digits[2]), int(digits[3])}
}

func (p Port) BigInt() big.Int {
	bi := big.Int{}
	bi.SetUint64(uint64(p))
	return bi
}

func FromBigInt(bi big.Int) Port {
	return Port(bi.Uint64())
}
