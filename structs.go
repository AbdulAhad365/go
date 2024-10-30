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

// to change the status
func (o *order) change_status(stat string) {
	o.status = stat
}

// return function
func (o order) getAmount() string {
	return o.status

}

func main() {
	//instance
	// var order order=

	//order 1
	myorder := order{
		id:     "1",
		amount: 50.00,
		status: "recieved",
	}
	//order 2
	myorder2 := order{
		id:     "2",
		amount: 100.0,
		status: "not delievered",
	}
	//change the stat
	myorder.change_status("paymenet done")
	myorder.createdAt = time.Now()
	//get the value of status
	fmt.Println(myorder.getAmount())
	fmt.Println("ORDER IS ", myorder)

	myorder2.createdAt = time.Now()
	fmt.Println("order 2 is ", myorder2)
	fmt.Println(myorder.id)
}
