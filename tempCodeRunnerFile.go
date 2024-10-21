package main

//iterating over the data structures
import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5}
	for i := 0; i < len(nums)+1; i++ {
		fmt.Print(i)
	}
}
