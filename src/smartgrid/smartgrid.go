package main

import (
	"fmt"
	. "libs/emslib"
	"libs/environment"
	"math/rand"
	"time"
	"pv"
	"ess"
)

func intializeEms() *Ems{
	var ems *Ems = new(Ems)
	
	var ess *Ess = new(Ess)
	var poc *Poc = new(Poc)
	var pv *Pv = new(Pv)

	ems.Ess = ess
	ems.Poc = poc
	ems.Pv = pv

	return ems
}

func main () {
	rand.Seed(time.Now().UTC().UnixNano())
	var userInput string
	ems := intializeEms()
	ess.InitializeEss(ems.Ess)

	for true {
		// simulate PV production
		pv.SimulatePv(ems.Pv)
		
		// read power productions and demands
		
		// core AI descision making
		
		// show all power levels
		environment.ShowEnvironment()
		pv.ShowPv(ems.Pv)
		ess.ShowEss(ems.Ess)

		// next hour on user input
		fmt.Println("Press a key to process time cycle.")
		fmt.Scanln(&userInput)
		environment.NextHour()

	}
}