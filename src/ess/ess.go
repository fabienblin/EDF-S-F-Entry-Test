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
}

// constants can be changed to create scenarios
const essCapacity KWattHour = 1.5

// for initialization
func randomizeEssCharge() KWattHour{
	return KWattHour(rand.Float64()) * essCapacity
}

// arbitrary Pmaxch and Pmaxdisch
func InitializeEss(ess *Ess){
	ess.Pess = 0
	ess.Pmaxch = 1
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
		if userInput == "y"{
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
	fmt.Print("Eess : ", ess.Eess, "kWh }\n")
}

// simulate charge or discharge
func (ess *Ess) SimulateEss() {
	ess.Eess = ess.Eess + KWattHour(ess.Pess)
}