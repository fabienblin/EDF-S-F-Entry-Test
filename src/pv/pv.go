package pv

import (
	. "libs/emslib"
	"libs/environment"
	"fmt"
	"poc"
)

type Pv struct {
	Ppv   Watt
	Pprod WattPerSqrMeter
	SetpointPPv Watt
}

// constants can be changed to create production scenarios
const nbSolarPanels int = 100
const solarPanelSurface int = 6
const solarTransformationEfficiency float64 = 0.2

var pvPeak KWatt = PvPeak()

func PvPeak() KWatt {
	pvTotalSurface := nbSolarPanels * solarPanelSurface
	peak := Watt(float64(pvTotalSurface) * solarTransformationEfficiency)
	return WattToKWatt(peak)
}

// defines solar park total power output in watts
// setpoint >= Ppv >= 0 is always true
func (pv *Pv) simulatePpv() {
	simPpv := Watt(environment.GetSunShine() * solarTransformationEfficiency * float64(nbSolarPanels))

	if simPpv > pv.SetpointPPv {
		simPpv = pv.SetpointPPv
	} else if simPpv < 0 { // this can't happen
		fmt.Println("FAILURE : simulatePpv < 0")
		simPpv = 0
	}

	pv.Ppv = simPpv
}

// simulates pyranometer's power output estimate in watts/m² (arbitrary value)
// IRL this is a measure, not a calculation
func (pv *Pv) simulatePprod() {
	pv.Pprod = WattPerSqrMeter(environment.GetSunShine() * 10)
}

// define PV variables depending on environment conditions
// return power production
func (pv *Pv) SimulatePv(poc *poc.Poc) {
	pv.simulatePpv()
	pv.simulatePprod()
	
	poc.SimulatePoc(WattToKWatt(pv.Ppv))
}

func (pv *Pv) Show() {
	fmt.Println(" PV { Ppv :", pv.Ppv, "W ; Pprod :", pv.Pprod, "W/m² ; SetPoint :", pv.SetpointPPv, "W }")
}