package main
import "fmt"
func main() {
	//while loop fashio
	i:=1;
	for i<=5{
		fmt.Println(i)
		i+=	1
	}
	//			2)		for infinite loop
	// for {
	// 	fmt.Println(i);
	// }
	fmt.Println("-------------")
	//			3)		for loop
	for j:=0;j<3;j++{
		if (j==1){
			continue
		}
		fmt.Println(j)
	}
	//			4)		for loop with range
	fmt.Println("-------------")
	for k:=range "golang"{
		fmt.Println(k)
	}
}