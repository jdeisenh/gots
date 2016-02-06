package ts

import "fmt"

func StreamType(value uint8) string {
	switch {
	case value == 0x00:
		return "ITU-T | ISO/IEC Reserved"
	case value == 0x01:
		return "ISO/IEC 11172-2 Video"
	case value == 0x02:
		return "ITU-T Rec. H.262 | ISO/IEC 13818-2 Video or ISO/IEC 11172-2 constrained parameter video stream"
	case value == 0x03:
		return "ISO/IEC 11172-3 Audio"
	case value == 0x04:
		return "ISO/IEC 13818-3 Audio"
	case value == 0X05:
		return "ITU-T Rec. H.222.0 | ISO/IEC 13818-1 private_sections"
	case value == 0x06:
		return "ITU-T Rec. H.222.0 | ISO/IEC 13818-1 PES packets containing private data"
	case value == 0x07:
		return "ISO/IEC 13522 MHEG"
	case value == 0x08:
		return "ITU-T Rec. H.222.0 | ISO/IEC 13818-1 Annex A DSM-CC"
	case value == 0x09:
		return "ITU-T Rec. H.222.1"
	case value == 0x0A:
		return "ISO/IEC 13818-6 type A"
	case value == 0x0B:
		return "ISO/IEC 13818-6 type B"
	case value == 0x0C:
		return "ISO/IEC 13818-6 type C"
	case value == 0x0D:
		return "ISO/IEC 13818-6 type D"
	case value == 0x0E:
		return "ITU-T Rec. H.222.0 | ISO/IEC 13818-1 auxiliary"
	case value == 0x0F:
		return "ISO/IEC 13818-7 Audio with ADTS transport syntax"
	case value == 0x10:
		return "ISO/IEC 14496-2 Visual"
	case value == 0x11:
		return "ISO/IEC 14496-3 Audio with the LATM transport syntax as defined in ISO/IEC 14496-3"
	case value == 0x12:
		return "ISO/IEC 14496-1 SL-packetized stream or FlexMux stream carried in PES packets"
	case value == 0X13:
		return "ISO/IEC 14496-1 SL-packetized stream or FlexMux stream carried in ISO/IEC 14496_sections"
	case value == 0x14:
		return "ISO/IEC 13818-6 Synchronized Download Protocol"
	case value == 0x15:
		return "Metadata carried in PES packets"
	case value == 0x16:
		return "Metadata carried in metadata_sections"
	case value == 0x17:
		return "Metadata carried in ISO/IEC 13818-6 Data Carousel"
	case value == 0x18:
		return "Metadata carried in ISO/IEC 13818-6 Object Carousel"
	case value == 0x19:
		return "Metadata carried in ISO/IEC 13818-6 Synchronized Download Protocol"
	case value == 0x1A:
		return "IPMP stream (defined in ISO/IEC 13818-11, MPEG-2 IPMP)"
	case value == 0x1B:
		return "AVC video stream as defined in ITU-T Rec. H.264 | ISO/IEC 14496-10 Video"
	case 0x1C <= value && value <= 0x7E:
		return "ITU-T Rec. H.222.0 | ISO/IEC 13818-1 Reserved"
	case value == 0x7F:
		return "IPMP stream"
	}
	return "User Private"

}

func (p Packet) String() string {
	var adf string
	if p.ContainsAdaptationField {
		adf = fmt.Sprintf("\n\tAdaptationField:\n%s", p.AdaptationField)
	}
	return fmt.Sprintf("SyncByte: %x\n"+
		"Transport Error Indicator: %t\n"+
		"Payload Unit Start Indicator: %t\n"+
		"Transport Priority: %t\n"+
		"PID: %d\n"+
		"Transport Scrambling Control: %x\n"+
		"Contains Adaptation Field: %t\n"+
		"Contains Payload: %t\n"+
		"Continuity Counter: %d"+
		adf,
		p.SyncByte,
		p.TransportErrorIndicator,
		p.PayloadUnitStartIndicator,
		p.TransportPriority,
		p.PID,
		p.TransportScramblingControl,
		p.ContainsAdaptationField,
		p.ContainsPayload,
		p.ContinuityCounter)
}

