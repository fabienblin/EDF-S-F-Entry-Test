package ess

import (
	"fmt"
	. "libs/emslib"
	"math/rand"
	"math"
	"poc"
)

type Ess struct {
	Pess      KWatt
	Pmaxch    KWatt
	Pmaxdisch KWatt
	Eess      KWattHour
	SetpointPEss KWatt
	IsFull bool
}

// constants can be changed to create scenarios
const essCapacity KWattHour = 20

// for initialization
func randomizeEssCharge() KWattHour {
	return KWattHour(rand.Float64()) * essCapacity
}

// arbitrary Pmaxch and Pmaxdisch
func InitializeEss(ess *Ess) {
	ess.Pess = 0
	ess.Pmaxch = -4 // must be negative
	ess.Pmaxdisch = 8
	if essChargeQuery() {
		ess.Eess = randomizeEssCharge()
	}
}

// for initialization
func essChargeQuery() bool {
	var userInput string
	var validIput bool
	var choice bool
	
	for !validIput {
		fmt.Print("Do you want to initialize the ESS with random charge ? (y/n) : ")
		fmt.Scanln(&userInput)
		if userInput == "y" {
			choice = true
			validIput = true
		} else if userInput == "n" {
			choice = false
			validIput = true
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
	fmt.Print("SetPoint : ", ess.SetpointPEss, " kW } ")
	if ess.IsFull {
		fmt.Println("FULL")
	} else if ess.Eess == 0 {
		fmt.Println("EMPTY")
	} else if ess.Pess < 0 {
		fmt.Println("CHARGING")
	} else if ess.Pess > 0 {
		fmt.Println("DISCHARGING")
	} else {
		fmt.Print("\n")
	}
}

// Pess is negative if charging : add setpoint to eess (limited by pmaxch and pPoc)
// Pess is positive if discharging : remove setpoint from eess (limited by pmaxdisch and eess)
// setpoint >= Pess >= -setpoint is always true
func (ess *Ess) simulatePess(poc *poc.Poc) {
	var isCharging bool = ess.SetpointPEss < 0
	var simPess KWatt = ess.SetpointPEss
	
	if isCharging {
		// pPoc := poc.Ppoc
		if ess.SetpointPEss < ess.Pmaxch {
			simPess = ess.Pmaxch
		}
		if ess.Pmaxch < poc.Ppoc && poc.Ppoc < 0 {
			simPess = poc.Ppoc
		}
	} else {
		if ess.SetpointPEss > ess.Pmaxdisch {
			simPess = ess.Pmaxdisch
		}
		if KWatt(ess.Eess) < simPess {
			simPess = KWatt(ess.Eess)
		}
	}

	if simPess == -0 {
		simPess = 0
	}

	ess.Pess = simPess
}

// add -pess to eess (pess is negative if charging)
// essCapacity >= eess >= 0 is always true
// return excess energy
func (ess *Ess) simulateEess() {
	ess.Eess -= KWattHour(ess.Pess)

	var excess KWatt = KWatt(ess.Eess - essCapacity)
	if excess > 0 {
		ess.Eess = essCapacity
		ess.IsFull = true
		if math.Abs(float64(excess + ess.Pess)) < 1e-10 {
			ess.Pess = 0
		} else {
			ess.Pess += excess // excess power is not used
		}
	}

	if ess.Eess < 0 {
		ess.Eess = 0
	}
}

// simulate charge or discharge
// return power production/demand
func (ess *Ess) SimulateEss(poc *poc.Poc){
	ess.simulatePess(poc)
	ess.simulateEess()

	poc.SimulatePoc(ess.Pess)
}

func (ess *Ess) Reset() {
	ess.Pess = 0
}