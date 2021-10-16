package environment

import(
	"fmt"
	"math"
	"math/rand"
	"time"
)

// initialization var
var initial bool = true
// hour of the day ranges between 0 and 23
// protected variable (read only) accessible via GetHour()
var hour int
// day weather changes every day, ranges from 0.2 to 1.0
var dayWeather float64
// hour weather simulates clouds depending on dayWeather
var hourWeather float64
// simulated sun shine ranges from 0.0 to 100.0 minus hourWeather
// proteced variable accessible via GetSunShine()
var sunShine float64


func ShowEnvironment() {
	if initial {
		fmt.Println("Initializing...")
	} else {
		fmt.Println("ENVIRONMENT {")
		fmt.Println(" hour         : ", hour)
		fmt.Println(" dayWeather   : ", dayWeather)
		fmt.Println(" hourWeather  : ", hourWeather)
		fmt.Println(" sunShine     : ", sunShine)
		fmt.Println("}")
	}
}

// Add an hour to day/night cycle and update dayWeather every 24h
// returns updated hour
func NextHour() {
	// simulations order matters
	// 0: increment hour
	// 1: simulateDayWeather
	// 2: simulateHourWeather
	// 3: simulateSunShine

	if initial == false { // setHour(+1) starts at second cycle for initialization surposes
		setHour(hour + 1)
	} else {
		rand.Seed(time.Now().UTC().UnixNano())
	}
	initial = false
	if hour == 0 { // every 24h
		simulateDayWeather()
	}
	simulateHourWeather()
	simulateSunShine()
}

func GetHour() int{
	return hour
}

func GetSunShine() float64{
	return sunShine
}

// protected function
// returns updated hour
func setHour(time int) int {
	if time <= 24 && time >= 0 {
		if time == 24{
			hour = 0
		} else {
			hour = time
		}
	}
	return hour
}

// sinusoidal function that simulates sunshine, randomized with weather conditions.
func simulateSunShine() {
	var simSunShine float64
	
	simSunShine = (math.Sin((float64((hour - 6)) * math.Pi) / 12) + 0.1) / 1.1 // range -0.8 to 1.0 ; max value with hour=12
	
	if simSunShine < 0 { // range 0.0 to 1.0
		simSunShine = 0
	}

	simSunShine = simSunShine * 100 // range 0.0 to 100.0
	
	simSunShine = simSunShine - (simSunShine * hourWeather) / 100 // apply weather conditions

	sunShine = simSunShine
}

// ranges from 20 (verry bad weather) to 100 (perfect weather)
func simulateDayWeather() {
	dayWeather = (rand.Float64() + 0.2) * 100
	fmt.Println(dayWeather)
	if dayWeather > 100 {
		dayWeather = 100
	}
}

// makes small variations (20% max) to the dayWeather, ranges from 16 to 80
func simulateHourWeather() {
	hourWeather = dayWeather - (dayWeather * rand.Float64() * 0.2)
}