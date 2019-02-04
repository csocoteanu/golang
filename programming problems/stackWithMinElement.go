package main

import (
    "fmt"
)

const MAX_QUEUE_SIZE = 20

/* ---------------------------------------------- */

type Heap struct {
	Contents 	[MAX_QUEUE_SIZE]int
	Size 		int
}

func (h *Heap) Add(Value int) {

	h.Contents[h.Size] = Value
	h.SiftUp(h.Size)

	h.Size++
}

func (h *Heap) Swap(i1 int, i2 int) {

	temp := h.Contents[i1]
	h.Contents[i1] = h.Contents[i2]
	h.Contents[i2] = temp
}

func (h *Heap) SiftDown(startIndex int) {

	for (startIndex < h.Size) {

		minValue  := h.Contents[startIndex]
		swapIndex := -1

		if (2 * startIndex + 1 < h.Size) && (minValue > h.Contents[2 * startIndex + 1]) {

			minValue  = h.Contents[2 * startIndex + 1]
			swapIndex = 2 * startIndex + 1 
		}		

		if (2 * startIndex + 2 < h.Size) && (minValue > h.Contents[2 * startIndex + 2]) {

			minValue  = h.Contents[2 * startIndex + 2]
			swapIndex = 2 * startIndex + 2 
		}

		if (minValue == h.Contents[startIndex]) {

			break
		}

		h.Swap(startIndex, swapIndex)
		startIndex = swapIndex
	}
}

func (h *Heap) SiftUp(startIndex int) {

	for (startIndex > 0) {

		parentIndex := startIndex / 2

		if (h.Contents[startIndex] >= h.Contents[parentIndex]) {

			break	
		}

		h.Swap(parentIndex, startIndex)
		startIndex = parentIndex
	}
}

func (h *Heap) IsEmpty() bool {

	return h.Size == 0
}

func (h *Heap) PopMin() int {

	if (h.IsEmpty()) {

		return -1
	}

	minValue := h.Contents[0]

	h.Size--
	h.Swap(0, h.Size)
	h.SiftDown(0)

	return minValue
}

func (h *Heap) Init(Array []int) {

	h.Size = len(Array)
	for i := range Array {
		h.Contents[i] = Array[i]
	}


	for i := h.Size / 2; i >= 0; i-- {
		h.SiftDown(i)
	}
}

func testMinHeap() {

	fmt.Printf("Testing minheap!\n")

	h := &Heap{}

	h.Add(7)
	h.Add(3)
	h.Add(6)
	h.Add(2)
	h.Add(1)
	h.Add(5)
	h.Add(8)

	for (!h.IsEmpty()) {
		fmt.Printf("--> %d\n", h.PopMin())
	}

	fmt.Printf("Done!\n")
}

func testBuildHeap() {

	fmt.Printf("Test build heap!\n")

	h := &Heap{}

	h.Init([]int{7, 3, 6, 2, 1, 5, 8})

	for (!h.IsEmpty()) {
		fmt.Printf("--> %d\n", h.PopMin())
	}

	fmt.Printf("Done!\n")
}

/* ---------------------------------------------- */

type Stack struct {
	Contents 	 [MAX_QUEUE_SIZE]int
	Size 		 int

	MinElements	 *Heap
}

func (s *Stack) Init() {

	s.MinElements = &Heap{}
}

func (s *Stack) Push(Value int) {

	s.Contents[s.Size] = Value
	s.Size++

	s.MinElements.Add(Value)
}

func (s *Stack) GetMinValue() int {

	return s.MinElements.Contents[0]
}

func (s *Stack) Pop() int {

	s.Size--
	value := s.Contents[s.Size]

	for i := 0; i < s.MinElements.Size; i++ {
		if (s.MinElements.Contents[i] == value) {

			s.MinElements.Size--
			s.MinElements.Swap(s.MinElements.Size, i)
			s.MinElements.SiftDown(i)

			break
		}
	}

	return value
}

func (s *Stack) IsEmpty() bool {

	return s.Size == 0
}

func testStack() {

	s := &Stack{}

	s.Init()
	s.Push(7)
	s.Push(3)
	s.Push(6)
	s.Push(2)
	s.Push(1)
	s.Push(5)
	s.Push(8)
	s.Push(0)

	fmt.Printf("Min: %d\n", s.GetMinValue())

	for (!s.IsEmpty()) {

		fmt.Printf("Pop: %d Min: %d\n", s.Pop(), s.GetMinValue())	
	}

	fmt.Printf("Is Full: stack: %v heap: %v\n", s.IsEmpty(), s.MinElements.IsEmpty())
}

/* ---------------------------------------------- */

func main() {

	testMinHeap()
	testBuildHeap()
	testStack()
}