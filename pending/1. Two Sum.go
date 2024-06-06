package main

// import "fmt"

// func main() {
//     nums := []int{2,7,11,15}
//     target := 9
//     fmt.Println(twoSum(nums, target))

//     nums = []int{3,2,4}
//     target = 6
//     fmt.Println(twoSum(nums, target))

//     nums = []int{3,3}
//     target = 6
//     fmt.Println(twoSum(nums, target))
// }

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