package main

import "fmt"

func main() {

	var number int
	var totalSum int

	fmt.Print("Enter a positive number N to see the sum of numbers from 1 to N : ")
	fmt.Scan(&number)

	for i:=1; i <= number; i++{
		totalSum += i
	}

	fmt.Println("The total sum from 1 to", number, "is : ", totalSum)
}
