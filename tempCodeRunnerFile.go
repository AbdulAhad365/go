package main

import "fmt"

func change_num(num int) {
	num = 5
	fmt.Print(num)
}
func main() {
	num := 1
	change_num(num)
	fmt.Println(num)
}
