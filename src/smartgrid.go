package main

import (
	"fmt"
	. "libs/emslib"
	"libs/environment"
	"math/rand"
	"math"
	"time"
	"ems"
	"poc"
	"github.com/fatih/color"
)

// constants can be changed to create scenarios
const pMaxSite KWatt = -8

// verify power distribution
// knowing (pMaxSite < Ppoc <= 0) must be true at all times
func verifyPowerLevels(poc *poc.Poc) {
	var draining bool = (pMaxSite >= poc.Ppoc && math.Abs(float64(pMaxSite - poc.Ppoc)) > 1e-10)
	var overload bool = (poc.Ppoc > 0)

	if  overload || draining {
		color.Set(color.FgRed)
		fmt.Println("! FAILURE !")
		if overload{
			fmt.Print("OVERLOAD : Ppoc > 0 : ")
			fmt.Println(poc.Ppoc," > 0")
		}
		if draining {
			fmt.Print("DRAINING : pMaxSite >= Ppoc : ")
			fmt.Println(pMaxSite, " >= ", poc.Ppoc)
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
		EMS.Poc.Reset(true) // new pload
		EMS.Pv.SimulatePv(EMS.Poc, true) // new sunshine
		EMS.Ess.SimulateEss(EMS.Poc)
		
		// core AI descision making
		EMS.Ai()

		EMS.Poc.Reset(false)
		EMS.Pv.SimulatePv(EMS.Poc, false)
		EMS.Ess.SimulateEss(EMS.Poc)

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