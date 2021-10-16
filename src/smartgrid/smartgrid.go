package main

import (
	"fmt"
	// . "libs/emslib"
	"libs/environment"
	"math/rand"
	"time"
	"pv"
	// "ess"
	"ems"
)

func main () {
	rand.Seed(time.Now().UTC().UnixNano())
	var userInput string
	EMS := ems.IntializeEms()

	for true {
		// simulate PV production
		pv.SimulatePv(EMS.Pv)
		
		// read power productions and demands
		
		// core AI descision making
		
		// show all power levels
		environment.ShowEnvironment()
		EMS.Show()

		// next hour on user input
		fmt.Println("Press a key to process time cycle.")
		fmt.Scanln(&userInput)
		environment.NextHour()

	}
}