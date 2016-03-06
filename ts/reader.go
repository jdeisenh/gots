package ts

import (
	"io"
	"math"
)

type Reader struct {
	reader io.Reader
	offset int

	Packet *Packet
	PAT    *ProgramAssociationTable
	PMT    *ProgramMapTable

	OnNewPacket func(*Packet)
	OnNewPAT    func(*ProgramAssociationTable)
	OnNewPMT    func(*ProgramMapTable)
}

func NewReader(r io.Reader, c func(*Packet), a func(*ProgramAssociationTable), m func(*ProgramMapTable)) *Reader {
	return &Reader{
		reader:      r,
		OnNewPacket: c,
		OnNewPAT:    a,
		OnNewPMT:    m}
}

func (r *Reader) Next() (*Packet, error) {
	data, err := r.readBytes(PacketSize)
	if err != nil {
		return nil, err
	}
	r.Packet, err = newPacket(data)
	if err != nil {
		return nil, err
	}
	r.OnNewPacket(r.Packet)
	if r.Packet.hasProgramAssociationTable() {
		r.PAT = newPogramAssociationTable(r.Packet.Payload)
		r.OnNewPAT(r.PAT)
	}
	if r.Packet.hasProgramMapTable(r.PAT) {
		r.PMT = newProgramMapTable(r.Packet.Payload)
		r.OnNewPMT(r.PMT)
	}
	return r.Packet, nil
}

func (r *Reader) Read(b []byte) (int, error) {
	read := 0
	for read < cap(b) {
		if r.Packet == nil || r.offset == len(r.Packet.Payload) {
			if _, err := r.Next(); err != nil {
				return read, err
			}
			r.offset = 0
		}
		rd := copy(b[read:], r.Packet.Payload[r.offset:int(math.Min(float64(r.offset+cap(b)-read), float64(len(r.Packet.Payload))))])
		r.offset += rd
		read += rd
	}
	return read, nil
}

func (r *Reader) readBytes(length int) ([]byte, error) {
	buffer := make([]byte, length)
	read := 0
	for read < length {
		b, err := r.reader.Read(buffer[read:])
		if err != nil {
			return nil, err
		}
		read += b
	}
	return buffer, nil
}
