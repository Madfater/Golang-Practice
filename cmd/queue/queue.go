package main

import (
	"fmt"
)

type queueElement interface{}

type Queue struct {
	queue []queueElement
}

func (q *Queue) Push(element queueElement) {
	q.queue = append(q.queue, element)
}

func (q *Queue) Pop() queueElement {
	if q.IsEmpty() {
		panic("Queue is empty")
	}

	var value queueElement = q.queue[0]
	q.queue = q.queue[1:]

	return value
}

func (q *Queue) Size() int {
	return len(q.queue)
}

func (q *Queue) IsEmpty() bool {
	return len(q.queue) == 0
}

func main() {
	var queue = Queue{}

	queue.Push("2")
	queue.Push(3)
	fmt.Println(queue.Pop())
	fmt.Println(queue.Size())
}
