package main

import "fmt"

func main() {
	var arrays = []int{3, 2, 4}
	var num = 6
	twoSum(arrays, num)
	fmt.Println(twoSum(arrays, num))
}

func twoSum(nums []int, target int) []int {
	var result []int
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				result = []int{j, i}
				break
			}
		}
	}
	return result
}
