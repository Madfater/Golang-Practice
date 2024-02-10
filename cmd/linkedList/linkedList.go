package main

import (
	"fmt"
)

type linkedListValue interface{}

type Node struct {
	nextNode *Node
	value    linkedListValue
}

type LinkedList struct {
	head *Node
	Size int
}

func (l *LinkedList) InsertAtHead(value linkedListValue) {
	var newNode = &Node{
		nextNode: l.head,
		value:    value,
	}

	l.head = newNode
	l.Size++
	l.PrintList()
}

func (l *LinkedList) InsertAtTail(value linkedListValue) {
	var newNode = &Node{
		nextNode: nil,
		value:    value,
	}

	var node *Node = l.head

	for node.nextNode != nil {
		node = node.nextNode
	}

	node.nextNode = newNode
	l.Size++
	l.PrintList()
}

func (l *LinkedList) InsertAtIndex(value linkedListValue, index int) {
	if l.Size < index {
		panic("over length")
	}

	if index == 0 {
		l.InsertAtHead(value)
		return
	}

	if index == l.Size {
		l.InsertAtTail(value)
		return
	}

	var prevNode *Node = nil
	var currentNode *Node = l.head
	currentIndex := 0

	for currentIndex != index {
		prevNode = currentNode
		currentNode = currentNode.nextNode
		currentIndex++
	}

	// 插入新節點
	newNode := &Node{
		nextNode: currentNode,
		value:    value,
	}
	prevNode.nextNode = newNode
	l.Size++

	l.PrintList()
}

func (l *LinkedList) DeleteAtHead() {
	if l.head == nil {
		panic("None to delete")
	}

	l.head = l.head.nextNode
	l.Size--

	l.PrintList()
}

func (l *LinkedList) DeleteAtTail() {
	if l.IsEmpty() {
		panic("None to delete")
	}

	if l.Size == 1 {
		l.DeleteAtHead()
		return
	}

	var prevNode *Node = nil
	var currentNode *Node = l.head

	for currentNode.nextNode != nil {
		prevNode = currentNode
		currentNode = currentNode.nextNode
	}

	// 刪除最後一個節點
	prevNode.nextNode = nil
	l.Size--
	l.PrintList()
}

func (l *LinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= l.Size {
		panic("index out of range")
	}

	if index == 0 {
		l.DeleteAtHead()
		return
	}

	var prevNode *Node = nil
	var currentNode *Node = l.head
	currentIndex := 0

	for currentIndex != index {
		prevNode = currentNode
		currentNode = currentNode.nextNode
		currentIndex++
	}

	// 刪除節點
	prevNode.nextNode = currentNode.nextNode
	l.Size--
	l.PrintList()
}

func (l *LinkedList) GetAtIndex(index int) linkedListValue {
	if l.Size < index {
		panic("out of range")
	}

	var node *Node = l.head
	var currentIndex int = 0

	for currentIndex != index {
		node = node.nextNode
		currentIndex++
	}

	return node.value
}

func (l *LinkedList) IsEmpty() bool {
	return l.head == nil
}

func (l *LinkedList) PrintList() {
	var node *Node = l.head

	for node != nil {
		fmt.Printf("%d ", node.value)
		node = node.nextNode
	}
	fmt.Print("\n")
}

func main() {
	var linkedList = &LinkedList{
		head: nil,
		Size: 0,
	}

	linkedList.InsertAtHead(1)
	linkedList.InsertAtHead(2)
	linkedList.InsertAtTail(3)
	linkedList.InsertAtIndex(4, 0)
	linkedList.InsertAtIndex(5, 4)
	linkedList.InsertAtIndex(6, 3)
	linkedList.DeleteAtHead()
	linkedList.DeleteAtTail()
	linkedList.DeleteAtIndex(1)
	fmt.Println("linkedlist[1] = ", linkedList.GetAtIndex(1))
	linkedList.DeleteAtTail()
	linkedList.DeleteAtTail()
	linkedList.DeleteAtTail()

	fmt.Println(linkedList.IsEmpty())
}
