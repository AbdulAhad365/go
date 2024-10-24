package main

import "fmt"

func add(a int, b int) int {
	return a + b
}
func getLanguages() (string, string, string, string, bool) {
	return "golang", "javascript", "c", "typescript", true
}
func main() {
	result := add(3, 5)
	fmt.Print(result)
	fmt.Println(getLanguages())
	//ADD OTHER FUNCTIONS CALLS
}
