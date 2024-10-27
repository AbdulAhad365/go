package main

import "fmt"

func change_num(num int) {
	num = 5
	fmt.Println(num)
}
func main() {
	num := 1
	change_num(num)
	// num will not be changed
	fmt.Println(num)
	//now more code
}
