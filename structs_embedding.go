package main

import (
	"fmt"
	"time"
)
// customer struct
type customer struct{
	phone	int
	email	string
}
// struct of the delivery
type delivery struct {
	id		string
	amount	float32
	status 	string
	createdAt	time.Time
	customer
}
//my main file here
int main(){
	  new_delivery:=delivery{
		id	"ahad",
		amount	3453,
		status	"new delivery done"
	  }
}