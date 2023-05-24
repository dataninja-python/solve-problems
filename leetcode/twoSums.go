// given an array of integers and a target, return the indices of the two numbers that add up to the target
// there is only one solution
// Input: nums = [2,7,11,15], target = 9
// Output: [0,1]
// Explanation: nums[0] + nums[1] == 9, so we return [0, 1]
// 2 <= nums.length <= 10^4
// -10^9 <= nums[i] <= 10^9
// -10^9 <= target <= 10^9

// packages
package main

import (
	"fmt"
)

// the function
func twoSum(nums []int, target int) []int {
	if len(nums) == 2 {
		return []int{0, 1}
	}
	var diffMap = make(map[int]int)
	for key, value := range nums {
		if key >= target {
			continue
		}
		if index, ok := diffMap[target-value]; ok {
			return []int{index, key}
		}
		diffMap[value] = key
	}
	return nil
}

func main() {
	// initialize the input array
	inputArray := []int{3, 2, 7, 11, 59, 15, 17, 25}
	// initialize the target
	target := 17
	// call the function
	result := twoSum(inputArray, target)
	fmt.Println("result: ", result)
}
