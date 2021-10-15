package emslib

type Watt float64

type Ess struct {
	Pess      Watt
	Pmaxch    Watt
	Pmaxdisch Watt
	Eess      Watt
}

type Pv struct {
	Ppv   Watt
	Pprod Watt
}

type Poc struct {
	Ppoc Watt
}

func Void(){

}