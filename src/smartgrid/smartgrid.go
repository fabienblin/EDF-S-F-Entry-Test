package main

import (
	"fmt"
	. "libs/emslib"
	"libs/environment"
	"math/rand"
	"time"
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

	for true {
		//Read power productions and demands
		// ceci est completement con :
		ems.Ess = ems.GetEssMeasure()
		ems.Pv = ems.GetPvMeasure()
		ems.Poc = ems.GetPocMeterMeasure()
		// core AI descision making
		
		// show all power levels

		// next hour on user input
		fmt.Println("Press a key to process time cycle.")
		fmt.Scanln(&userInput)
		environment.NextHour()

	}
}