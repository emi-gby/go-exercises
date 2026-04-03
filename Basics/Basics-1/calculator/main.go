package main

import "fmt"

func calculate(num1, num2 float64, op string) (float64, error){
	switch op {
		case "+" : return num1 + num2, nil
		case "-" : return num1 - num2, nil
		case "*" : return num1 * num2, nil
		case "/" : {
			if num2 == 0{
				return 0, fmt.Errorf("Division by zero")
			}
			return num1 / num2, nil
		}
	}
	return 0, fmt.Errorf("Operation invalid")
}

func main() {
	var num1, num2 float64
	var op string

	fmt.Println("Enter 2 integer numbers and the operation(+,-,*,/): ")
	fmt.Scanln(&num1)
	fmt.Scanln(&num2)
	fmt.Scanln(&op)

	result, error := calculate(num1, num2, op)

	if error == nil {
		fmt.Println("Result = ",result)
	}else{
		fmt.Println("Error :", error)
	}
}
