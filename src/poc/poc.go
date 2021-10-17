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
var pload KWatt

// deifine hidden pload, ranging from 0 to 5 kW
func simulatePload() {
	pload = KWatt(rand.Float64()) * 5
}

// simulate visible Ppoc
func (poc *Poc) simulatePpoc(pAcBus KWatt) {
	poc.Ppoc = pAcBus + pload
}

func (poc *Poc) SimulatePoc(pAcBus KWatt) {
	simulatePload()
	poc.simulatePpoc(pAcBus)
}

func (poc *Poc) Show() {
	fmt.Println(" POC { Ppoc :", poc.Ppoc, "kW }")
}

func ShowFacility() {
	fmt.Println("FACILITY {")
	fmt.Println(" Pload :", pload, "kW")
	fmt.Println("}")
}