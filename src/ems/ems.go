package ems

import (
	"fmt"
	. "libs/emslib"
	"ess"
	"pv"
	"poc"
)

type Ems struct {
	Ess  *ess.Ess
	Pv   *pv.Pv
	Poc  *poc.Poc
	Pess Watt
	Ppv  Watt
}

func (ems *Ems) SetpointPEss(setpointPEss Watt) {
	ems.Pess = setpointPEss
}

func (ems *Ems) SetpointPPv(setpointPPv Watt) {
	ems.Ppv = setpointPPv
}

func (ems *Ems) GetEssMeasure() *ess.Ess {
	return ems.Ess
}

func (ems *Ems) GetPvMeasure() *pv.Pv {
	return ems.Pv
}

func (ems *Ems) GetPocMeterMeasure() *poc.Poc {
	return ems.Poc
}

func (ems *Ems) Show() {
	fmt.Println("EMS {")
	ems.Ess.Show()
	ems.Pv.Show()
	ems.Poc.Show()
	fmt.Println(" pEss : ", ems.Pess)
	fmt.Println(" pPv  : ", ems.Ppv)
	fmt.Println("}")
}

func IntializeEms() *Ems{
	var ems *Ems = new(Ems)
	
	var ESS *ess.Ess = new(ess.Ess)
	var poc *poc.Poc = new(poc.Poc)
	var pv *pv.Pv = new(pv.Pv)

	ems.Ess = ESS
	ems.Poc = poc
	ems.Pv = pv

	ess.InitializeEss(ems.Ess)

	return ems
}