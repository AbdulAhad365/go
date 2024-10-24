package main

import "fmt"

func add(a int, b int) int {
	return a + b
}
func getLanguages() (string, string, string, string, bool) {
	return "golang", "javascript", "c", "typescript", true
}
func process() func(a int) int {
	return func(a int) int {
		return a
	}
}
func main() {
	result := add(3, 5)
	fmt.Print(result)
	// way 1 to get the return function from getlnsng
	fmt.Println(getLanguages())
	// way 2:
	lan1, lang2, lang3, lang4, _ := getLanguages()
	// PRINT THR LANGUAGES
	fmt.Print(lan1, lang2, lang3, lang4)
	//ADD OTHER FUNCTIONS CALLS
	proces := process()
	fmt.Print(proces(4))

	//call the function	func ()  {

}
