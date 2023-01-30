//windchill prompts the collection of a temperature in Fahrenheit and wind-speed in miles per hour, returning
//the estimated wind chill. The -table flag may be used to view a wind chill table instead of being prompted for a calculation.
package main

import (
	"flag"
	"fmt"
	"math"
)

// Constants
const degreeF = string(0x00B0) + "F" // Fahrenheit unit

// Flags
var table = flag.Bool("table", false, "Display Wind Chill Table instead of prompting for a calculation.")

func main() {
	flag.Parse()

	if *table {
		displayTable()
	} else {
		temp := collectTemp()           // an address
		windSpeed := collectWindSpeed() //an address

		// display the calculated wind chill
		fmt.Printf("\nEstimated Wind Chill: %5.2f %s\n", calcWindChill(*temp, *windSpeed), degreeF) // pass the values, not the address
	}
}

//collectTemp continually prompts the collection of a temperature value in range [-58, 41] until a
//valid value is provided and returns the address of the collected value.
func collectTemp() *float64 {
	var temp float64
	for {
		fmt.Printf("Temp (%s): ", degreeF)
		_, err := fmt.Scan(&temp)
		// validate input
		if err != nil {
			fmt.Printf("[ERROR 1] Invalid input. Please provide a number between -58%s and 41%s. Got: %v\n", degreeF, degreeF, err)
			continue
		}
		if (temp < -58) || (temp > 41) {
			fmt.Printf("[ERROR 2] Invalid input. %v %s is out of range. Please provide a number between -58%s and 41%s.\n", temp, degreeF, degreeF, degreeF)
			continue
		}
		break
	}
	return &temp // returns an address
}

//collectWindSpeed continually prompts the collection of a wind speed value greater than or equal to 2 until a
//valid value is provided and returns the address of the collected value.
func collectWindSpeed() *float64 {
	var windSpeed float64
	for {
		fmt.Printf("Wind Speed (mph): ")
		_, err := fmt.Scan(&windSpeed)
		if err != nil {
			fmt.Printf("[ERROR 3] Invalid input. Please provide a number greater than or equal to 2 mph. Got: %v\n", err)
			continue
		}
		if windSpeed < 2 {
			fmt.Printf("[ERROR 4] Invalid input. %v is out of range. Please provide a number greater than or equal to 2.\n", windSpeed)
			continue
		}
		break
	}
	return &windSpeed // returns an address
}

//calcWindChill calculates the wind chill given a temperature in Fahrenheit and wind-speed in miles per hour.
func calcWindChill(temp, windSpeed float64) float64 {
	exp := math.Pow(windSpeed, 0.16)
	return 35.74 + (0.6215 * temp) - (35.75 * exp) + (0.4275 * temp * exp)
}

//createTable builds and returns the Wind Chill Chart.
// Note: windChill and temp are iterated once before they're used, thus they are initialized with an offset.
func createTable() [13][19]float64 {
	var wcTable [13][19]float64

	windChill := 0.0          // never resets, we only go through the rows once
	for i := 0; i < 13; i++ { // row number
		temp := 45.0              // resets every row
		for j := 0; j < 19; j++ { // column number
			// temp axis labels
			if i == 0 && j > 0 {
				wcTable[i][j] = temp
			}
			// wind chill axis labels
			if i > 0 && j == 0 {
				wcTable[i][j] = windChill
			}
			// the calculated table values
			if i > 0 && j > 0 {
				wcTable[i][j] = calcWindChill(float64(temp), float64(windChill))
			}
			temp -= 5 // next column, iterate temp
		}
		windChill += 5 // next row, iterate wind speed
	}
	return wcTable
}

//displayTable prints out the Wind Chill Chart. Calls createTable to create the Wind Chill Chart.
func displayTable() {
	// user friendly stuff...
	fmt.Printf("X-Axis (first row): Temperature (%s)\n", degreeF)
	fmt.Printf("Y-Axis (first column): Wind Speed (mph)\n\n")

	table := createTable() // gotta make the table before displaying...

	// begin displaying the table item by item
	for i := 0; i < 13; i++ {
		for j := 0; j < 19; j++ {
			// for formatting axis labels
			if i == 0 {
				fmt.Printf(" [%3v] ", table[i][j])
				continue
			}
			if i > 0 && j == 0 {
				fmt.Printf(" [%3v] ", table[i][j])
				continue
			}
			// not an axis label
			fmt.Printf(" %5.1f ", table[i][j])
		}
		fmt.Println() // new row
	}
}
