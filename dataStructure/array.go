package datastructure

import "fmt"

func TestArray() {

	fmt.Print("Array:\n")

	//var arr = [5]int{1, 2, 3, 4, 5}

	//arr := [5]int{1, 2, 3, 4, 5}

	var arr = [...]int{1, 2, 3}

	for _, v := range arr {
		fmt.Printf("%d ", v)
	}

	fmt.Print("\n")
}

func Test2DArray() {

	fmt.Print("2DArray:\n")

	var arr = [...][3]int{{1, 2, 3}, {1, 2, 3}}

	for _, subArr := range arr {
		for _, v := range subArr {
			fmt.Printf("%d ",v)
		}
	}

	fmt.Print("\n")
}
