package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"golang.org/x/tour/tree"
)

func walkTree (tree *tree.Tree, channel chan int) {
	if (tree.Left != nil) {
		walkTree(tree.Left, channel)
	}

	channel <- tree.Value

	if (tree.Right != nil) {
		walkTree(tree.Right, channel)
	}
}

func isSameTree (tree1 *tree.Tree, tree2 *tree.Tree) bool {
	c1 := make(chan int, 10)
	c2 := make(chan int, 10)

	go walkTree(tree1, c1)
	go walkTree(tree2, c2)

	for i := 0; i < 10; i++ {
        value1 := <-c1
        value2 := <-c2

    	fmt.Printf("Comparing: %d with %d\n", value1, value2)

        if (value1 != value2) {
        	return false
        }
    }

    return true
}

func printHttpResponse(url string) {
	response := make(chan *http.Response, 1)
	errors := make(chan *error)

	go func() {
		resp, err := http.Get(url)
		if err != nil {
			errors <- &err
		}
		response <- resp
	}()
	for {
		select {
		case r := <-response:
			fmt.Printf("Got response: %s", r.Body)
			return
		case err := <-errors:
			log.Fatal(*err)
		/*case <-time.After(200 * time.Millisecond):
			fmt.Printf("Timed out!")
			return*/
		}
	}
}

func sum(a []int, c chan int) {
	sum := 0

	for _, v := range a {
		sum += v
	}

	fmt.Printf("Sending sum: %d\n", sum)
	c <- sum // send sum to c
	fmt.Printf("We have send sum: %d\n", sum)
}

func testSum() {
	a := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)
	x, y := <-c, 0 // <-c // receive from c

	fmt.Println(x, y, x+y)

	time.Sleep(5 * time.Second)
}

func fibonacci(n int, c chan int) {
    x, y := 0, 1
    for i := 0; i < n; i++ {
        c <- x
        x, y = y, x+y
    }
    close(c)
}

func testFibonacci() {
    c := make(chan int, 10)

    go fibonacci(cap(c), c)
    for i := range c {
        fmt.Println(i)
    }
}

func main() {
	// printHttpResponse("http://matt.aimonetti.net/")
	
	fmt.Println(isSameTree(tree.New(1), tree.New(1)))
	fmt.Println(isSameTree(tree.New(1), tree.New(2)))
}
