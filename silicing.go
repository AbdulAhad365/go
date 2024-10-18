package main
import "fmt"
func main(){
	var nums[]int
	// fmt.Println(nums)
	fmt.Println(nums==nil)
	// fmt.Println(len(nums))

	// part 2
	// it resize the capticty by double when more space is req
	// initially the length is 0 then there is no element in the array 
	var n1=make([]int, 3,5)
	//to change data in the array
	n1[0]=34
	//capacity is max num
	fmt.Println(n1)
	n1=append(n1,2)
	// fmt.Println(n1)
	// fmt.Println(n1)
	// fmt.Println(n1)
	// fmt.Println(n1)
	// fmt.Println(n1)
	fmt.Println(cap(n1),len(n1))



	//copy from one element to other element
	var n2=make([]int,len(n1))
	copy(n2,nums)
	fmt.Println(len(n2),len(n1))
	//now here is the second way for silicing
	fmt.Println("silicing")
	fmt.Println(len(n2))


	//more functions

	// 1) slice operator
	var n3=[]int{1,2,3}
	fmt.Println(n3[0:2]) 
}