package main

// O(n^2) time | O(1) space

import "fmt"

func main() {
	test_array := []int{3, 5, -4, 8, 11, 1, -1, 6}
	test_sum := 10
	fmt.Println(twoNumberSum(test_array, test_sum))
}

func twoNumberSum(array []int, target int) []int {
	for index := 0; index < len(array)-1; index++ {
		firstNum := array[index]
		//The numbers are not allowed to repeat therefor we add 1 to every index of the
		//first iteration in order to itirate through the array for a different num that
		//adds up to the target.
		for secondIndex := index + 1; secondIndex < len(array); secondIndex++ {
			secondNum := array[secondIndex]
			if firstNum+secondNum == target {
				return []int{firstNum, secondNum}
			}
		}
	}
	//The purpose of this array is the negative scenario that no numbers matched the target.
	return []int{}
}
