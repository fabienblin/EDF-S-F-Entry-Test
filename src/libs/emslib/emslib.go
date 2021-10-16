package emslib

import(
	"fmt"
)

// Physical units
type Watt float64 // W
type KWatt float64 // kW
type KWattHour float64 // kWh
type WattPerSqrMeter float64 // W/m²

// ESS definitions
type Ess struct {
	Pess      KWatt
	Pmaxch    KWatt
	Pmaxdisch KWatt
	Eess      KWattHour
}

// PV definitions
type Pv struct {
	Ppv   Watt
	Pprod WattPerSqrMeter
}

// POC definitions
type Poc struct {
	Ppoc KWatt
}

// EMS definition
type Ems struct {
	Ess  *Ess
	Pv   *Pv
	Poc  *Poc
	Pess Watt
	Ppv  Watt
}

func (ems *Ems) SetpointPEss(setpointPEss Watt) {
	ems.Pess = setpointPEss
}

func (ems *Ems) SetpointPPv(setpointPPv Watt) {
	ems.Ppv = setpointPPv
}

func (ems *Ems) GetEssMeasure() *Ess {
	return ems.Ess
}

func (ems *Ems) GetPvMeasure() *Pv {
	return ems.Pv
}

func (ems *Ems) GetPocMeterMeasure() *Poc {
	return ems.Poc
}

func (ems *Ems) Print() {
	fmt.Println("\nEMS {")
	fmt.Println(" ess  : ", ems.Ess)
	fmt.Println(" pv   : ", ems.Pv)
	fmt.Println(" poc  : ", ems.Poc)
	fmt.Println(" pEss : ", ems.Pess)
	fmt.Println(" pPv  : ", ems.Ppv)
	fmt.Println("}")
}

// unit convertion function
func WattToKWatt(watt Watt) KWatt {
	return KWatt(watt / 1000)
}

// unit convertion function
func KWattToWatt(kwatt KWatt) Watt {
	return Watt(kwatt * 1000)
}