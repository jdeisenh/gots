package ts

import (
	"reflect"
	"testing"
)

var packetWithAdaptationField = []byte{
	0x47, 0x41, 0x00, 0x30, 0x07, 0x10, 0x00, 0x17, 0x15, 0xc4, 0x7e, 0x00, 0x00, 0x00, 0x01, 0xe0,
	0xe4, 0xfc, 0x80, 0xc0, 0x0a, 0x31, 0x00, 0xbd, 0x7b, 0x81, 0x11, 0x00, 0xbd, 0x43, 0x41, 0x00,
	0x00, 0x00, 0x01, 0x09, 0xf0, 0x00, 0x00, 0x00, 0x01, 0x67, 0x4d, 0x40, 0x1f, 0xec, 0xa0, 0x50,
	0x17, 0xfc, 0xb0, 0x80, 0x00, 0x00, 0x03, 0x00, 0x80, 0x00, 0x00, 0x19, 0x47, 0x8c, 0x18, 0xcb,
	0x00, 0x00, 0x00, 0x01, 0x68, 0xef, 0xbc, 0x80, 0x00, 0x00, 0x01, 0x65, 0x88, 0x84, 0x00, 0xf6,
	0xff, 0xe1, 0x17, 0x86, 0x54, 0x23, 0x96, 0x1f, 0xf4, 0x21, 0xe5, 0x8f, 0xf1, 0x08, 0xd7, 0x7d,
	0x2a, 0xc1, 0xb9, 0x2d, 0x9f, 0x9c, 0xfe, 0xfa, 0x9e, 0xc4, 0x55, 0x71, 0xa8, 0xb8, 0x2b, 0xaa,
	0x82, 0x9f, 0xb5, 0x03, 0xb1, 0x6c, 0x66, 0xfb, 0x56, 0xa0, 0x93, 0xf8, 0xcf, 0x61, 0xde, 0x56,
	0xf0, 0xdc, 0xa2, 0x1a, 0xd6, 0xcc, 0x4b, 0x04, 0x9c, 0xf0, 0x4c, 0x38, 0x29, 0x98, 0x1b, 0x95,
	0x02, 0x61, 0x49, 0x66, 0xf2, 0xb4, 0x07, 0xe6, 0x42, 0x33, 0x76, 0x70, 0xa7, 0x0e, 0xc1, 0x33,
	0x94, 0xe8, 0xe4, 0x37, 0xe3, 0x25, 0x6c, 0xc0, 0x5d, 0xd2, 0x66, 0xd3, 0x4d, 0x30, 0xe1, 0x7a,
	0x67, 0x80, 0x08, 0xb8, 0x26, 0x84, 0x3f, 0xcf, 0x60, 0x22, 0x16, 0x75}

func TestNewPacket(t *testing.T) {
	tests := []struct {
		name    string
		data    []byte
		want    *Packet
		wantErr bool
	}{
		{
			name:    "Empty []byte",
			data:    []byte{},
			wantErr: true,
		}, {
			name:    "Not starting with sync byte",
			data:    make([]byte, 188),
			wantErr: true,
		}, {
			name:    "Packet with adaptation field",
			data:    packetWithAdaptationField,
			wantErr: false,
			want: &Packet{
				SyncByte:                  0x47,
				TransportErrorIndicator:   false,
				PayloadUnitStartIndicator: true,
				TransportPriority:         false,
				PID:                       256,
				TransportScramblingControl: 0,
				ContainsAdaptationField:    true,
				ContainsPayload:            true,
				ContinuityCounter:          0,
				AdaptationField: &AdaptationField{
					AdaptationFieldLength:             7,
					DiscontinuityIndicator:            false,
					RandomAccessIndicator:             false,
					ElementaryStreamPriorityIndicator: false,
					ContainsPCR:                       true,
					ContainsOPCR:                      false,
					ContainsSplicingPoint:             false,
					ContainsTransportPrivateData:      false,
					ContainsAdaptationFieldExtension:  false,
					PCR:                        []byte{0x00, 0x17, 0x15, 0xc4, 0x7e, 0x00},
					OPCR:                       nil,
					SpliceCountdown:            0,
					TransportPrivateDataLenght: 0,
					PrivateData:                nil,
				},
				Payload: packetWithAdaptationField[12:],
			},
		},
	}
	for _, tt := range tests {
		got, err := newPacket(tt.data)
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. NewPacket() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. NewPacket() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestPacketHasProgramAssociationTable(t *testing.T) {
	tests := []struct {
		name string
		pid  uint16
		want bool
	}{
		{
			name: "Contains association table",
			pid:  ProgramAssociationTableID,
			want: true,
		}, {
			name: "Doesn't contain association table",
			pid:  0x1234,
			want: false,
		},
	}
	for _, tt := range tests {
		p := Packet{PID: tt.pid}
		if got := p.hasProgramAssociationTable(); got != tt.want {
			t.Errorf("%q. Packet.hasProgramAssociationTable() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestPacketHasProgramMapTable(t *testing.T) {
	tests := []struct {
		name string
		pid  uint16
		pat  *ProgramAssociationTable
		want bool
	}{
		{
			name: "Contains program map table",
			pid:  0x1234,
			pat: &ProgramAssociationTable{
				Programs: []*Program{
					&Program{ProgramMapPID: 0x1234},
				},
			},
			want: true,
		}, {
			name: "Doesn't contain program map table",
			pid:  0x4321,
			pat: &ProgramAssociationTable{
				Programs: []*Program{
					&Program{ProgramMapPID: 0x1234},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		p := Packet{PID: tt.pid}
		if got := p.hasProgramMapTable(tt.pat); got != tt.want {
			t.Errorf("%q. Packet.hasProgramMapTable() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
