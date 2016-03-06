package pes

import "fmt"

const (
	StartCodePrefix uint32 = 0x000001
	MarkerBits      uint16 = 0x02
)

type Packet struct {
	CodePrefix   uint32
	StreamID     uint8
	PacketLength uint16
	Header       *Header
}

func newPacket(data []byte) (*Packet, error) {
	cp := uint32(data[0])<<8 | uint32(data[1])<<4 | uint32(data[2])
	if cp != StartCodePrefix {
		return nil, fmt.Errorf("Invalid PES packet, must start with proper code prefix, got %x expect %x", cp, StartCodePrefix)
	}
	p := &Packet{
		CodePrefix:   cp,
		StreamID:     data[3],
		PacketLength: uint16(data[4])<<8 | uint16(data[5])}

	if hasHeader(p.StreamID) {
		p.Header = newHeader(data)
	}
	return p, nil
}

func hasHeader(id uint8) bool {
	return id != 0xBC && id != 0xBE && id != 0xBF &&
		id != 0xF0 && id != 0xF1 && id != 0xFF &&
		id != 0xF2 && id != 0xF8
}

func IsAudio(id uint8) bool {
	return 0xC0 <= id && id <= 0xDF
}

func IsVideo(id uint8) bool {
	return 0xE0 <= id && id <= 0xEF
}
