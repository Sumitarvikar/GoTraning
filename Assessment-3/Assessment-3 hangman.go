package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func printGussedLetters(entries map[string]bool) {
	for i := range entries {
		fmt.Printf("%s , ", i)
	}
}

func compareWordandUpdatedGussedLetters(entries *map[string]bool, placeholder *[]string, chances *int, word string, inputWord string) {
	entriesWords := *entries
	placeholderWords := *placeholder
	for i := range word {
		c := string(word[i])
		if inputWord == c && placeholderWords[i] != c {
			entriesWords[c] = true
			placeholderWords[i] = c
		}
	}
	if ok := entriesWords[inputWord]; !ok {
		entriesWords[inputWord] = true
		*chances -= 1
	}
}

func printRequirdData(entries map[string]bool, placeholder []string, chances *int) {
	fmt.Printf("\n")
	fmt.Println(placeholder)
	fmt.Printf("\n")
	chancesWord := *chances                  // render the placeholder
	fmt.Printf("chances :  %d", chancesWord) // render the chances left
	fmt.Printf("\n")
	fmt.Printf("Guessed a letter :")
	printGussedLetters(entries)
	fmt.Printf("\n")
	fmt.Println("enter char or word :")
}

func getWord() string {
	resp, err := http.Get("http://example.com/")
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	var wordlist []string
	body, err := io.ReadAll(resp.Body)
	err = json.Unmarshal(body, &wordlist)
	if err != nil {
		return "elephant"
	}
	for _, word := range wordlist {
		if len(word) > 4 && len(word) < 9 {
			return word
		}
	}
	return wordlist[0]
}

func main() {
	word := getWord()
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
		printRequirdData(entries, placeholder, &chances)
		str := ""
		_, err := fmt.Scanln(&str)

		if err != nil {
			fmt.Println("*****************************")
			fmt.Println("Please enter valid input")
			continue
		}
		if str == word {
			fmt.Printf("[%v]", word)
			fmt.Printf("\n")
			fmt.Println("Win the game")
			break
		}
		// compare and update entries, placeholder and chances.
		compareWordandUpdatedGussedLetters(&entries, &placeholder, &chances, word, str)
	}
}
