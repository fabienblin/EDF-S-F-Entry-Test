package poc

import (
	"fmt"
	. "libs/emslib"
	"libs/environment"
	"math"
	"math/rand"
)

type Poc struct {
	Ppoc KWatt
}

// change me
const facilityConsumptionFactor KWatt = 2

// negative value
var pload KWatt

// deifine hidden pload, ranging from 0 to -5 kW
func simulatePload() {
	// ranges from -0.1 to -2.1
	simPload := -KWatt(math.Sin(((float64(environment.GetHour() - 6)) * math.Pi) / 12) + 1.1)
	
	simPload -= KWatt(rand.Float64()) // add variations ranging from 0 to -1

	simPload *= facilityConsumptionFactor

	pload = simPload
}

// PmaxSite < Ppoc <= 0 must always be true  ; this can't be controlled from here
func (poc *Poc) simulatePpoc(pAcBus KWatt) {
	poc.Ppoc += pAcBus
}

func (poc *Poc) SimulatePoc(pAcBus KWatt) KWatt{
	poc.simulatePpoc(pAcBus)

	return poc.Ppoc
}

func (poc *Poc) Show() {
	fmt.Println(" POC { Ppoc :", poc.Ppoc, "kW }")
}

func ShowFacility() {
	fmt.Println("FACILITY {")
	fmt.Println(" Pload :", pload, "kW")
	fmt.Println("}")
}

func (poc *Poc) Reset() {
	simulatePload()
	poc.Ppoc = pload
}