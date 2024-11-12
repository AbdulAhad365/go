package main

import (
	"fmt"
	"time"
)

// import "ti/me"
func who_am_i(i interface{}) {
	switch i.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	}

}

// main function of the switch statement
func main() {
	i := 3
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("other")
	}

	//---------------------------

	// multiple conditon switch
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("its weekand")
	default:
		fmt.Println("its a work day")
	}

	//type switch
	whoami := func(i interface{}) {
		switch i.(type) {
		case int:
			fmt.Println("int value")
		case string:
			fmt.Println("its a string")
		case bool:
			fmt.Println("its a boolesn")
		default:
			fmt.Println("Other")
		}

	}

	whoami("astrhad")
	who_am_i("value")
}
