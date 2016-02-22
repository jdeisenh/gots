GoTS [![Build Status](https://travis-ci.org/damienlevin/gots.svg?branch=master)](https://travis-ci.org/damienlevin/gots)
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
    cd $GOPATH
    go get -u github.com/dml/gots

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

    ./parser "https://devimages.apple.com.edgekey.net/streaming/examples/bipbop_4x3/gear1/fileSequence179.ts" | less
    
    
``` 
...
============================================================
TS packet [87]
============================================================
SyncByte: 47
Transport Error Indicator: false
Payload Unit Start Indicator: true
Transport Priority: false
PID: 4096
Transport Scrambling Control: 0
Contains Adaptation Field: false
Contains Payload: true
Continuity Counter: 2
------------------------------------------------------------
Payload (184) 
------------------------------------------------------------
00 02 b0 17 00 01 c1 00 00 e1 00 f0 00 1b e1 00 
f0 00 03 e1 01 f0 00 4e 59 3d 1e ff ff ff ff ff 
ff ff ff ff ff ff ff ff ff ff ff ff ff ff ff ff 
ff ff ff ff ff ff ff ff ff ff ff ff ff ff ff ff 
ff ff ff ff ff ff ff ff ff ff ff ff ff ff ff ff 
ff ff ff ff ff ff ff ff ff ff ff ff ff ff ff ff 
ff ff ff ff ff ff ff ff ff ff ff ff ff ff ff ff 
ff ff ff ff ff ff ff ff ff ff ff ff ff ff ff ff 
ff ff ff ff ff ff ff ff ff ff ff ff ff ff ff ff 
ff ff ff ff ff ff ff ff ff ff ff ff ff ff ff ff 
ff ff ff ff ff ff ff ff ff ff ff ff ff ff ff ff 
ff ff ff ff ff ff ff ff 
------------------------------------------------------------
PMT 
------------------------------------------------------------
        Table Id: 2
        Section Syntax Indicator: true
        Section Lenght: 23
        Program Number: 1
        Version Number: 0
        Current Next Indicator: true
        Section Number: 0
        Last Section Number: 0
        PCR PID: 256
        Program Info Length: 0
        Stream Descriptors: []
                Streams:
                Stream Type: 1b [AVC video stream as defined in ITU-T Rec. H.264 | ISO/IEC 14496-10 Video]
                Elementaty PID: 256
                ES Info Length: 0
                Data: 03e101f000
                Stream Type: 3 [ISO/IEC 11172-3 Audio]
                Elementaty PID: 257
                ES Info Length: 0
                Data: 
        CRC32: 4e593d1e

...
```
## License

This project is released under the  [Apache License v2](http://www.apache.org/licenses/LICENSE-2.0)
