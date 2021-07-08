package main
//O(n^2) This iterates through the array/slice multiple times each time it gets to a new element.
func main() {
  //insertionSort(array)
}

func insertionSort(array []int) []int {
	for element := range array {
    //In the first Iteration it skips the execution because it is only looking at one element. 
    //On the second Iteration it checks array[1] > array[0] then does logic. 
    //On the third < Iterations the logic will check every single value to make sure it is not bigger than a previous element. 
		for elementComparer := element; elementComparer > 0 && array[elementComparer] < array[elementComparer -1]; elementComparer-- {
			array[elementComparer], array[elementComparer -1] = array[elementComparer -1], array[elementComparer]
			log := Sprintf("%d Compared to %d\n", array[elementComparer], array[elementComparer-1])
			
		}
	}
	return array
}
