package pes

type Header struct {
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

	PTS uint64
	DTS uint64
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

func newHeader(data []byte) *Header {
	h := &Header{
		ScramblingControl:          data[6] & 0x30 >> 4,
		Priority:                   data[6]&0x08>>3 == 1,
		DataAlignmentIndicator:     data[6]&0x04>>2 == 1,
		Copyright:                  data[6]&0x02>>1 == 1,
		Original:                   data[6]&0x01 == 1,
		ContainsPTS:                data[7]&0x80>>7 == 1,
		ContainsDTS:                data[7]&0x40>>6 == 1,
		ContainsESCR:               data[7]&0x10>>4 == 1,
		ContainsESRate:             data[7]&0x10>>4 == 1,
		ContainsDSMTrickMode:       data[7]&0x08>>3 == 1,
		ContainsAdditionalCopyInfo: data[7]&0x04>>2 == 1,
		ContainsCRC:                data[7]&0x02>>1 == 1,
		ContainsExtension:          data[7]&0x01 == 1,
		HeaderLength:               data[8]}

	if h.ContainsPTS && !h.ContainsDTS {
		h.PTS = uint64(data[9]&0x0E>>1) << 30
		h.PTS = h.PTS | (uint64(data[10])<<8|uint64(data[11]))>>1<<15
		h.PTS = h.PTS | (uint64(data[12])<<8|uint64(data[13]))>>1
	}

	if h.ContainsPTS && h.ContainsDTS {
		h.PTS = uint64(data[9]&0x0E>>1) << 30
		h.PTS = h.PTS | (uint64(data[10])<<8|uint64(data[11]))>>1<<15
		h.PTS = h.PTS | (uint64(data[12])<<8|uint64(data[13]))>>1

		h.DTS = uint64(data[14]&0x0E>>1) << 30
		h.DTS = h.DTS | (uint64(data[15])<<8|uint64(data[16]))>>1<<15
		h.DTS = h.DTS | (uint64(data[17])<<8|uint64(data[18]))>>1
	}

	return h
}
