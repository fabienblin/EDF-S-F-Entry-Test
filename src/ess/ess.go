package ess

import (
	"fmt"
	. "libs/emslib"
	// "libs/environment"
	"math/rand"
	// "time"
)

type Ess struct {
	Pess      KWatt
	Pmaxch    KWatt
	Pmaxdisch KWatt
	Eess      KWattHour
}

// constants can be changed to create scenarios
const nbPowerBank int = 10
const powerBankCapacity KWattHour = 1.5

// for initialization
func randomizeEssCharge() KWattHour{
	return KWattHour(rand.Float64()) * powerBankCapacity * KWattHour(nbPowerBank)
}

func InitializeEss(ess *Ess){
	ess.Pess = 0
	ess.Pmaxch = .5
	ess.Pmaxdisch = .7
	ess.Eess = randomizeEssCharge()
}

func (ess *Ess) Show() {
	fmt.Print(" Ess : ")
	fmt.Print("Pess : ", ess.Pess, "kW ; ")
	fmt.Print("Pmaxch : ", ess.Pmaxch, "kW ; ")
	fmt.Print("Pmaxdisch : ", ess.Pmaxdisch, "kW ; ")
	fmt.Println("Eess : ", ess.Eess, "kWh")
}