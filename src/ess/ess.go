package main

import (
	"fmt"
	. "libs/emslib"
	"libs/environment"
	"math/rand"
	"time"
)


func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	var userInput string
	var ess *Ess = new(Ess)

	for true {
		fmt.Println("Press a key to process time cycle.")
		fmt.Scanln(&userInput)

		ess.Pess = simulatePess()
		ess.Pmaxch = simulatePmaxch()
		ess.Pmaxdisch = simulatePmaxdisch()
		ess.Eess = simulateEess()

		environment.NextHour()
	}
}
