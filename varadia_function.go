package main

import "fmt"

// function for finding the sum
func sum(nums ...int) int {
	// func sum(nums ...interface{}) int {
	//we get a slice from it
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

// bring change
// Main function
func main() {
	// fmt.Println(1, 2, 3, 4, "hello there ")
	total := sum(1, 2, 3, 4)
	fmt.Print(total)
	fmt.Print("--")
	//way 2 to pass the slice values directly
	nums := []int{1, 2, 3, 4, 5, 6, 6, 7, 8, 9, 10}
	t := sum(nums...)
	fmt.Println(t)
}
