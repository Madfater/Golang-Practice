package main

import "fmt"

type stackElement interface{}

type Stack struct {
	stack []stackElement
}

func (s *Stack) Push(element stackElement){
	s.stack = append(s.stack, element)
}

func (s *Stack) Pop() stackElement {
	var value stackElement = s.stack[s.Size()-1]
	s.stack=s.stack[:s.Size()-1]
	return value
}

func (s *Stack) Top() stackElement {
	var value stackElement = s.stack[s.Size()-1]
	return value
}

func (s *Stack) Size() int {
	return len(s.stack)
}

func (s *Stack) IsEmpty() bool {
	return len(s.stack) == 0
}

func main(){
	var stack = Stack{}

	stack.Push("2")
	stack.Push(3)
	fmt.Println(stack.Pop())
	fmt.Println(stack.Size())
}