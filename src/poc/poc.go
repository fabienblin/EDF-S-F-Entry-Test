package main

import (
	"fmt"
	. "libs/emslib"
	"libs/environment"
	"math/rand"
	"time"
)

// protected var must not be used other than with simulateFacilityConsumption() for exercice purposes
var facilityConsumption KWatt

// ranging from 0 to 5 kW
func simulateFacilityConsumption() {
	facilityConsumption =  KWatt(rand.Float64()) * 5
}

func simulatePpoc() KWatt {
	return 0
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	var poc *Poc = new(Poc)
	var userInput string

	for true {
		fmt.Println("Press a key to process time cycle.")
		fmt.Scanln(&userInput)

		simulateFacilityConsumption()

		poc.Ppoc = simulatePpoc()
		fmt.Println(poc)
		
		environment.NextHour()
	}

}
