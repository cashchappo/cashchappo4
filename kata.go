package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, -11, -12, -13, -14, -15}
	fmt.Println(CountPositivesSumNegatives(numbers))
}

func CountPositivesSumNegatives(numbers []int) []int {
	var res []int
	a1 := len(numbers)
	var counter1 int
	var counter2 int

	for n1 := 0; n1 != a1; n1++ {
		if numbers[n1] > 0 {
			counter1 = counter1 + 1
		} else {
			counter2 = counter2 + numbers[n1]
		}
	}

	res = append(res, counter1)
	res = append(res, counter2)

	return res
}
