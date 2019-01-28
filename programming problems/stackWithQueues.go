package main

import (
    "fmt"
)

const MAX_QUEUE_SIZE = 2

type Queue struct {
	Contents 	[MAX_QUEUE_SIZE]int
	Size 	 	int
	PopIndex 	int
	PushIndex 	int
}

func (q *Queue) IsFull() bool {
	return q.Size == MAX_QUEUE_SIZE
}

func (q *Queue) IsEmpty() bool {
	return q.Size == 0
}

func (q *Queue) Push(value int) {
	if (q.IsFull()) {
		fmt.Printf("Queue is full!\n")
		return
	}

	q.Contents[q.PushIndex] = value

	q.Size++
	q.PushIndex = (q.PushIndex + 1) % MAX_QUEUE_SIZE
}

func (q *Queue) Pop() int {
	if (q.IsEmpty()) {
		fmt.Printf("Queue is empty!\n")
		return -1
	}

	value := q.Contents[q.PopIndex]

	q.Size--
	q.PopIndex = (q.PopIndex + 1) % MAX_QUEUE_SIZE

	return value
}

func (q *Queue) Top() int {
	if (q.IsEmpty()) {
		fmt.Printf("Queue is empty!\n")
		return -1
	}

	return q.Contents[q.PopIndex]	
}

/* ------------------------------------------------- */

type Stack struct {
	q1 			 *Queue
	q2 			 *Queue
	popFromQueue *Queue
}

func (s *Stack) Size() int {
	return s.q1.Size + s.q2.Size
}

func (s *Stack) IsEmpty() bool {
	return s.q1.IsEmpty() && s.q2.IsEmpty()
}

func (s *Stack) IsFull() bool {
	return s.q1.IsFull() && s.q2.IsFull()
}

func (s *Stack) Push(Value int) {
	s.popFromQueue.Pop()
}

func (s *Stack) Pop() int {

}

func (s *Stack) Top() int {
	return s.popFromQueue.Top()
}

/* ------------------------------------------------- */

func main() {
	q := &Queue{}

	for i:=0; i < 5; i++ {
		q.Push(i)
	}

	for !q.IsEmpty() {
		value := q.Pop()
		fmt.Printf("%d\n", value)
	}

	for i:=0; i < 5; i++ {
		q.Push(i)
	}

	for !q.IsEmpty() {
		value := q.Pop()
		fmt.Printf("%d\n", value)
	}

	fmt.Printf("Done!\n")
}