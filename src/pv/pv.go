package pv

import (
	. "libs/emslib"
	"libs/environment"
	"fmt"
)

type Pv struct {
	Ppv   Watt
	Pprod WattPerSqrMeter
}

// constants can be changed to create production scenarios
const nbSolarPanels int = 100
const solarPanelSurface int = 6
const solarTransformationEfficiency float64 = 0.2

var pvPeak KWatt = PvPeak()

func PvPeak() KWatt{
	pvTotalSurface := nbSolarPanels * solarPanelSurface
	peak := Watt(float64(pvTotalSurface) * solarTransformationEfficiency)
	return WattToKWatt(peak)
}

// defines solar park total power output in watts
func simulatePpv() Watt{
	return Watt(environment.GetSunShine() * solarTransformationEfficiency * float64(nbSolarPanels))
}

// simulates pyranometer's power output estimate in watts/m² (arbitrary value)
// IRL this is a measure, not a calculation
func simulatePprod() WattPerSqrMeter{
	return WattPerSqrMeter(environment.GetSunShine() * 10)
}

// define PV variables depending on environment conditions
func (pv *Pv) SimulatePv() {
	pv.Ppv = simulatePpv()
	pv.Pprod = simulatePprod()
}

func (pv *Pv) Show (){
	fmt.Println(" PV { Ppv :", pv.Ppv, "W ; Pprod :", pv.Pprod, "W/m² }")
}