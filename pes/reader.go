package pes

import "github.com/damienlevin/gots/ts"

type Reader struct {
	reader    *ts.Reader
	remainder []byte

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
	ac := -6
	for p == nil || ac < int(p.PacketLength) {
		ts, err := r.reader.Next()
		if err != nil {
			return nil, err
		}
		if pmt := r.reader.PMT; pmt != nil && pmt.HasElementaryStream(ts.PID) {
			if ts.PayloadUnitStartIndicator {
				p, err = NewPacket(ts.Payload)
				if err != nil {
					return nil, err
				}
			}
			ac = ac + len(ts.Payload)
		}
	}
	r.OnNewPacket(p)
	return p, nil
}
