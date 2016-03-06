GoTS ![License](https://img.shields.io/badge/license-Apache%202.0-blue.svg) [![Build Status](https://travis-ci.org/damienlevin/gots.svg?branch=master)](https://travis-ci.org/damienlevin/gots) [![codebeat badge](https://codebeat.co/badges/c63c07d0-e4ff-45b9-9283-9861f6d9c720)](https://codebeat.co/projects/github-com-damienlevin-gots)
======
[MPEG](https://en.wikipedia.org/wiki/MPEG_transport_stream) transport stream parser written in [Go](golang.org).

- [x] TS packet parser
    - [x] Program Association Table
    - [x] Program Map Table
    - [ ] Conditional Access Table
    - [ ] Network Information Table
    - [ ] Transport Stream Description Table
    - [ ] IPMP Control Information Table  
    
- [x] PES packet parser
    - [x] PTS
    - [x] DTS
    - [ ] Full support


## Installation
    go get -u github.com/damienlevin/gots

## Usage

``` 
func main() {
	f, _ := os.Open(os.Args[1])
	// Create new TS Reader and provide
	// callbacks for TS,PAT and PMT
	t := ts.NewReader(f, callbackTS, callbackPAT, callbackPMT)
	// Create new PES Reaser and provide callback for PES
	p := pes.NewReader(t, calbackPES)

	// Iterate through the PES packets
	for {
		_, err := p.Next()
		if err != nil {
			return
		}
	}
}

``` 


    
Example usage from `parser.go` :

    ./parser  -u "https://devimages.apple.com.edgekey.net/streaming/examples/bipbop_4x3/gear1/fileSequence179.ts" | less
    
    
``` 
...
============================================================
TS packet [4]
============================================================
SyncByte: 47
Transport Error Indicator: false
Payload Unit Start Indicator: true
Transport Priority: false
PID: 257
Transport Scrambling Control: 0
Contains Adaptation Field: true
Contains Payload: true
Continuity Counter: 7
        AdaptationField:
        Adaptation Field Length: 7
        Discontinuity Indicator: false
        RandomAccess Indicator: false
        Elementary StreamPriority Indicator: false
        Contains PCR: true
        Contains OPCR: false
        Contains Splicing Point: false
        Contains Transport Private Data: false
        Contains Adaptation Field Extension: false
        PCR: 04d113fa7eed
        OPCR: 
        Splice Countdown: 0
        Transport Private Data Lenght: 0
------------------------------------------------------------
Payload (176) 
------------------------------------------------------------
00 00 01 e0 00 00 84 c0 0a 31 26 89 75 61 11 26 
89 5d f1 00 00 00 01 09 f0 00 00 01 06 05 11 03 
87 f4 4e cd 0a 4b dc a1 94 3a c3 d4 9b 17 1f 00 
80 00 00 01 21 e1 09 08 37 ff ff ff ff ff ff ff 
ff ff d4 bf e2 bf fa ff ff ef ff d7 da 31 10 b1 
50 26 fd 64 60 e0 74 d2 e5 21 ad 31 57 89 4f 0e 
a8 37 b3 4c 53 5c bc a5 eb 8d bb ff cd fe 30 0f 
85 60 21 07 49 8d a9 27 87 e1 37 a4 99 fb e6 ff 
4d 71 58 57 e1 dc d1 83 fa 9a 61 9b ea bf ff ff 
ff ff fc 9e 6f ff cc 01 d8 89 8a 72 1d 8a 06 6a 
7f f5 a3 31 84 40 6c bc c4 c0 47 d7 20 41 80 89 
------------------------------------------------------------
PES packet [2]
------------------------------------------------------------
        Code Prefix: 1
        Stream ID: e0 [ITU-T Rec. H.262 | ISO/IEC 13818-2, ISO/IEC 11172-2, ISO/IEC 14496-2 or ITU-T Rec. H.264 | ISO/IEC 14496-10 video stream number xxxx]
        Packet Length: 0
        ScramblingControl: 0
        Priority: false
        DataAlignmentIndicator: true
        Copyright: false
        Original: false
        Contains PTS: true
        Contains DTS: true
        Contains ESCR: false
        Contains ESRate: false
        Contains DSMTrickMode: false
        Contains AdditionalCopyInfo: false
        Contains CRC: false
        Contains Extension: false
        HeaderLength: 10
        PTS: 161626800 [29m55.853s]
        DTS: 161623800 [29m55.82s]
...
```
## License

This project is released under the  [Apache License v2](http://www.apache.org/licenses/LICENSE-2.0)
