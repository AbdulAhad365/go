package main

import (
	"fmt"
	"maps"
)

func main() {
	//maps -> hash,object, dictionary

	//1) create the map
	m := make(map[string]string)
	n := make(map[string]int)
	// type [string], and its value it int
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
	// clear(m)c

	//			WAY 2:
	m1 := map[string]int{"price": 24, "phones": 34}
	m2 := map[string]int{"price": 24, "phones": 34}
	v, ok := m1["price"]

	fmt.Print("--------------- ")
	fmt.Print(v, ok)

	// fmt.Print(m1, m1["price"])
	//compare and & check if two maps are equal
	fmt.Println("Map1==Map2", maps.Equal(m1, m2))
}
