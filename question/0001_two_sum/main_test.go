package main

import (
	"testing"
)

// func Test_twoSum(t *testing.T)  {
// 	nums := []int{3,2,4}
// 	target := 6
// 	result1 := twoSum(nums, target)
// 	t.Log("result1=",result1)
// }

// func Test_twoSum2(t *testing.T)  {
// 	nums := []int{3,2,4}
// 	target := 6
// 	result1 := twoSum2(nums, target)
// 	t.Log("result2=",result1)
// }

// func Test_twoSum3(t *testing.T)  {
// 	nums := []int{3,2,4}
// 	target := 6
// 	result1 := twoSum3(nums, target)
// 	t.Log("result3=",result1)
// }

func Test_twoSum4_no1(t *testing.T)  {
	nums := []int{2,7,11,15}
	target := 9
	result1 := twoSum3(nums, target)
	t.Log("result3=",result1)
}

func Test_twoSum4_no2(t *testing.T)  {
	nums := []int{3,2,4}
	target := 6
	result1 := twoSum3(nums, target)
	t.Log("result3=",result1)
}

func Test_twoSum4_no3(t *testing.T)  {
	nums := []int{3,3}
	target := 6
	result1 := twoSum3(nums, target)
	t.Log("result3=",result1)
}

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
