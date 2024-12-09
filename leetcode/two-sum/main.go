package main

import "fmt"

func twoSum(nums []int, target int) []int {
	var result []int

	for i := 0; i < len(nums)-1; i++ {
		for j := i+1; j < len(nums); j++ {
			if nums[i] + nums[j] == target {
				result = append(result, i)
				result = append(result, j)
			}
		}
	}

	return result
}


func main() {
	array := []int{3,2,4}
	fmt.Println(twoSum(array, 6))
}
