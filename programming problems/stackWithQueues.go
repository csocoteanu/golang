package main

import (
    "fmt"
)

const MAX_QUEUE_SIZE = 20

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
	primaryQueue 	*Queue
	secondaryQueue  *Queue
}

func (s *Stack) Size() int {
	return s.primaryQueue.Size + s.secondaryQueue.Size
}

func (s *Stack) IsEmpty() bool {
	return s.primaryQueue.IsEmpty() && s.secondaryQueue.IsEmpty()
}

func (s *Stack) IsFull() bool {
	return s.primaryQueue.IsFull() && s.secondaryQueue.IsFull()
}

func (s *Stack) Push(Value int) {
	if (!s.primaryQueue.IsEmpty()) {
		s.secondaryQueue.Push(s.primaryQueue.Pop())
	}

	s.primaryQueue.Push(Value)	
}

func (s *Stack) Pop() int {
	if (s.IsEmpty()) {
		return -1	
	}
	
	poppedValue := s.primaryQueue.Pop();

	for i := 0; i < s.secondaryQueue.Size - 1; i++ {
		s.primaryQueue.Push(s.secondaryQueue.Pop())
	}

	tempQueuePointer := s.primaryQueue
	s.primaryQueue   = s.secondaryQueue
	s.secondaryQueue = tempQueuePointer

	return poppedValue
}

func (s *Stack) Top() int {
	return s.primaryQueue.Top()
}

/* ------------------------------------------------- */

func main() {
	s := &Stack{primaryQueue: &Queue{} , secondaryQueue: &Queue{} }

	s.Push(3)
	s.Push(5)
	s.Push(7)

	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.IsEmpty())
	fmt.Println(s.Pop())
	fmt.Printf("------------>  primary %v\n", s.primaryQueue)
	fmt.Printf("------------>  secondary %v\n", s.secondaryQueue)

	fmt.Printf("Done!\n")
}
