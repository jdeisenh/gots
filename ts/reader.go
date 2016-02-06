package ts

import (
	"fmt"
	"io"
)

type Reader struct {
	reader    io.Reader
	remainder []byte
	off       int

	Packet      *Packet
	PAT         *ProgramAssociationTable
	PMT         *ProgramMapTable
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

func (r *Reader) Read(p []byte) (n int, err error) {
	var ac, rd int
	for ac < cap(p) {
		if r.Packet == nil || r.off == len(r.Packet.Payload) {
			_, err := r.Next()
			r.off = 0
			if err != nil {
				return 0, err
			}
		}
		if ac+r.off < len(r.Packet.Payload) {
			rd = copy(p[ac:cap(p)], r.Packet.Payload[r.off:])
		} else {
			rd = copy(p[ac:len(r.Packet.Payload)-ac], r.Packet.Payload[r.off:])
		}
		r.off = r.off + rd
	}
	return ac, nil
}

func (r *Reader) Next() (*Packet, error) {
	if err := r.syncToPacket(); err != nil {
		return nil, err
	}
	if err := r.readUntilPacket(); err != nil {
		return nil, err
	}

	var err error
	r.Packet, err = NewPacket(r.remainder[:PacketSize])
	if err == nil {
		r.OnNewPacket(r.Packet)
		if r.Packet.hasProgramAssociationTable() {
			r.PAT = newPogramAssociationTable(r.Packet.Payload)
			r.OnNewPAT(r.PAT)
		}
		if r.Packet.hasProgramMapTable(r.PAT) {
			r.PMT = newProgramMapTable(r.Packet.Payload)
			r.OnNewPMT(r.PMT)
		}
	}
	r.remainder = r.remainder[PacketSize:]
	return r.Packet, err
}

func (r *Reader) syncToPacket() error {
	if len(r.remainder) == 0 || r.remainder[0] != SyncByte {
		buffer := make([]byte, 1)
		for {
			b, err := r.reader.Read(buffer)
			if err != nil {
				return err
			}
			if b == 0 {
				return fmt.Errorf("Failed to find sync byte")
			}
			if buffer[0] == SyncByte {
				r.remainder = buffer
				return nil
			}
		}
	}
	return nil
}

func (r *Reader) readUntilPacket() error {
	buffer := make([]byte, 64)
	for len(r.remainder) < PacketSize {
		b, err := r.reader.Read(buffer)
		if err != nil {
			return err
		}
		if b == 0 {
			return fmt.Errorf("Failed to read full packet got %d out of %d bytes", len(r.remainder), PacketSize)
		}
		r.remainder = append(r.remainder, buffer...)
	}
	return nil
}
