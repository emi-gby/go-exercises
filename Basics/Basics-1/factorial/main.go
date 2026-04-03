package main

import "fmt"

func factorial(number int) int {
	if number == 0{
		return 0
	}

	total := 1
	for i:=1; i<=number; i++ {
		total = total * i
	}
	return total
}



func main(){

	var number int

	fmt.Print("Enter a positive integer: ")
	fmt.Scanln(&number)

	total := factorial(number)
	fmt.Println("The factorial of", number, "is :", total)

}
