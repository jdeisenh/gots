package main

import (
	"fmt"
	"github.com/damienlevin/gots/pes"
	"github.com/damienlevin/gots/ts"
	"net/http"
	"os"
)

var TSIndex = 1

func main() {
	rsp, err := http.Get(os.Args[1])
	if err != nil {
		panic(err)
	}
	t := ts.NewReader(rsp.Body, displayTSPacket, displayPAT, displayPMT)
	p := pes.NewReader(t, displayPES)

	for {
		_, err := p.Next()
		if err != nil {
			return
		}
	}
}

func displayTSPacket(p *ts.Packet) {
	fmt.Println("============================================================")
	fmt.Printf("TS packet [%d]\n", TSIndex)
	fmt.Println("============================================================")
	fmt.Printf("%s\n", p)
	fmt.Println("------------------------------------------------------------")
	fmt.Printf("Payload (%d) \n", len(p.Payload))
	fmt.Println("------------------------------------------------------------")
	displayPayload(p.Payload)
	TSIndex++
}

func displayPAT(m *ts.ProgramAssociationTable) {
	fmt.Println("------------------------------------------------------------")
	fmt.Printf("PAT \n")
	fmt.Println("------------------------------------------------------------")
	fmt.Printf("%s\n", m)
}

func displayPMT(m *ts.ProgramMapTable) {
	fmt.Println("------------------------------------------------------------")
	fmt.Printf("PMT \n")
	fmt.Println("------------------------------------------------------------")
	fmt.Printf("%s\n", m)
}

func displayPES(m *pes.Packet) {
	fmt.Println("------------------------------------------------------------")
	fmt.Printf("PES \n")
	fmt.Println("------------------------------------------------------------")
	fmt.Printf("%s\n", m)
}

func displayPayload(bytes []byte) {
	for i, b := range bytes {
		if (i+1)%16 == 0 || i+1 == len(bytes) {
			fmt.Printf("%02x \n", b)
			continue
		}
		fmt.Printf("%02x ", b)
	}
}
