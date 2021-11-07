package iterator

import (
	"fmt"
	"testing"
)

func Test01(t *testing.T) {
	s1 := &Station{name: "s1", frequency: 96.8}
	s2 := &Station{name: "s2", frequency: 101.5}

	radioStation := &RadioStation{writeIndex: 0, stations: make([]*Station, 10)}
	radioStation.addStation(s1)
	radioStation.addStation(s2)

	for radioStation.hasNext() {
		r := radioStation.getNext()
		fmt.Printf("Station: %s, Frequency: %5.1f \n", r.name, r.frequency)
	}
}
