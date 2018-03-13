package main

import "fmt"

func twoSum(nums []int, target int) []int {
	num_map := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		pair := target - nums[i]
		if j, ok := num_map[pair]; ok {
			return []int{j, i}
		} else {
			num_map[nums[i]] = i
		}
	}
	return nil
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(nums, target))
}
