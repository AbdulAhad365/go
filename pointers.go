package main

import "fmt"

func after_change_num(num *int) {
	*num = 5
	fmt.Println("changed value is: ", *num)
}
func change_num(num int) {
	num = 5
	fmt.Println(num)
}
func main() {
	num := 1
	fmt.Println("initial value is: ", num)
	change_num(num)
	// num will not be changed
	after_change_num(&num)
	fmt.Println(num)
	//now more code
}
