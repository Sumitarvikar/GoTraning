package main

// Qus - 1. Write a program to calculate the simple interest
// First-line has the comma-separated values of the Principal, rate and time (in years) respective
// *constraints: Round simple interest float value to 2 decimal places

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	var input string
	fmt.Print("Enter principal, rate of interest (in %), and time (in years), separated by commas: ")
	fmt.Scanln(&input)

	input = strings.ReplaceAll(input, " ", "")
	parts := strings.Split(input, ",")

	if len(parts) != 3 {
		fmt.Println("Please enter valid values")
		return
	}

	principal, err := strconv.ParseFloat(parts[0], 64)
	if err != nil {
		fmt.Println("Error: Invalid principal value.")
		return
	}

	rate, err := strconv.ParseFloat(parts[1], 64)
	if err != nil {
		fmt.Println("Error: Invalid rate of interest value.")
		return
	}

	time, err := strconv.ParseFloat(parts[2], 64)
	if err != nil {
		fmt.Println("Error: Invalid time value.")
		return
	}

	interest := (principal * rate * time) / 100

	roundedInterest := math.Round(interest*100) / 100

	fmt.Printf("The simple interest is: %.2f\n", roundedInterest)
}
