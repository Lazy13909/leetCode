package main

import "fmt"

func main() {
	var arrays = []int{2, 5, 5, 11}
	var num = 10
	twoSum(arrays, num)
	fmt.Println(twoSum(arrays, num))
}

func twoSum(nums []int, target int) []int {
	var result []int
	var sum int
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			sum = nums[i] + nums[j]
			if sum == target {
				return []int{j, i}
			}
		}
	}
	return result
}
