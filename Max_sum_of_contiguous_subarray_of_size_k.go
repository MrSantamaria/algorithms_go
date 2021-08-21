package main

import (
	"fmt"
)

//Given an array of positive numbers and a positive number ‘k,’ find the maximum sum of any contiguous subarray of size ‘k’.
//Input: [2, 1, 5, 1, 3, 2], k=3 
//Output: 9
//Explanation: Subarray with maximum sum is [5, 1, 3].

func main() {
	test1 := []int{2, 1, 5, 1, 3, 2}
	fmt.Printf("Your result was %d\n", Max_sum_of_contiguous_subarray(test1, 3))
  //Time Complexity O(N) since we iterate through the array once using the Sliding Window Approach
}


func Max_sum_of_contiguous_subarray(array []int, size int) int {
	windowStart, max_sum, sum := 0, 0, 0 
	for windowEnd, _ := range array {
		sum += array[windowEnd]
    //Checks if we created the first subarray. Aftet this point the if conditions runs everytime since we are adding a new element to our window
    //BUT also removing the last index. This is what makes it better than a BruteForce at O(N*K)
		if windowEnd >= size - 1{
      //To make this slighty cleaner we could use parse this logic out to it's own function. 
			if max_sum < sum {
				max_sum = sum
			}
			sum -= array[windowStart]
			windowStart += 1
		}
	}
	return max_sum
}
