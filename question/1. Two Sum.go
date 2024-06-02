package main

import "fmt"

func main() {
	nums := []int{3,2,4,3}
	target := 6
	fmt.Println(twoSum(nums, target))
}

func twoSum(nums []int, target int) []int {
    tempMap := make(map[int]int)

    var resolution []int
    for i := range nums {
        another := target - nums[i]
 
        if _, exists := tempMap[another]; exists {
            resolution = []int{tempMap[another], i}
            break
        } 

        tempMap[nums[i]] = i 
    }
    return resolution
}