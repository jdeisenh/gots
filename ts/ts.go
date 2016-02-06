package ts

import "fmt"

const (
	PacketSize                int    = 188
	SyncByte                  uint8  = 0x47
	ProgramAssociationTableID uint16 = 0x0000
)

type Packet struct {
	SyncByte                   uint8
	TransportErrorIndicator    bool
	PayloadUnitStartIndicator  bool
	TransportPriority          bool
	PID                        uint16
	TransportScramblingControl uint8
	ContainsAdaptationField    bool
	ContainsPayload            bool
	ContinuityCounter          uint8
	AdaptationField            *AdaptationField
	Payload                    []byte
}

func NewPacket(data []byte) (*Packet, error) {
	if data[0] != SyncByte {
		return nil, fmt.Errorf("Invalid TS packet, must start with sync byte, got %x expect %x", data[0], SyncByte)
	}
	p := &Packet{
		SyncByte:                  data[0],
		TransportErrorIndicator:   data[1]&0x80>>7 == 1,
		PayloadUnitStartIndicator: data[1]&0x40>>6 == 1,
		TransportPriority:         data[1]&0x20>>5 == 1,
		PID:                       uint16(data[1]&0x1F)<<8 | uint16(data[2]),
		TransportScramblingControl: data[3] & 0xC0,
		ContainsAdaptationField:    data[3]&0x20>>5 == 1,
		ContainsPayload:            data[3]&0x10>>4 == 1,
		ContinuityCounter:          data[3] & 0xf}

	if p.ContainsAdaptationField {
		p.AdaptationField = newAdaptationField(data[4 : 5+int(data[4])])
	}
	if p.ContainsPayload {
		if p.ContainsAdaptationField {
			p.Payload = data[5+p.AdaptationField.AdaptationFieldLength : PacketSize]
		} else {
			p.Payload = data[4:PacketSize]
		}
	}
	return p, nil
}

func (p Packet) hasProgramMapTable(pat *ProgramAssociationTable) bool {
	if pat != nil {
		for _, r := range pat.Programs {
			if r.ProgramMapPID == p.PID {
				return true
			}
		}
	}
	return false
}

func (p Packet) hasProgramAssociationTable() bool {
	if p.PID == ProgramAssociationTableID {
		return true
	}
	return false
}
