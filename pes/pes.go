package pes

import "fmt"

const (
	StartCodePrefix uint32 = 0x000001
	MarkerBits      uint16 = 0x02
)

type Packet struct {
	CodePrefix                 uint32
	StreamID                   uint8
	PacketLength               uint16
	MarkerBits                 uint8
	ScramblingControl          uint8
	Priority                   bool
	DataAlignmentIndicator     bool
	Copyright                  bool
	Original                   bool
	ContainsPTS                bool
	ContainsDTS                bool
	ContainsESCR               bool
	ContainsESRate             bool
	ContainsDSMTrickMode       bool
	ContainsAdditionalCopyInfo bool
	ContainsCRC                bool
	ContainsExtension          bool
	HeaderLength               uint8
	PTS                        uint64
	DTS                        uint64
	//TODO Parse me
	ESCR                 uint64
	ESRate               uint32
	TrickModeControl     uint8
	FieldId              uint8
	IntraSliceRefresh    bool
	FrequencyTruncation  uint8
	RepCntrl             uint8
	AdditionalCopyInfo   uint8
	PreviousPESPacketCRC uint16
}

func NewPacket(data []byte) (*Packet, error) {

	cp := uint32(data[0])<<8 | uint32(data[1])<<4 | uint32(data[2])
	if cp != StartCodePrefix {
		return nil, fmt.Errorf("Invalid PES packet, must start with proper code prefix, got %x expect %x", cp, StartCodePrefix)
	}

	p := &Packet{
		CodePrefix: cp,
		StreamID:   data[3]}

	if p.StreamID != 0xBC && p.StreamID != 0xBE && p.StreamID != 0xBF &&
		p.StreamID != 0xF0 && p.StreamID != 0xF1 && p.StreamID != 0xFF &&
		p.StreamID != 0xF2 && p.StreamID != 0xF8 {

		p.CodePrefix = cp
		p.StreamID = data[3]
		p.PacketLength = uint16(data[4])<<8 | uint16(data[5])
		p.MarkerBits = data[6] & 0xC0 >> 6
		p.ScramblingControl = data[6] & 0x30 >> 4
		p.Priority = data[6]&0x08>>3 == 1
		p.DataAlignmentIndicator = data[6]&0x04>>2 == 1
		p.Copyright = data[6]&0x02>>1 == 1
		p.Original = data[6]&0x01 == 1
		p.ContainsPTS = data[7]&0x80>>7 == 1
		p.ContainsDTS = data[7]&0x40>>6 == 1
		p.ContainsESCR = data[7]&0x10>>4 == 1
		p.ContainsESRate = data[7]&0x10>>4 == 1
		p.ContainsDSMTrickMode = data[7]&0x08>>3 == 1
		p.ContainsAdditionalCopyInfo = data[7]&0x04>>2 == 1
		p.ContainsCRC = data[7]&0x02>>1 == 1
		p.ContainsExtension = data[7]&0x01 == 1
		p.HeaderLength = data[8]

		if p.ContainsPTS && !p.ContainsDTS {
			p.PTS = uint64(data[9]&0x0E>>1) << 30
			p.PTS = p.PTS | (uint64(data[10])<<8|uint64(data[11]))>>1<<14
			p.PTS = p.PTS | (uint64(data[12])<<8|uint64(data[13]))>>1
		}

		if p.ContainsPTS && p.ContainsDTS {
			p.PTS = uint64(data[9]&0x0E>>1) << 30
			p.PTS = p.PTS | (uint64(data[10])<<8|uint64(data[11]))>>1<<15
			p.PTS = p.PTS | (uint64(data[12])<<8|uint64(data[13]))>>1

			p.DTS = uint64(data[14]&0x0E>>1) << 30
			p.DTS = p.DTS | (uint64(data[15])<<8|uint64(data[16]))>>1<<15
			p.DTS = p.DTS | (uint64(data[17])<<8|uint64(data[18]))>>1
		}
	}
	return p, nil
}
