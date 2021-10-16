package main

import (
	"fmt"
	. "libs/emslib"
	"libs/environment"
)

// global constants can be changed to create production scenarios
const nbSolarPanels int = 100
const solarPanelSurface int = 6
const solarTransformationEfficiency float64 = 0.2

// defines solar park total power output in watts
func simulatePpv() Watt{
	return Watt(environment.GetSunShine() * solarTransformationEfficiency * float64(nbSolarPanels))
}

// simulates pyranometer's power output estimate in watts/m² for all solar panels
// IRL this is a measure, not a calculation
func simulatePprod() float64{
	return environment.GetSunShine() / float64(solarPanelSurface * nbSolarPanels)
}

// if we had motorized solar panels, the sun's hour angle would be usefull: https://en.wikipedia.org/wiki/Sunrise_equation
// tracking the sun's position in the sky would allow to optimize the solarTransformationEfficiency
func sunHourAngle(hour int) int{
	return 0
}

func main() {
	var pv *Pv = new(Pv)
	var userInput string

	for true {
		fmt.Println("Press a key to process time cycle.")
		fmt.Scanln(&userInput)

		pv.Ppv = simulatePpv()
		pv.Pprod = simulatePprod()
		fmt.Println(pv)
		environment.ShowEnvironment()
		
		environment.NextHour()
	}
}
