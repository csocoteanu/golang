package main

import (
    "fmt"
    "time"
    "sync"
)

const READERS_COUNT = 10
const WRITERS_COUNT = 3

var (
	currentReaders int = READERS_COUNT
	currentWriters int = WRITERS_COUNT
)

func ReaderFunction(readersMutex *sync.Mutex, resourceMutex *sync.Mutex) {
	readerId := 0

	readersMutex.Lock()
	readerId = currentReaders
	fmt.Printf("Prep reading.......%d\n", readerId)

	if (currentReaders == READERS_COUNT) {
		resourceMutex.Lock()
		fmt.Println("Locking resource for readers")
	}

	currentReaders--
	readersMutex.Unlock()

	fmt.Printf("Reader#%d is reading!\n", readerId)

	readersMutex.Lock()
	if (currentReaders == 0) {
		fmt.Println("Unlocking resource for readers")
		resourceMutex.Unlock()
	}
	readersMutex.Unlock()
}

func WriterFunction(writtersMutex *sync.Mutex, resourceMutex *sync.Mutex) {
	writerId := 0

	writtersMutex.Lock()
	writerId = currentWriters
	fmt.Printf("Prep writting.......%d\n", writerId)

	if (currentWriters == WRITERS_COUNT) {
		resourceMutex.Lock()
		fmt.Println("Locking resource for writers")
	}

	currentWriters--
	writtersMutex.Unlock()

	fmt.Printf("Writer#%d is writting!\n", writerId)

	writtersMutex.Lock()
	if (currentWriters == 0) {
		fmt.Println("Unlocking resource for writers")
		resourceMutex.Unlock()
	}
	writtersMutex.Unlock()
}

func main() {
	var readersMutex 	sync.Mutex
	var writtersMutex   sync.Mutex
	var resourceMutex	sync.Mutex

	for writer := 0; writer < WRITERS_COUNT; writer++ {
		go WriterFunction(&writtersMutex, &resourceMutex)
	}

	for reader := 0; reader < READERS_COUNT; reader++ {
		go ReaderFunction(&readersMutex, &resourceMutex);
	}

	time.Sleep(3600 * time.Second)
	fmt.Println("Done!");
}
