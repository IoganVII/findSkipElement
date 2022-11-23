package main

import "fmt"

func main() {

	a := []int{1, 4}
	result := Solution(a)
	fmt.Println(result)

}

func Solution(inputArray []int) int {
	minElement := inputArray[0]
	maxElement := inputArray[0]
	result := -1

	for i := 1; i < len(inputArray); i++ {
		if inputArray[i] < minElement {
			minElement = inputArray[i]
		}
		if inputArray[i] > maxElement {
			maxElement = inputArray[i]
		}
	}
	var countInput int
	for elemtnValue := minElement; elemtnValue <= maxElement; elemtnValue++ {
		countInput = 0
		for j := 0; j < len(inputArray); j++ {
			if inputArray[j] == elemtnValue {
				countInput++
				break
			}
		}
		if countInput == 0 {
			result = elemtnValue
			break
		}
	}

	return result

}
