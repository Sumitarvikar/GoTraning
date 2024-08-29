package main

import (
	"fmt"
)

func findDay(index int) (day string) {
	daysInWeek := map[int]string{
		1: "Monday",
		2: "Tuesday",
		3: "Wednesday",
		4: "Thursday",
		5: "Friday",
		6: "Saturday",
		7: "Sunday",
	}
	dayValue := daysInWeek[index]
	if dayValue == "" {
		day = "Not a day"
	} else {
		day = dayValue
	}
	return
}

func main() {
	fmt.Println("Please enter day (index) : ")
	var index int
	fmt.Scanln(&index)
	fmt.Println(findDay(index))
}
