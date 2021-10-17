package ess

import (
	"fmt"
	. "libs/emslib"
	// "libs/environment"
	"math/rand"
	// "time"
)

type Ess struct {
	Pess      KWatt
	Pmaxch    KWatt
	Pmaxdisch KWatt
	Eess      KWattHour
	SetpointPEss KWatt
}

// constants can be changed to create scenarios
const essCapacity KWattHour = 1.5

// for initialization
func randomizeEssCharge() KWattHour {
	return KWattHour(rand.Float64()) * essCapacity
}

// arbitrary Pmaxch and Pmaxdisch
func InitializeEss(ess *Ess) {
	ess.Pess = 0
	ess.Pmaxch = -1 // must be negative
	ess.Pmaxdisch = 1
	if essChargeQuery() {
		ess.Eess = randomizeEssCharge()
	}
}

// for initialization
func essChargeQuery() bool {
	var userInput string
	var invalidIput bool = true
	var choice bool
	
	for invalidIput {
		fmt.Print("Do you want to initialize the ESS with random charge ? (y/n) : ")
		fmt.Scanln(&userInput)
		if userInput == "y" {
			choice = true
			invalidIput = false
		} else if userInput == "n" {
			choice = false
			invalidIput = false
		} else {
			fmt.Println("Invalid answer. Try again.")
		}
	}
	return choice
}
		
func (ess *Ess) Show() {
	fmt.Print(" Ess { ")
	fmt.Print("Pess : ", ess.Pess, " kW ; ")
	fmt.Print("Pmaxch : ", ess.Pmaxch, " kW ; ")
	fmt.Print("Pmaxdisch : ", ess.Pmaxdisch, " kW ; ")
	fmt.Print("Eess : ", ess.Eess,  " kWh ; ")
	fmt.Print("SetPoint : ", ess.SetpointPEss, " kW }\n")
}

// Pess is negative if charging : add setpoint to eess (limited by pmaxch)
// Pess is positive if discharging : remove setpoint from eess (limited by pmaxdisch)
// setpoint >= Pess >= -setpoint is always true
func (ess *Ess) simulatePess() {
	var isCharging bool = ess.SetpointPEss < 0
	var simPess KWatt = ess.SetpointPEss

	if isCharging {
		if ess.SetpointPEss > ess.Pmaxch {
			simPess = ess.Pmaxch
		}
	} else {
		if ess.SetpointPEss > ess.Pmaxdisch {
			simPess = ess.Pmaxdisch
		}
		if KWatt(ess.Eess) < simPess {
			simPess = KWatt(ess.Eess)
		}
	}

	ess.Pess = simPess
}

// add -pess to eess (pess is negative if charging)
// essCapacity >= eess >= 0 is always true
// return excess energy
func (ess *Ess) simulateEess() KWatt{
	ess.Eess -= KWattHour(ess.Pess)

	var excess KWatt = KWatt(ess.Eess - essCapacity)
	if excess > 0 {
		ess.Eess = essCapacity
		return excess
	}

	if ess.Eess < 0 {
		ess.Eess = 0
	}

	return 0
}

// simulate charge or discharge
func (ess *Ess) SimulateEss() KWatt{
	ess.simulatePess()
	excess := ess.simulateEess()

	return excess
}