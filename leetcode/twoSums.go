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

import "fmt"

// the function
func twoSums(nums []int, target int) []int {
	// Initialize a map to hold the result
	resultArray := make(map[int]int)
	// What is the value of the initial array
	fmt.Println("The initial value of 'resultArray': ", resultArray)
	// Check if the target is even
	// If even this means, we need to find two odd numbers or two even numbers
	// If odd this means, we need to find one odd number and one even number
	targetIsEven := getIsEven(target)
	// Initialize a variable to hold if the current number is even
	var currentIsEven bool
	fmt.Println("targetIsEven: ", targetIsEven)
	for key, value := range nums {
		currentIsEven = getIsEven(value)
		fmt.Println("key: ", key, " value: ", value, " isEven: ", currentIsEven, "\n")
	}
	return nil
}

// check if a number is odd
// Number theory says that adding an even + even = even and odd + odd = even
// 2 + 4 = 6; 4 + 6 = 10; 6 + 8 = 14;; 1 + 3 = 4; 3 + 5 = 8; 5 + 7 = 12;;
// But, even + odd = odd or odd + even = odd
func getIsEven(num int) bool {
	if num%2 == 0 {
		return true
	}
	return false
}

func main() {
	// initialize the input array
	inputArray := []int{2, 7, 11, 15}
	// initialize the target
	target := 9
	// call the function
	_ = twoSums(inputArray, target)
}
