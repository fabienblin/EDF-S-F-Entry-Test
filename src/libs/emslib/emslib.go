package emslib

import(
	// "fmt"
)

// Physical units
type Watt float64 // W
type KWatt float64 // kW
type KWattHour float64 // kWh
type WattPerSqrMeter float64 // W/m²

// unit convertion function
func WattToKWatt(watt Watt) KWatt {
	return KWatt(watt / 1000)
}

// unit convertion function
func KWattToWatt(kwatt KWatt) Watt {
	return Watt(kwatt * 1000)
}