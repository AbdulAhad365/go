package main

//iterating over the data structures
import "fmt"

func main() {
	//array
	nums := []int{1, 2, 3, 4, 5}
	//dictionary
	num_dic := map[string]string{"ahad": "eed", "ahmad": "ewdfd"}

	//		WAY 1
	for i := 0; i < len(nums); i++ {
		fmt.Print(nums[i])
	}
	fmt.Print("   ")

	//		WAY 2
	// for i, n1 := range nums {
	// 	//i=index, n1=number
	// 	fmt.Print(i, n1)
	// 	fmt.Print("----")
	// }
	// These are some methord to make changes in the file
	for j, n2 := range num_dic {
		fmt.Print(j, n2)
		fmt.Print("-----")
	}

	//to get the unicode in the letter or strings
	for i, c := range "ahadname" {
		fmt.Print(i, c)
		fmt.Print("---chr
		")
	}
}
