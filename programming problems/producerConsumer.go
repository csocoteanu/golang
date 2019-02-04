package main 

import (
	"fmt"
	"time"
	"sync"
	"math/rand"
)

const PRODUCER_SLEEP_TIME = 2 * time.Second
const CONSUMER_SLEEP_TIME = 1 * time.Second
const PRODUCER_COUNT = 2
const CONSUMER_COUNT = 5

const MAX_BUFFER_SIZE = 100

func ProducerFunction(ProducerID int, channel chan int, done chan bool) {

	Value := 0

	for {

		time.Sleep(PRODUCER_SLEEP_TIME)

		seed := rand.NewSource(time.Now().UnixNano() + int64(ProducerID))
		random := rand.New(seed)
    	Value = random.Intn(1000)

		fmt.Printf("Producer #%d....%d\n", ProducerID, Value)		
		channel <- Value
	}

	done <- true
}

func ConsumerFunction(ConsumerID int, channel chan int, done chan bool) {

	Value := 0

	for {

		Value = <- channel

		fmt.Printf("Consumer #%d....%d\n", ConsumerID, Value)
		time.Sleep(CONSUMER_SLEEP_TIME)
	}

	done <- true
}

func ProducerConsumerWithChannel() {

	channel := make(chan int)
	done := make(chan bool)

	for i := 0; i < CONSUMER_COUNT; i++ {

		go ConsumerFunction(i, channel, done)
	}

	for i := 0; i < PRODUCER_COUNT; i++ {

		go ProducerFunction(i, channel, done)
	}

	for i := 0; i < CONSUMER_COUNT + PRODUCER_COUNT; i++ {
		<- done
	}
}

type RWBuffer struct {
	Lock      sync.Mutex
	Value 	  [MAX_BUFFER_SIZE]int
	Size      int
}

func (b *RWBuffer) Add(Value int) {

	b.Lock.Lock()
	defer b.Lock.Unlock()

	if (b.Size == MAX_BUFFER_SIZE) {
		return
	}

	b.Value[b.Size] = Value
	b.Size++
}

func (b *RWBuffer) Get() int {

	b.Lock.Lock()
	defer b.Lock.Unlock()

	if (b.Size == 0) {
		return -1
	}

	b.Size--
	return b.Value[b.Size]
}

func (b *RWBuffer) IsEmpty() bool {

	b.Lock.Lock()
	defer b.Lock.Unlock()

	return b.Size == 0
}

func (b *RWBuffer) IsFull() bool {

	b.Lock.Lock()
	defer b.Lock.Unlock()

	return b.Size == MAX_BUFFER_SIZE
}

func ProduceUntilFull(ProducerID int, b *RWBuffer, produceChannel chan bool, consumeChannel chan bool, done chan bool) {

	for {
		produceCount := 0

		<- produceChannel
		fmt.Printf("Producer #%d started....\n", ProducerID)

		for !b.IsFull() {
			b.Add(0)
			produceCount++
		}

		time.Sleep(PRODUCER_SLEEP_TIME)
		fmt.Printf("Producer #%d completed %d....\n", ProducerID, produceCount)

		consumeChannel <- true
	}

	done <- true
}

func ConsumeUntilEmpty(ConsumerID int, b *RWBuffer, consumeChannel chan bool, produceChannel chan bool, done chan bool) {

	for {
		consumeCount := 0

		<- consumeChannel
		fmt.Printf("Consumer #%d started....\n", ConsumerID)

		for !b.IsEmpty() {
			b.Get()
			consumeCount++
		}

		time.Sleep(CONSUMER_SLEEP_TIME)
		fmt.Printf("Consumer #%d completed %d....\n", ConsumerID, consumeCount)

		produceChannel <- true
	}

	done <- true
}

func ProducerConsumerWithBuffer() {

	buffer 		    := &RWBuffer{}
	produceChannel  := make(chan bool, CONSUMER_COUNT)
	consumerChannel := make(chan bool, CONSUMER_COUNT)
	done		    := make(chan bool)

	for i := 0; i < CONSUMER_COUNT; i++ {
		go ProduceUntilFull(i, buffer, produceChannel, consumerChannel, done)
	}

	for i := 0; i < CONSUMER_COUNT; i++ {
		go ConsumeUntilEmpty(i, buffer, consumerChannel, produceChannel, done)
	}

	for i:=0; i < CONSUMER_COUNT; i++ {
		produceChannel <- true
	}

	<- done
}

func main() {

	ProducerConsumerWithBuffer()
}
