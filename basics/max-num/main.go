package main

import "fmt"

func max(numbers []int) int{

	max_num := numbers[0]

	for i:=1; i<len(numbers); i++ {
		if numbers[i] > max_num {
			max_num = numbers[i]
		}
	}
	return max_num
}

func main() {

	numbers := []int{3,5,67,3,534,6,45}

	max_num := max(numbers)

	fmt.Println("The largest number is", max_num)
}

