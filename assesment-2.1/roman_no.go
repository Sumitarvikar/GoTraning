package main

import (
	"fmt"
	"strings"
)

func romanToInt(s string) int {
	romanValues := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	total := 0
	length := len(s)

	for i := 0; i < length; i++ {
		current := romanValues[rune(s[i])]
		if i+1 < length && romanValues[rune(s[i+1])] > current {
			total -= current
		} else {
			total += current
		}
	}

	return total
}

func main() {
	var romanInput string
	fmt.Println("Enter a Roman numeral:")
	fmt.Scan(&romanInput)
	romanInput = strings.ToUpper(romanInput)
	result := romanToInt(romanInput)
	fmt.Printf("The integer value of %s is %d\n", romanInput, result)
}
