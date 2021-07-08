package main

import (
 "sort" 
)
// n+ O(nlogn) -> Iterating through an array and adding values would be constant, but because I decided to sort at the end the O(nlogn)
//O(n)
//To improve performance here I'd have to implement a sorted append function which could decrease it to O(n)

func main() {
  //SortedSquaredArray(array)
}

func SortedSquaredArray(array []int) []int {
	squaredSlice := make([]int, len(array))
	for element :=  range array {
		squaredSlice[element] = array[element] * array[element]
	}
	if len(squaredSlice) > 0 {
		sort.Ints(squaredSlice)
		return squaredSlice
	}
	return nil
}
