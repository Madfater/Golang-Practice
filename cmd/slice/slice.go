package main

import "fmt"

func testSlice() {

	fmt.Println("slice:")

	var slice = make([]int, 3, 3)

	//slice:=[]int{1,2,3}

	slice = append(slice, 4)

	for _, v := range slice {
		fmt.Printf("%d ", v)
	}
	fmt.Print("\n")

	slice=slice[:len(slice)-1]

	for _, v := range slice {
		fmt.Printf("%d ", v)
	}
	fmt.Print("\n")
}

func main(){
	testSlice()
}
