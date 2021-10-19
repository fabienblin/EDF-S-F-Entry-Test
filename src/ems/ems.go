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
func (ems *Ems) setpointPEss(setpointPEss KWatt) {
	ems.Ess.SetpointPEss = setpointPEss
}

// define PV power output to smart grid
func (ems *Ems) setpointPPv(setpointPPv Watt) {
	ems.Pv.SetpointPPv = setpointPPv
}

// returns current state of ESS
func (ems *Ems) getEssMeasure() *ess.Ess {
	return ems.Ess
}

// returns current state of PV
func (ems *Ems) getPvMeasure() *pv.Pv {
	return ems.Pv
}

// returns current state of POC
func (ems *Ems) getPocMeterMeasure() KWatt {
	return ems.Poc.SimulatePoc(0)
}

// log the EMS current state
func (ems *Ems) Show() {
	fmt.Println("EMS {")
	ems.Pv.Show()
	ems.Ess.Show()
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
func (ems *Ems) getPLoad() KWatt {
	return ems.Poc.Ppoc - ems.Ess.Pess - WattToKWatt(ems.Pv.Ppv)
}

// pourcent of Ppoc for ESS
var essShareRatio float64

var initial bool = true
// use only given functions from exercise
func (ems *Ems) Ai() {
	// collect data
	ems.getEssMeasure()
	ems.getPocMeterMeasure()
	ems.getPvMeasure()

	var pvSet KWatt = WattToKWatt(ems.Pv.SetpointPPv)
	var essSet KWatt = ems.Ess.SetpointPEss
	if initial {
		pvSet = 1000
		essSet = 0
		initial = false
	}

	if ems.Poc.Ppoc > 0 { // PV overproductive
		if ems.Ess.IsFull { // ESS fully charged
			fmt.Println("STOP PV")
			pvSet = -ems.getPLoad()
			essSet = 0
		} else { // store excess PV prod
			fmt.Println("STORE EXCESS")
			pvSet = -ems.getPLoad() - ems.Ess.Pmaxch
			essSet = -(WattToKWatt(ems.Pv.Ppv) + ems.getPLoad())
		}
	} else { // machine learning descision
		fmt.Println("MACHINE LEARNING")
	}

	fmt.Println("ems.setpointPPv(KWattToWatt(", pvSet, "))")
	fmt.Println("ems.setpointPEss(", essSet, ")")

	ems.setpointPEss(essSet)
	ems.setpointPPv(KWattToWatt(pvSet))

}