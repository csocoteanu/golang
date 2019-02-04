package main 

import (
	"fmt"
)

type LinkedList struct {
	Value int
	Next  *LinkedList
}

func AddFront(list **LinkedList, value int) {

	newNode := &LinkedList{value, *list}

	*list = newNode
}

func AddBack(list **LinkedList, value int) {

	if (*list == nil) {
		*list = &LinkedList{value, nil}
		return
	}

	listPointer := *list

	for ; listPointer.Next != nil; listPointer = listPointer.Next { }

	listPointer.Next = &LinkedList{value, nil}
}

func ReverseList(list **LinkedList) {
	
	var refNode *LinkedList = nil
	startNode := *list
	iterator := *list

	for ; iterator != nil && iterator.Next != nil ; {

		fmt.Printf("before swap: it.v=%d startNode.v=%d refNode=%v, it.Next = %v\n", iterator.Value, startNode.Value, refNode, iterator.Next)
		iterator = iterator.Next.Next

		tempRef := startNode.Next
		startNode.Next = refNode
		tempRef.Next = startNode

		fmt.Printf("after swap: it.v=%v startNode.v=%v refNode=%v \n", iterator, startNode, refNode)
	
		if (iterator != nil) {
			startNode = iterator
		}

		refNode = tempRef
	}

	*list = startNode
}

func PrintList(list *LinkedList) {

	fmt.Printf("[")

	for ; list != nil; list = list.Next {

		fmt.Printf(" %d", list.Value)
	}

	fmt.Printf(" ]\n")
}

func main() {

	var list *LinkedList = nil

	for i := 0; i < 2; i++ {
		AddBack(&list, i)
	}

	PrintList(list)

	ReverseListV1(&list)

	PrintList(list)
}