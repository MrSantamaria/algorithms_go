package main

// O(n^2) time | O(1) space

import "fmt"

func main() {
	test_array := []int{3, 5, -4, 8, 11, 1, -1, 6}
	test_sum := 10
	fmt.Println(twoNumberSumBruteForce(test_array, test_sum))
	fmt.Println(twoNumberSumHashTable(test_array, test_sum))
}

func twoNumberSumBruteForce(array []int, target int) []int {
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

// O(n) time | O(n) space
func twoNumberSumHashTable(array []int, target int) []int {
	//Implementing a hashtable here instead of two foor loops keeps this solutions at constant time. 
	//Because we are adding each value to a table the space complexity increases from O(1) to O(n)
	secondNum := map[int]bool{}
	for index := 0; index < len(array); index++ {
		pairNumber :=  target - array[index]
		if _, ok := secondNum[pairNumber]; ok {
			return []int{array[index], pairNumber}
		} else {
			secondNum[array[index]] = true
		}
	}
	return []int{}
}

func twoNumberSumPointers(array []int, target int) []int{
	pointerOne, pointerTwo := 0, len(array)-1
	resultArray := []int{}
	for len(resultArray) < 1 {
		sum := array[pointerOne] + array[pointerTwo]
		if pointerOne >= pointerTwo {
			break
		} else if sum < target {
			pointerOne += 1
		} else if sum > target {
			pointerTwo -= 1
		} else if sum == target {
			resultArray = append(resultArray, pointerOne)
			resultArray = append(resultArray, pointerTwo)
			return resultArray
		}
	}
	return resultArray
}
