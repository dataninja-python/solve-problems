// Find the largest product of any two adjacent numbers in the given array.
// Example: []int{3, 6, -2, -5, 7, 3} => 21

package main

import "fmt"

func solution(inputArray []int) int {
	arrLen := len(inputArray)
	endIndex := arrLen - 1
	var current int
	var next int
	var largestProduct int
	var currentProduct int
	var isPositive bool
	if arrLen == 2 {
		return product(inputArray[0], inputArray[1])
	}
	for i := 0; i < arrLen; i++ {
		// As long as the next number is not the last then do get product using normal algorithm
		if i+1 != endIndex {
			// If this is our first time through the loop then do all calculations as normal
			// because we already handled the case where the array length is 2
			if i == 0 {
				current = inputArray[i]
				next = inputArray[i+1]
				largestProduct = product(current, next)
				if largestProduct > 0 {
					isPositive = true
				} else {
					isPositive = false
				}
			} else {
				next = inputArray[i+1]
				if isPositive && next <= 0 {
					continue
				}
				current = inputArray[i]
				currentProduct = product(current, next)
				largestProduct = getLargest(largestProduct, currentProduct)
			}
		} else {
			// Use special algorithm for when the next number is the last
			next = inputArray[i+1]
			if isPositive && next <= 0 {
				break
			}
			current = inputArray[i]
			currentProduct = product(current, next)
			largestProduct = getLargest(largestProduct, currentProduct)
			break
		}
	}
	return largestProduct
}

func product(a int, b int) int {
	return a * b
}

func getLargest(a int, b int) int {
	switch {
	case b < a:
		return a // a is larger
	case b > a:
		return b // b is larger
	default:
		return a // a and b are equal
	}
}

func main() {
	// a := []int{3, 6, -2, -5, 7, 3}
	a := []int{3, 6, -2, -5, 7, 3, 3, 0}
	x := solution(a)
	fmt.Println(x)
}
