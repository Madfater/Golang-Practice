package main

import (
	"encoding/json"
	"fmt"
)

//一定要大寫才會被json encoding到
type Person struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Gender bool   `json:"gender"`
}

func testStruct() {
	var p = &Person{
		Name:   "CCcat",
		Age:    21,
		Gender: true,
	}

	jsons, errs := json.Marshal(p)
	if errs != nil {
		fmt.Println("errs", errs)
	}
	fmt.Println("json data :", string(jsons))
}

func main() {
	testStruct()
}
