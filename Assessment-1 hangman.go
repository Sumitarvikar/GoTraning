package main

import (
	"fmt"
	"strings"
)

func main() {
	word := "elephant"

	// lookup for entries made by the user.
	entries := map[string]bool{}

	// list of "_" corrosponding to the number of letters in the word. [ _ _ _ _ _ ]
	placeholder := []string{}
	for i := 0; i < len(word); i++ {
		placeholder = append(placeholder, "_")
	}

	chances := 8
	for {
		fmt.Println("*****************************")
		// evaluate a loss! If user guesses a wrong letter or the wrong word, they lose a chance.
		inputWord := strings.Join(placeholder, "")
		if chances == 0 && (word != inputWord) {
			fmt.Println("Lost the game")
			break
		}
		// evaluate a win!
		if word == inputWord {
			fmt.Println(placeholder)
			fmt.Println("Win the game")
			break
		}
		// Console display
		fmt.Printf("\n")
		fmt.Println(placeholder)
		fmt.Printf("\n")                     // render the placeholder
		fmt.Printf("chances :  %d", chances) // render the chances left
		fmt.Printf("\n")
		fmt.Printf("Guess a letter :")

		for i := range entries {
			fmt.Printf("%s", i)
		}
		fmt.Printf("\n")
		fmt.Println("enter char :")
		str := ""
		_, err := fmt.Scanln(&str)

		if err != nil || len(str) > 1 {
			fmt.Println("*****************************")
			fmt.Println("Please enter valid input")
			continue
		}

		// compare and update entries, placeholder and chances.
		for i := range word {
			c := string(word[i])
			if str == c && placeholder[i] != c {
				entries[c] = true
				placeholder[i] = c
			}
		}

		if ok := entries[str]; !ok {
			entries[str] = true
			chances--
		}

	}
}
