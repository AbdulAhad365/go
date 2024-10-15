package main
import "fmt"
func main(){
	//specific length is defined
	var arr[4]int;
	// //element add in the array
	arr[0]=32
	// //length of array
	fmt.Println(len(arr))
	// //get element
	fmt.Println(arr)
	for i:=0;i<len(arr);i++{
		fmt.Println(arr[i])
	}

	//declare methord 2
	declare_var:=[3]int{1,2,3}
	//print it
	fmt.Println(declare_var)
	//2d array)
	m:=[1][2]int{{1,2}}
	fmt.Println(m)
	
}