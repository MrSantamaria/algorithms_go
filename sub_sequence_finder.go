package main

import (
  "fmt"
)
//O(N) time-ish since the sequence could be found before the whole array is traversed. 
//O(1) space since we are only creating pointers. 

func main() {
  sequence := []int{2, 20, 11, 9}
  array := []int{1, 2, 6, 20, 40, 80, 11, 100, 9, 12}
  fmt.Println(IsValidSubsequence(array, sequence))
}

func IsValidSubsequence(array []int, sequence []int) bool {
	arrayIndex := 0
	sequenceIndex := 0
	for sequenceIndex < len(sequence) && arrayIndex < len(array) {
		if array[arrayIndex] == sequence[sequenceIndex] {
			sequenceIndex += 1
		}
		arrayIndex += 1
	}
	return sequenceIndex == len(sequence)
}
