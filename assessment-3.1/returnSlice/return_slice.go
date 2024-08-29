package main

import (
	"fmt"
)

func printSlice(index1 int, index2 int) {
	wordArray := [8]string{"qwe", "wer", "ert", "rty", "tyu", "yui", "uio", "iop"}
	if index1 < 0 || index2 < 0 || index1 >= len(wordArray) || index2 >= len(wordArray) || index1 > index2 {
		fmt.Println("Incorrect Indexes")
		return
	}
	fmt.Println(wordArray[:index1+1])
	fmt.Println(wordArray[index1 : index2+1])
	fmt.Println(wordArray[index2:])
}

func main() {
	var index1, index2 int
	fmt.Println("Please enter two index ")
	fmt.Scanln(&index1, &index2)
	printSlice(index1, index2)
}
