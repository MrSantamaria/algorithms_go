package main

//Variable Speed N(O) this depends on how many times it has to iterate for the isSorted value to not be changed back to false. 
//Most likely being more than once N(O^2) would be my initial thought.
//Space N(1) since only pointers are being used, swap is hapenning in the same referenced array or slice. 

func main() {
  //bubbleSort(array)
}

func bubbleSort(array []int) []int {
	isSorted := false
	for !isSorted {
		isSorted = true
		for index :=0;  index < len(array) -1 ; index++ {
			firstBubble, secondBubble := array[index], array[index + 1]
			if firstBubble > secondBubble {
				array[index] = secondBubble
				array[index + 1] = firstBubble
				isSorted=false
			}
		}
				
	}
	return array
}

