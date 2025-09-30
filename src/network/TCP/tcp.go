package TCP

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"slices"
	"strconv"

	. "github.com/lfknudsen/golib/src/logging"

	"github.com/lfknudsen/golib/src/network"
	. "github.com/lfknudsen/golib/src/structs"
)

type Header struct {
	SrcPort      network.Port // uint16
	DstPort      network.Port // uint16
	Seq          uint32       // Sequence number (relevant if Flags(SYN) is true)
	Ack          uint32       // Acknowledge number (relevant if Flags(ACK) is true)
	OffsetAndRes uint8        // Unused
	Flags        FlagField    // Bitfield (uint8)
	WindowSize   uint16       // Unused
	Checksum     uint16       // Unused
	UrgentPtr    uint16       // Unused
	Options      uint32       // Unused
}

func (p *Header) String() string {
	return fmt.Sprintf("Src: %s Dst: %s\nSeq: %d Ack: %d\nFlags: %s",
		p.SrcPort, p.DstPort, p.Seq, p.Ack, p.Flags.String())
}

func (p *Header) SetSourcePort(srcPort network.Port) {
	p.SrcPort = srcPort
}

func (p *Header) SetLocalAddr(local net.TCPAddr) {
	p.SrcPort = network.Port(local.Port)
}

func (p *Header) SetDestinationPort(dstPort network.Port) {
	p.DstPort = dstPort
}

func (p *Header) SetRemoteAddr(remote net.TCPAddr) {
	p.SrcPort = network.Port(remote.Port)
}

func (p *Header) SetSeq(seq uint32) {
	p.Seq = seq
}

func (p *Header) SetAck(ack uint32) {
	p.Ack = ack
}

func (p *Header) SetFlag(flag Flag, bit Bit) {
	bitfield, err := p.Flags.PutFlag(flag, bit != 0)
	FatalCheck(err)
	p.Flags = bitfield
}

func (p *Header) GetFlag(flag Flag) bool {
	return bool(p.Flags.At(flag))
}

func (p *Header) GetFlags() Bitfield8 {
	return Bitfield8(p.Flags)
}

func (p *Header) Write(w io.Writer) error {
	return binary.Write(w, binary.BigEndian, p)
}

func (p *Header) AllSet(flags ...Flag) bool {
	for _, flag := range flags {
		if !p.GetFlag(flag) {
			return false
		}
	}
	return true
}

func (p *Header) IsSyn() Bool {
	return p.Flags.At(SYN) && !p.Flags.At(ACK)
}

func (p *Header) IsSynAck() Bool {
	return p.Flags.At(SYN) && p.Flags.At(ACK)
}

func (p *Header) IsAck() Bool {
	return p.Flags.At(ACK) && !p.Flags.At(SYN)
}

func (p *Header) MarshalJSON() ([]byte, error) {
	head := []byte(`{"data":{`)
	src := []byte(`"SrcPort":"` + p.SrcPort.String() + `",`)
	dst := []byte(`"DstPort":"` + p.DstPort.String() + `",`)
	seq := []byte(`"Seq":"` + strconv.FormatUint(uint64(p.Seq), 10) + `",`)
	ack := []byte(`"Ack":"` + strconv.FormatUint(uint64(p.Ack), 10) + `",`)
	offset := []byte(`"OffsetAndRes":"` + strconv.FormatUint(uint64(p.OffsetAndRes), 10) + `",`)
	flags := []byte(`"Flags":"` + strconv.FormatUint(uint64(p.Flags), 10) + `",`)
	windowSize := []byte(`"WindowSize":"` + strconv.FormatUint(uint64(p.WindowSize), 10) + `",`)
	checksum := []byte(`"Checksum":"` + strconv.FormatUint(uint64(p.Checksum), 10) + `",`)
	urgentPtr := []byte(`"UrgentPtr":"` + strconv.FormatUint(uint64(p.UrgentPtr), 10) + `",`)
	options := []byte(`"Options":"` + strconv.FormatUint(uint64(p.Options), 10) + `"`)
	tail := []byte(`}}`)
	return slices.Concat(head, src, dst, seq, ack, offset, flags, windowSize,
		checksum, urgentPtr, options, tail), nil
}

func (p *Header) UnmarshalJSON(b []byte) error {
	fmt.Println("Inside Header.UnmarshalJSON")
	var data = make(map[string]map[string]string)
	err := json.Unmarshal(b, &data)
	fmt.Println("After unmarshalling into a map.")
	FatalCheck(err)
	dict := data["data"]
	src, err := strconv.ParseUint(dict["SrcPort"], 10, 16)
	dst, err := strconv.ParseUint(dict["DstPort"], 10, 16)
	seq, err := strconv.ParseUint(dict["Seq"], 10, 32)
	ack, err := strconv.ParseUint(dict["Ack"], 10, 32)
	offset, err := strconv.ParseUint(dict["OffsetAndRes"], 10, 8)
	flags, err := strconv.ParseUint(dict["Flags"], 10, 8)
	windowSize, err := strconv.ParseUint(dict["WindowSize"], 10, 16)
	checksum, err := strconv.ParseUint(dict["Checksum"], 10, 16)
	urgentPtr, err := strconv.ParseUint(dict["UrgentPtr"], 10, 16)
	options, err := strconv.ParseUint(dict["Options"], 10, 32)

	p.SrcPort = network.Port(src)
	p.DstPort = network.Port(dst)
	p.Seq = uint32(seq)
	p.Ack = uint32(ack)
	p.OffsetAndRes = uint8(offset)
	p.Flags = FlagField(flags)
	p.WindowSize = uint16(windowSize)
	p.Checksum = uint16(checksum)
	p.UrgentPtr = uint16(urgentPtr)
	p.Options = uint32(options)
	return err
}
