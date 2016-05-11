package pes

import "github.com/jdeisenh/gots/ts"

type Reader struct {
	reader *ts.Reader

	Packet      *Packet
	OnNewPacket func(*Packet)
}

func NewReader(r *ts.Reader, c func(*Packet)) *Reader {
	return &Reader{
		reader:      r,
		OnNewPacket: c}
}

func (r *Reader) Next() (*Packet, error) {
	var p *Packet
	var err error
	for p == nil {
		if _, err := r.reader.Next(); err != nil {
			return nil, err
		}
		t := r.reader.Packet
		pmt := r.reader.PMT
		if t != nil && pmt != nil && pmt.HasElementaryStream(t.PID) {
			if t.PayloadUnitStartIndicator {
				p, err = newPacket(t.Payload)
				if err != nil {
					return nil, err
				}
				r.OnNewPacket(p)
				return p, nil
			}
		}
	}
	return nil, nil
}
