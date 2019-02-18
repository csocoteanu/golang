package main 

import (
	"fmt"
)

const MAX_LRU_CACHE_SIZE = 20

type LRUCacheEntry struct {
	Value      int
	Prev	   *LRUCacheEntry
	Next 	   *LRUCacheEntry
}

func (l *LRUCacheEntry) Add(node *LRUCacheEntry) *LRUCacheEntry {

	if (node == nil) {

		return nil
	}

	trailingNode := l.Next

	l.Next = node
	node.Prev = l
	node.Next = trailingNode

	return node
}

func (source *LRUCacheEntry) SwapNodes(destination *LRUCacheEntry) (*LRUCacheEntry, *LRUCacheEntry) {
	
	if (source == destination) { return source, source }

	if (source.Next == destination) {

		temp := source
		source.Next = destination.Next
		destination.Next = temp

		return destination, source
	}

	prevSource 		:= source.Prev
	nextSource 	    := source.Next
	prevDestination := destination.Prev
	nextDestination := destination.Next

	// set destination links
	if (prevSource != nil) { prevSource.Next = destination }
	destination.Prev = prevSource

	destination.Next = nextSource
	if (nextSource != nil) { nextSource.Prev = destination }

	// set source links
	if (prevDestination != nil) { prevDestination.Next = source }
	source.Prev = prevDestination

	source.Next = nextDestination
	if (nextDestination != nil) { nextDestination.Prev = source }

	return destination, source
}

func (l *LRUCacheEntry) MoveNodeToHead(node *LRUCacheEntry) *LRUCacheEntry {

	// fmt.Printf("--> MoveNodeToHead: head=%d, node=%d, node == head -> %v\n", l.Value, node.Value, l == node)

	if (node == nil) { return nil }

	if (l == node) { return l }

	headingNode  := l.Prev
	trailingNode := l.Next

	movedNodePrev := node.Prev
	movedNodeNext := node.Next

	// set the refs for the target node
	if (headingNode != nil) {
		headingNode.Next = node
	}
	node.Prev = headingNode
	node.Next = trailingNode
	trailingNode.Prev = node

	// set the refs for the head of the list
	movedNodePrev.Next = l
	l.Prev = movedNodePrev
	if (movedNodeNext != nil) {
		movedNodeNext.Prev = l
	}
	l.Next = movedNodeNext

	return node
}

func (l *LRUCacheEntry) PrintPrevNodes() {

	fmt.Printf("[")

	for iterator := l; iterator != nil; iterator = iterator.Prev {

		fmt.Printf(" %d ", iterator.Value)
	}

	fmt.Printf("]\n")
}

func (l *LRUCacheEntry) PrintNextNodes() {

	fmt.Printf("[")

	for iterator := l; iterator != nil; iterator = iterator.Next {

		fmt.Printf(" %d ", iterator.Value)
	}

	fmt.Printf("]\n")
}

func testLRUCacheEntry() {

	fmt.Println("<>------- Testing LRUCacheEntry double linked list -------<>")

	head  := &LRUCacheEntry{0, nil, nil}
	node1 := &LRUCacheEntry{1, nil, nil}
	node2 := &LRUCacheEntry{2, nil, nil}
	node3 := &LRUCacheEntry{3, nil, nil}
	node4 := &LRUCacheEntry{4, nil, nil}
	tail  := node4

	head.Add(node1)
	node1.Add(node2)
	node2.Add(node3)
	node3.Add(node4)

	head.PrintNextNodes()
	tail.PrintPrevNodes()

	/*head, _ = head.SwapNodes(node1)
	head.PrintNextNodes()

	head, _ = head.SwapNodes(node2)
	head.PrintNextNodes()

	return
	head, _ = head.SwapNodes(node4)
	head.PrintNextNodes()

	return
	
	head = head.SwapNodes(node2)
	head.PrintNextNodes()

	return*/

	head, _ = head.SwapNodes(node3)
	head.PrintNextNodes()

	head, _ = head.SwapNodes(node4)
	head.PrintNextNodes()
}

/* ------------------------------------------------ */

type LRUCache struct {
	Contents 		[MAX_LRU_CACHE_SIZE]*LRUCacheEntry
	Head   			*LRUCacheEntry
	Tail   			*LRUCacheEntry
	Size			int
}

func (cache *LRUCache) Set(key, value int) {

	index := key % MAX_LRU_CACHE_SIZE

	if (cache.Contents[index] == nil) {

		node := &LRUCacheEntry{value, nil, nil}

		cache.Size++

		if (cache.Head == nil) {
			cache.Head = node
			cache.Tail = node
		} else {
			cache.Tail = cache.Tail.Add(node)
		}

		cache.Contents[index] = node
	} else {

		cache.Contents[index].Value = value
		//temp := cache.Tail
		//cache.Tail, _ = temp.SwapNodes(cache.Contents[index])
	}
}

func (cache *LRUCache) Get(key int) int {

	index := key % MAX_LRU_CACHE_SIZE
	value := cache.Contents[index]

	if (value == nil) {
		return -1
	}

	cache.Head, _ = cache.Head.SwapNodes(cache.Contents[index])

	return value.Value
}

func testLRUCache() {

	fmt.Println("<>------- Testing LRUCache -------<>")

	cache := &LRUCache{}

	cache.Set(1, 20)
	cache.Set(2, 8)
	cache.Set(0, 5)
	cache.Set(2, 19)
	
	fmt.Printf("------------->%d\n", cache.Get(2))
	fmt.Printf("------------->%d\n", cache.Get(0))

	cache.Head.PrintNextNodes()
	cache.Tail.PrintPrevNodes()
}

/* ------------------------------------------------ */

func main() {

	testLRUCacheEntry()
	testLRUCache()

	fmt.Println("Done!")
}