func (a AdaptationField) String() string {
	var pv string
	if a.ContainsTransportPrivateData {
		pv = fmt.Sprintf("Private Data:\n%s\n", a.PrivateData)
	}
	return fmt.Sprintf("\tAdaptation Field Length: %d\n"+
		"\tDiscontinuity Indicator: %t\n"+
		"\tRandomAccess Indicator: %t\n"+
		"\tElementary StreamPriority Indicator: %t\n"+
		"\tContains PCR: %t\n"+
		"\tContains OPCR: %t\n"+
		"\tContains Splicing Point: %t\n"+
		"\tContains Transport Private Data: %t\n"+
		"\tContains Adaptation Field Extension: %t\n"+
		"\tPCR: %x\n"+
		"\tOPCR: %x\n"+
		"\tSplice Countdown: %d\n"+
		"\tTransport Private Data Lenght: %d"+
		pv,
		a.AdaptationFieldLength,
		a.DiscontinuityIndicator,
		a.RandomAccessIndicator,
		a.ElementaryStreamPriorityIndicator,
		a.ContainsPCR,
		a.ContainsOPCR,
		a.ContainsSplicingPoint,
		a.ContainsTransportPrivateData,
		a.ContainsAdaptationFieldExtension,
		a.PCR,
		a.OPCR,
		a.SpliceCountdown,
		a.TransportPrivateDataLenght)
}

func (a ProgramAssociationTable) String() string {
	prg := "\t\tPrograms:\n"
	for _, p := range a.Programs {
		prg = prg + fmt.Sprintf("%s", p)
	}

	return fmt.Sprintf("\tTableId: %x\n"+
		"\tSection Syntax Indicator: %t\n"+
		"\tSection Lenght: %d\n"+
		"\tTransport Stream Id: %d\n"+
		"\tVersion Number: %d\n"+
		"\tCurrent Next Indicator: %t\n"+
		"\tSection Number: %d\n"+
		"\tLast Section Number: %d\n"+
		prg+
		"\tCRC32: %x",
		a.TableId,
		a.SectionSyntaxIndicator,
		a.SectionLenght,
		a.TransportStreamId,
		a.VersionNumber,
		a.CurrentNextIndicator,
		a.SectionNumber,
		a.LastSectionNumber,
		a.CRC32)
}

func (p Program) String() string {
	return fmt.Sprintf("\t\tNumber: %d\n"+
		"\t\tNetwork PID: %d\n"+
		"\t\tProgram Map PID: %d\n",
		p.Number,
		p.NetworkPID,
		p.ProgramMapPID)
}

func (a ProgramMapTable) String() string {
	streams := "\t\tStreams:\n"
	for _, s := range a.ElementaryStreams {
		streams = streams + fmt.Sprintf("%s", s)
	}
	return fmt.Sprintf("\tTable Id: %x\n"+
		"\tSection Syntax Indicator: %t\n"+
		"\tSection Lenght: %d\n"+
		"\tProgram Number: %d\n"+
		"\tVersion Number: %d\n"+
		"\tCurrent Next Indicator: %t\n"+
		"\tSection Number: %d\n"+
		"\tLast Section Number: %d\n"+
		"\tPCR PID: %d\n"+
		"\tProgram Info Length: %d\n"+
		"\tStream Descriptors: %v\n"+
		streams+
		"\tCRC32: %x\n",
		a.TableId,
		a.SectionSyntaxIndicator,
		a.SectionLenght,
		a.ProgramNumber,
		a.VersionNumber,
		a.CurrentNextIndicator,
		a.SectionNumber,
		a.LastSectionNumber,
		a.PCRPID,
		a.ProgramInfoLength,
		a.StreamDescriptors,
		a.CRC32)
}

func (d StreamDescriptor) String() string {
	return fmt.Sprintf("\t\tDescriptor Tag: %x\n"+
		"\t\tDescriptor Length: %x\n"+
		"\t\tDescriptor Data: %x\n",
		d.DescriptorTag,
		d.DescriptorLength,
		d.Data)
}

func (s ElementaryStream) String() string {
	return fmt.Sprintf("\t\tStream Type: %x [%s]\n"+
		"\t\tElementaty PID: %d\n"+
		"\t\tES Info Length: %d\n"+
		"\t\tData: %x\n",
		s.StreamType,
		StreamType(s.StreamType),
		s.ElementaryPID,
		s.ESInfoLength,
		s.Data)
}
