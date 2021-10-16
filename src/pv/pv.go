package pv

import (
	. "libs/emslib"
	"libs/environment"
	"fmt"
)

// constants can be changed to create production scenarios
const nbSolarPanels int = 100
const solarPanelSurface int = 6
const solarTransformationEfficiency float64 = 0.2

// defines solar park total power output in watts
func simulatePpv() Watt{
	return Watt(environment.GetSunShine() * solarTransformationEfficiency * float64(nbSolarPanels))
}

// simulates pyranometer's power output estimate in watts/m² (arbitrary value)
// IRL this is a measure, not a calculation
func simulatePprod() WattPerSqrMeter{
	return WattPerSqrMeter(environment.GetSunShine() * 10)
}

// if we had motorized solar panels, the sun's hour angle would be usefull: https://en.wikipedia.org/wiki/Sunrise_equation
// tracking the sun's position in the sky would allow to optimize the solarTransformationEfficiency
func sunHourAngle(hour int) int{
	return 0
}

// define PV variables depending on environment conditions
func SimulatePv(pv *Pv) {
	pv.Ppv = simulatePpv()
	pv.Pprod = simulatePprod()
}

func ShowPv(pv *Pv){
	fmt.Println("Pv : Ppv : ", pv.Ppv, "W ; Pprod : ", pv.Pprod, "W/m²")
}