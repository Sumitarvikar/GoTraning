package main

// Qus - Write a program to calculate the area of the circle, First line has a value of the radius of the circle
// constraint
// 1. Use const PI from the package math package
// 2. Use the Pow function from the math package
// 3. Round area float value to 2 decimal places
import (
	"fmt"
	"math"
)

func main() {
	var radius float64
	fmt.Print("Enter the radius of the circle: ")
	_, err := fmt.Scanln(&radius)
	if err != nil {
		fmt.Println("Error: Invalid input.")
		return
	}
	area := math.Pi * math.Pow(radius, 2)
	roundedArea := math.Round(area*100) / 100
	fmt.Printf("The area of the circle is: %.2f\n", roundedArea)
}
