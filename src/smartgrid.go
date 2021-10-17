package main

import (
	"fmt"
	. "libs/emslib"
	"libs/environment"
	"math/rand"
	"time"
	// "pv"
	// "ess"
	"ems"
	"poc"
	"github.com/fatih/color"
)

// constants can be changed to create scenarios
const pMaxSite KWatt = 5

// summ of Pess and Ppv
var pAcBus KWatt
func simulatePAcBus(Ppv Watt, Pess KWatt, essExcess KWatt) {
	pAcBus = Pess + WattToKWatt(Ppv) + essExcess
}

// verify power distribution
// knowing (pMaxSite < Ppoc <= 0) must be true at all times
func verifyPowerLevels(poc *poc.Poc) {
	var overload bool = (pMaxSite <= poc.Ppoc)
	var draining bool = (poc.Ppoc < 0)

	if  overload || draining {
		color.Set(color.FgRed)
		fmt.Println("! FAILURE !")
		if overload {
			fmt.Print("OVERLOAD : pMaxSite >= Ppoc : ")
			fmt.Println(pMaxSite, " >= ", poc.Ppoc)
		}
		if draining {
			fmt.Print("DRAINING : Ppoc > 0 : ")
			fmt.Println(poc.Ppoc," > 0")
		}
		color.Unset()
	} else {
		color.Green("Smart Grid Not Overflowing Nor Draining")
	}
}

func main () {
	rand.Seed(time.Now().UTC().UnixNano())
	var userInput string
	EMS := new(ems.Ems).IntializeEms(pMaxSite)

	// 1 iteration = 1 hour
	for true {
		// simulate power supplies and demands
		environment.SimulateEnvironment()
		EMS.Pv.SimulatePv()
		essExcess := EMS.Ess.SimulateEss()
		simulatePAcBus(EMS.Pv.Ppv, EMS.Ess.Pess, essExcess)
		EMS.Poc.SimulatePoc(pAcBus)
		
		// core AI descision making
		EMS.Ai()
		EMS.SetpointPEss(1)
		
		// log environment and smart grid
		environment.ShowEnvironment()
		poc.ShowFacility()
		EMS.Show()

		// log problems
		verifyPowerLevels(EMS.Poc)

		// next hour on user input
		fmt.Println("Press ENTER to process time cycle.")
		fmt.Scanln(&userInput)
		environment.NextHour()
	}
}