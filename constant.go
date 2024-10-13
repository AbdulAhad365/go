package main
import "fmt"
func main(){
	const name="ahad"
// it is not possible to reassign
	// name="ahmad"
	fmt.Println(name)
	//to declare multiple constants

	const (
		port=5000
		nam_e="ahad"
	)
	fmt.Println(port,nam_e)

}