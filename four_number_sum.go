package main

func main() {
  FourNumberSumBruteforce([1, 2, 3, 4, 5, -5, 6, -6], 5)
}

//This is a bruteforce but the complexity is the worse. 
//N(O^4) for speed. 
//N(O) for space
func FourNumberSumBruteforce(array []int, target int) [][]int {
	responseArray := [][]int{}
	for elementUno := 0;  elementUno < len(array); elementUno++ {
		for elementDos := elementUno + 1;  elementDos < len(array); elementDos++ {
			for three := elementDos + 1; three < len(array); three++ {
				for four := three + 1; four < len(array); four++ {
					if array[elementUno] + array[elementDos] + array[three] + array[four] == target {
						responseArray = append(responseArray, []int{array[elementUno], array[elementDos], array[three], array[four]})
					}
				}
			}
		}
	}
	return responseArray
}

  
