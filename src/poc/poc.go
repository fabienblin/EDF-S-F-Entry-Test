package poc

import (
	"fmt"
	. "libs/emslib"
	"math/rand"
)

type Poc struct {
	Ppoc KWatt
}

// protected var must not be used other than with simulateFacilityConsumption() for exercice purposes
var facilityConsumption KWatt

// ranging from 0 to 5 kW
func simulateFacilityConsumption() {
	facilityConsumption =  KWatt(rand.Float64()) * 5
}

func simulatePpoc() KWatt {
	return 0
}

func (poc *Poc) Show() {
	fmt.Println(" Poc : Ppoc : ", poc.Ppoc, "W")
}
