package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	/*
		sequential()
		parallel()
	*/
	/*
		var c chan int = make(chan int)
			go add(0, 4, c)
			go add(0, 3, c)
			go add(0, 2, c)
			go add(0, 1, c)
			var result int
			for i := 0; i < 4; i++ {
				result = <-c
				fmt.Println(result)
			}
	*/
	/*
		var c1 chan int = make(chan int)
		var c2 chan int = make(chan int)
		go add(0, 2, c2)
		go add(0, 1, c1)
		for i := 0; i < 2; i++ {
			// Select the case that runs first
			select {
			case x1 := <-c1:
				fmt.Println(x1)
			case x2 := <-c2:
				fmt.Println(x2)
			}
		}
	*/

	/*
		var c chan bool = make(chan bool)
		go receive(c)
		send(c)
		fmt.Println("done")
	*/

	//buffered()
	//raceCondition()
	//noRaceCondition()
	noRaceCondition2()
}

func run() {
	fmt.Println("started")
	time.Sleep(500 * time.Millisecond)
	fmt.Println("finished")
}

func sequential() {
	fmt.Println("sequential")
	run()
	run()
	run()
}

func parallel() {
	fmt.Println("parallel")
	go run()
	go run()
	go run()
	// should use channels to block here
}

func add(a int, b int, c chan int) {
	for i := 0; i < b; i++ {
		time.Sleep(100 * time.Millisecond)
	}
	c <- a + b
}

func send(c chan bool) {
	fmt.Println("Sending")
	c <- true
}

func receive(c chan bool) {
	fmt.Println("Receiving")
	<-c
}

func buffered() {
	// buffered channel
	var ch chan bool = make(chan bool, 2)
	// More than 2 blocks for a receive
	ch <- true
	ch <- false
	<-ch
	fmt.Println(1)
	<-ch
	fmt.Println(2)
}

type Count struct {
	value int
}

func increment(count *Count) {
	count.value++
	fmt.Println(count.value)
}

func raceCondition() {
	var count Count = Count{0}
	for i := 0; i < 1000; i++ {
		go increment(&count)
	}
	// wait some time for the routines to finish
	time.Sleep(2 * time.Second)
	// Got 993, which is less than 1000
	fmt.Printf("Final value = %v\n", count.value)
}

// Put a lock to eliminate race condition
type CountMutex struct {
	value int
	lock  sync.Mutex
}

// The channel is used to wait for the threads to finish.
// Will use a send only channel
func incrementSafe(count *CountMutex, ch chan<- bool) {
	// Get a lock, to eliminate the race condition
	count.lock.Lock()
	// Defer releasing the lock, this guarantess that it will be release
	// even when an error occurs
	defer count.lock.Unlock()
	count.value++
	ch <- true
}

// Count again, without race condition
func noRaceCondition() {
	var count CountMutex = CountMutex{0, sync.Mutex{}}
	var ch chan bool = make(chan bool)
	nthreads := 1000
	for i := 0; i < nthreads; i++ {
		go incrementSafe(&count, ch)
	}

	// wait for all threads to finish
	for i := 0; i < nthreads; i++ {
		<-ch
	}

	fmt.Printf("Final value = %v\n", count.value)
}

// Instead of using channels, we can use wait groups
// It has the convenience of not needing a type.
// The type bool was used previously for the channels.

func incrementSafe2(count *CountMutex, waitGroup *sync.WaitGroup) {
	count.lock.Lock()
	// Defer releasing the lock, this guarantess that it will be release
	// even when an error occurs
	defer count.lock.Unlock()
	count.value++
	waitGroup.Done()
}

func noRaceCondition2() {
	var count CountMutex = CountMutex{0, sync.Mutex{}}
	nthreads := 1000
	var waitGroup sync.WaitGroup = sync.WaitGroup{}
	waitGroup.Add(nthreads)
	for i := 0; i < nthreads; i++ {
		go incrementSafe2(&count, &waitGroup)
	}

	// wait for all threads to finish
	waitGroup.Wait()
	fmt.Printf("Final value = %v\n", count.value)
}
