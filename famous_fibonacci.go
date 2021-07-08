package main

import . "fmt"

func main(){
  Println(GetNthFib_Bruteforce(55))
  Println(GetNthFib(55))
}

//Base Fibonacci Formula. Fun fact...calling 55 breaks https://play.golang.org/
//O(N^2) | O(N) Space -> Brute Force solution. 
func GetNthFib_Bruteforce(n int) int {
	if n == 1 {
		return 0
	} else if n == 2 {
		return 1
	} else {
		return GetNthFib(n-1) + GetNthFib(n-2)
	}
}

//By Using a map in this case the execution is now linear upgrading the base Fibonachi to O(n)
//The Performance is huge! Now https://play.golang.org/ does not break. 
func GetNthFib(n int) int {
  //This lines creates prepoulate K:V for the 1 and 2 edgecases.)
	return fibMemory(n, map[int]int {
		1:0,
		2:1,
	})
}

func fibMemory(n int, memory map[int]int) int {
	if val, exist := memory[n]; exist {
		return val
	}
	memory[n] = fibMemory(n-1, memory) + fibMemory(n-2, memory)
	return memory[n]
}
