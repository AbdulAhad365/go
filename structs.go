package main

import (
	"fmt"
	"time"
)

// struct order
// struct : to group multiple things at one place
type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time //nanosecond precision
}

func main() {
	//instance
	// var order order=
	myorder := order{
		id:     "1",
		amount: 50.00,
		status: "recieved",
	}
	myorder.createdAt = time.Now()
	fmt.Println("ORDER IS ", myorder)
	fmt.Println(myorder.id)
}
