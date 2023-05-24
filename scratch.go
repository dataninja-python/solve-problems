package main

import "fmt"

func shoes(units int, price int) int {
	return units * price
}

func main() {
	weeklySales := []int{10, 20, 30, 40, 50}
	for _, value := range weeklySales {
		fmt.Println(shoes(value, 175))
	}
}
