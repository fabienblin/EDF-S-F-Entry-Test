package poc

import (
	"fmt"
	. "libs/emslib"
	"math/rand"
)

type Poc struct {
	Ppoc KWatt
}

const maxPload KWatt = 5

// negative value
var pload KWatt

// deifine hidden pload, ranging from 0 to -5 kW
func simulatePload() {
	pload = -KWatt(rand.Float64()) * maxPload
}

// PmaxSite < Ppoc <= 0 must always be true  ; this can't be controlled from here
func (poc *Poc) simulatePpoc(pAcBus KWatt) {
	poc.Ppoc += pAcBus
}

func (poc *Poc) SimulatePoc(pAcBus KWatt) KWatt{
	// simulatePload()
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
	poc.Ppoc = pload
}