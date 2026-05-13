package main

import "fmt"

func main() {

	var number, timesNum int
	
	fmt.Print("Enter the number to see its multiplication table :")
	fmt.Scan(&number)
	fmt.Print("Enter the number for how far you want to go : ")
	fmt.Scan(&timesNum)

	for i := 0; i <= timesNum; i++ {
		fmt.Println(number, " * ", i, " = ", number*i)
	}

}
