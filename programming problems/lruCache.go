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

func (l *LRUCacheEntry) Add(node *LRUCacheEntry) {

	if (node == nil) {

		return
	}

	trailingNode := l.Next

	l.Next = node
	node.Prev = l
	node.Next = trailingNode
}

func (l *LRUCacheEntry) MoveNodeToHead(node *LRUCacheEntry) *LRUCacheEntry {

	headingNode := l.Prev
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

	head  := &LRUCacheEntry{0, nil, nil}
	node1 := &LRUCacheEntry{1, nil, nil}
	node2 := &LRUCacheEntry{2, nil, nil}
	node3 := &LRUCacheEntry{3, nil, nil}
	node4 := &LRUCacheEntry{4, nil, nil}

	head.Add(node1)
	node1.Add(node2)
	node2.Add(node3)
	node3.Add(node4)

	head.PrintNextNodes()
	node4.PrintPrevNodes()

	newList := head.MoveNodeToHead(node4)
	newList.PrintNextNodes()
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
			cache.Tail.Next = node
			cache.Tail = node
		}

		cache.Contents[index] = node
	} else {

		cache.Contents[index].Value = value
	}
}

func (cache *LRUCache) Get(key int) int {

	index := key % MAX_LRU_CACHE_SIZE
	value := cache.Contents[index]

	if (value == nil) {
		return -1
	}

	return cache.Contents[index].Value
}

func testLRUCache() {

	cache := &LRUCache{}

	cache.Set(1, 20)
	cache.Set(0, 5)
	cache.Set(2, 8)
	cache.Set(2, 19)
	
	fmt.Printf("------------->%d\n", cache.Get(1))
	fmt.Printf("------------->%d\n", cache.Get(1))

	fmt.Printf("------------> %v\n", cache)

	cache.Head.PrintNextNodes()
}

/* ------------------------------------------------ */

func main() {

	testLRUCacheEntry()
	fmt.Println("Done!")
}
