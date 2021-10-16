package main

import (
	"fmt"
	. "libs/emslib"
	"libs/environment"
)

func initEnv() *Ems{
	// ess := new(Ess)
	// pv := new(Pv)
	// poc := new(Poc)
	ems := new(Ems)

	// ems.ess = ess
	// ems.pv = pv
	// ems.poc = poc

	return ems
}

func main() {
	ems := initEnv()
	var userInput string
	
	for true {
		// get user input
		fmt.Println("Press a key to process time cycle.")
		fmt.Scanln(&userInput)

		//Read power productions and demands
		
		// core AI descision making
		
		// show all power levels
		fmt.Println("Current time : ", environment.GetHour())
		ems.Print()

		// add 1 hour to day/night cycle
		environment.NextHour()
	}
}
