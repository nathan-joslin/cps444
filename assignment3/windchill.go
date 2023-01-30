//windchill prompts the collection of a temperature in Fahrenheit and wind-speed in miles per hour, returning
//the estimated wind chill. The -table flag may be used to view a wind chill table instead.
package main

import (
	"fmt"
	"math"
)

func main() {
	// Get temp value
	temp := collectTemp() // an address

	// Get wind speed value
	windSpeed := collectWindSpeed() //an address

	fmt.Printf("Recieved: Temp(%v) Wind-Speed(%v).\n", *temp, *windSpeed)
	fmt.Println()

	displayTable()
}

//collectTemp continually prompts the collection of a temperature integer value in range [-58, 41] until a
//valid value is provided.
func collectTemp() *float64 {
	var temp float64
	for {
		fmt.Printf("Temp: ")
		_, err := fmt.Scan(&temp)
		if err != nil {
			fmt.Printf("[ERROR 1] Invalid input. Please provide a number between -58F and 41F. Got: %v\n", err)
			continue
		}
		if (temp < -58) || (temp > 41) {
			fmt.Printf("[ERROR 2] Invalid input. %v is out of range. Please provide a number between -58F and 41F.\n", temp)
			continue
		}
		break
	}
	return &temp
	// fmt.Printf("Recieved temp: %v\n", temp)
}

//collectWindSpeed continually prompts the collection of a wind speed integer value greater than or equal to 2 until a
//valid value is provided.
func collectWindSpeed() *float64 {
	var windSpeed float64
	for {
		fmt.Printf("Wind Speed: ")
		_, err := fmt.Scan(&windSpeed)
		if err != nil {
			fmt.Printf("[ERROR 3] Invalid input. Please provide a number greater than or equal to 2. Got: %v\n", err)
			continue
		}
		if windSpeed < 2 {
			fmt.Printf("[ERROR 4] Invalid input. %v is out of range. Please provide a number greater than or equal to 2.\n", windSpeed)
			continue
		}
		break
	}
	return &windSpeed
	// fmt.Printf("Recieved wind speed: %v\n", windSpeed)
}

//calcWindChill calculates the wind chill given a temperature in Fahrenheit and wind-speed in miles per hour.
func calcWindChill(temp, windSpeed float64) float64 {
	exp := math.Pow(windSpeed, 0.16)
	return 35.74 + (0.6215 * temp) - (35.75 * exp) + (0.4275 * temp * exp)
}

//createTable builds and returns the Wind Chill Chart.
func createTable() [13][19]float64 {
	var wcTable [13][19]float64
	column := 1
	// fill in temperature axis
	for i := 40.0; i >= -45; i -= 5 {
		wcTable[0][column] = i
		column++
	}
	// fill in wind axis
	row := 1
	for i := 5.0; i <= 60; i += 5 {
		wcTable[row][0] = i
		row++
	}
	// TODO: fill in table values
	return wcTable
}

//displayTable prints out the Wind Chill Chart. Calls createTable to create the Wind Chill Chart.
func displayTable() {
	// user friendly stuff...
	fmt.Println("X-Axis: Temperature (F)")
	fmt.Println("Y-Axis: Wind Speed (mph)")

	table := createTable()

	// begin displaying the table item by item
	for i := 0; i < 13; i++ {
		for j := 0; j < 19; j++ {
			// for formatting axis labels
			if i == 0 {
				fmt.Printf(" [%v] ", table[i][j])
				continue
			}
			if i > 0 && j == 0 {
				fmt.Printf(" [%v] ", table[i][j])
				continue
			}
			// not an axis label
			fmt.Printf(" %v ", table[i][j])
		}
		// new row
		fmt.Println()
	}
}
