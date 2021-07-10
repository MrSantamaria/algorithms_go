package main

func main() {
  //singleNumber([1, 2, 1, 4, 2, 3, 4, 5]
}

//O(N)
//The first iteration would be O(n)
//Since all operations for hashtables are at O(n) we are technically doing O(n) for inserting and then O(n) for looking up a value. 
//Technically this would be a O(3n) at worse scenario but would still be considered O(n)
func singleNumber(nums []int) int {
    mappy := make(map[int]bool)
    for element := 0; element < len(nums); element++ {
        if mappy[nums[element]] {
            mappy[nums[element]] = false
        } else {
            mappy[nums[element]] = true
        }
	}
    for k, v := range mappy{
        if v == true {
            return k
        }
    }
    return 0
}
