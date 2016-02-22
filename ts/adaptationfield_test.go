package ts

import (
	"reflect"
	"testing"
)

var adaptationField = []byte{0x07, 0x10, 0x00, 0x17, 0x15, 0xc4, 0x7e, 0x00}

func TestNewAdaptationField(t *testing.T) {
	tests := []struct {
		name string
		data []byte
		want *AdaptationField
	}{
		{
			name: "Adaptation field",
			data: adaptationField,
			want: &AdaptationField{
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
		},
	}
	for _, tt := range tests {
		if got := newAdaptationField(tt.data); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. newAdaptationField() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
