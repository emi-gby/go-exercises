package main

import "fmt"	


func reverse(word string) string {
	runes := []rune(word)
	runesReversed := []rune{}

	fmt.Println(runes)

	for i := (len(runes)-1); i>=0; i-- {
		runesReversed = append(runesReversed, runes[i])
	}
	wordReversed := string(runesReversed)
	return wordReversed
}

func main() {

	var word string
	
	fmt.Print("Enter a word to reverse : ")
	fmt.Scanln(&word)

	wordReversed := reverse(word)

	fmt.Println("Result : ", wordReversed)
}
