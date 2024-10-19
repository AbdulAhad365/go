package main

import (
	"fmt"
)

func main() {
	//maps -> hash,object, dictionary

	//1) create the map
	m := make(map[string]string)
	n := make(map[string]int)
	//setting the element
	m["name"] = "golang"
	m["area"] = "backend"

	n["val"] = 4
	//key value not exit return the zero value
	//for string --> "" ,int-->0, boolean-->false
	fmt.Println(m["name"], m["area"], m["w"], n["val"], n["m"])

	//len
	fmt.Print(len(m))

	//delete for deleting
	delete(m, "area")
	fmt.Print(m)

	//to empty everything use clear
	// clear(m)
}
