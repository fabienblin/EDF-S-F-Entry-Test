package main

import (
	"fmt"
	. "emslib"
)

// EMS definition

type emsEnv struct {
	ess  *Ess
	pv   *Pv
	poc  *Poc
	PEss Watt
	PPv  Watt
}

func (ems *emsEnv) setpointPEss(setpointPEss Watt) {
	ems.PEss = setpointPEss
}

func (ems *emsEnv) setpointPPv(setpointPPv Watt) {
	ems.PPv = setpointPPv
}

func (ems *emsEnv) getEssMeasure() *Ess {
	return ems.ess
}

func (ems *emsEnv) getPvMeasure() *Pv {
	return ems.pv
}

func (ems *emsEnv) getPocMeterMeasure() *Poc {
	return ems.poc
}

var _ = Void
func main() {
	var w Watt = 0

	fmt.Println(w)
}
