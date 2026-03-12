package main

import "fmt"

func main() {
	var number int

	fmt.Print("Enter a number to verify if is even or odd : ")
	fmt.Scanln(&number)

	if number % 2 == 0 {
		fmt.Println("The number ", number, " is even.")
	}else{
		fmt.Println("The number ", number, " is odd.")
	}
}