package ems

import (
	"fmt"
	. "libs/emslib"
	"ess"
	"pv"
	"poc"
)

// constants can be changed to simulate scenarios
var pMaxSite KWatt

type Ems struct {
	Ess  *ess.Ess
	Pv   *pv.Pv
	Poc  *poc.Poc
}

// define ESS power output to smart grid
func (ems *Ems) SetpointPEss(setpointPEss KWatt) {
	ems.Ess.SetpointPEss = setpointPEss
}

// define PV power output to smart grid
func (ems *Ems) SetpointPPv(setpointPPv Watt) {
	ems.Pv.SetpointPPv = setpointPPv
}

// returns current state of ESS
func (ems *Ems) GetEssMeasure() *ess.Ess {
	return ems.Ess
}

// returns current state of PV
func (ems *Ems) GetPvMeasure() *pv.Pv {
	return ems.Pv
}

// returns current state of POC
func (ems *Ems) GetPocMeterMeasure() KWatt {
	return ems.Poc.SimulatePoc(0)
}

// log the EMS current state
func (ems *Ems) Show() {
	fmt.Println("EMS {")
	ems.Ess.Show()
	ems.Pv.Show()
	ems.Poc.Show()
	fmt.Println("}")
}

// set the initial state of EMS
func (ems *Ems) IntializeEms(pMaxSite KWatt) *Ems {
	var ESS *ess.Ess = new(ess.Ess)
	var poc *poc.Poc = new(poc.Poc)
	var pv *pv.Pv = new(pv.Pv)

	ems.Ess = ESS
	ems.Poc = poc
	ems.Pv = pv

	ess.InitializeEss(ems.Ess)

	return ems
}

// returns facility power demand
// knowing Ppoc = Pess + Ppv + Pload
func (ems *Ems) GetPLoad() KWatt {
	return ems.Poc.Ppoc - ems.Ess.Pess - WattToKWatt(ems.Pv.Ppv)
}

// returns smart grid power supply/demand
// knowing Ppoc = Pess + Ppv + Pload
func (ems *Ems) GetPPoc() KWatt {
	return ems.Ess.Pess + WattToKWatt(ems.Pv.Ppv) + ems.GetPLoad()
}

// use only given functions from exercise
func (ems *Ems) Ai() {
	// if PV produces more than facility demand : charge ESS
	// if ESS is full : limit PV production
	ems.SetpointPPv(1000000)
}