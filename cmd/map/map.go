package main

import "fmt"

func testMap() {
	//var p map[int]string = make(map[int]string)
	//p := make(map[int]string)
	var p = map[int]string{}
	p[0] = "test"
	fmt.Println(p)
}

func main() {
	testMap()
}
