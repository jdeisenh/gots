package pes

import (
	"fmt"
	"time"
)

func StreamId(value uint8) string {
	switch {
	case value == 0xBC:
		return "program_stream_map"
	case value == 0xBD:
		return "private_stream_1"
	case value == 0xBE:
		return "padding_stream"
	case value == 0xBF:
		return "private_stream_2"
	case 0xC0 <= value && value <= 0xDF:
		return "ISO/IEC 13818-3 or ISO/IEC 11172-3 or ISO/IEC 13818-7 or ISO/IEC 14496-3 audio stream number x xxxx"
	case 0xE0 <= value && value <= 0xEF:
		return "ITU-T Rec. H.262 | ISO/IEC 13818-2, ISO/IEC 11172-2, ISO/IEC 14496-2 or ITU-T Rec. H.264 | ISO/IEC 14496-10 video stream number xxxx"
	case value == 0xF0:
		return "ECM_stream"
	case value == 0xF1:
		return "EMM_stream"
	case value == 0xF2:
		return "ITU-T Rec. H.222.0 | ISO/IEC 13818-1 Annex A or ISO/IEC 13818-6_DSMCC_stream"
	case value == 0xF3:
		return "ISO/IEC_13522_stream"
	case value == 0xF4:
		return "ITU-T Rec. H.222.1 type A"
	case value == 0xF5:
		return "ITU-T Rec. H.222.1 type B"
	case value == 0xF6:
		return "ITU-T Rec. H.222.1 type C"
	case value == 0xF7:
		return "ITU-T Rec. H.222.1 type D"
	case value == 0xF8:
		return "ITU-T Rec. H.222.1 type E"
	case value == 0xF9:
		return "ancillary_stream"
	case value == 0xFA:
		return "ISO/IEC 14496-1_SL-packetized_stream"
	case value == 0xFB:
		return "ISO/IEC 14496-1_FlexMux_stream"
	case value == 0xFC:
		return "metadata stream"
	case value == 0xFD:
		return "extended_stream_id"
	case value == 0xFE:
		return "reserved data stream"
	case value == 0xFF:
		return "program_stream_directory"
	}
	return "N/A"

}

func toTime(h uint64) string {
	t := time.Duration((h / 90)) * time.Millisecond
	return fmt.Sprintf("%s", t)
}

func (p Packet) String() string {
	return fmt.Sprintf("\tCode Prefix: %x\n"+
		"\tStream ID: %x [%s]\n"+
		"\tPacket Length: %d\n"+
		"%s",
		p.CodePrefix,
		p.StreamID,
		StreamId(p.StreamID),
		p.PacketLength,
		p.Header)
}

func (h Header) String() string {
	return fmt.Sprintf("\tScramblingControl: %x\n"+
		"\tPriority: %t\n"+
		"\tDataAlignmentIndicator: %t\n"+
		"\tCopyright: %t\n"+
		"\tOriginal: %t\n"+
		"\tContains PTS: %t\n"+
		"\tContains DTS: %t\n"+
		"\tContains ESCR: %t\n"+
		"\tContains ESRate: %t\n"+
		"\tContains DSMTrickMode: %t\n"+
		"\tContains AdditionalCopyInfo: %t\n"+
		"\tContains CRC: %t\n"+
		"\tContains Extension: %t\n"+
		"\tHeaderLength: %d\n"+
		"\tPTS: %d [%s]\n"+
		"\tDTS: %d [%s]",
		h.ScramblingControl,
		h.Priority,
		h.DataAlignmentIndicator,
		h.Copyright,
		h.Original,
		h.ContainsPTS,
		h.ContainsDTS,
		h.ContainsESCR,
		h.ContainsESRate,
		h.ContainsDSMTrickMode,
		h.ContainsAdditionalCopyInfo,
		h.ContainsCRC,
		h.ContainsExtension,
		h.HeaderLength,
		h.PTS,
		toTime(h.PTS),
		h.DTS,
		toTime(h.DTS))
}
