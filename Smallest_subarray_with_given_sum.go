package main

import (
	"fmt"
)

func main() {
	fmt.Printf("Smallest subarray length: %d\n",  smallest_subarray_with_given_sum([]int{2, 1, 5, 2, 3, 2}, 7))
}

func smallest_subarray_with_given_sum(array []int, sum int) int {
	windowStart, windowSum  := 0, 0 
  //I initialized smallestSubarray to a bigger size than the provided array in order to compare with the smallest 
  //array we find. If not, because the smallestSubarray was bigger than the provided array, no existing solution will exist.
  smallestSubarray := len(array) + 1
	for windowEnd, _ := range array {
		windowSum += array[windowEnd]
		//Once we find a window that is at least or equal to our sum:
		//We will substract windowStart until windowSum is smaller than sum.
		for windowSum >= sum {
			smallestSubarray = min(smallestSubarray, windowEnd - windowStart + 1)
			windowSum -= array[windowStart]
			windowStart += 1
		}
	}
  if smallestSubarray > len(array) {
    smallestSubarray = 0
  }
	return smallestSubarray
}

//Moved min to it's own function in order to clean up core solution. 
//Golang does not not have a math.min function for integers as of this time. 
func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
