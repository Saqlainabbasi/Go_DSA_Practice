package main

import (
	"Go_DSA_Practice/queue"
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	q := queue.SafeQueue[int]{}
	//for loop to run 5 go routine
	//will use the go routine to add in parallel
	//using the wait group to wait for the go routine to finish
	wg := sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(jobNo int) {
			defer wg.Done()
			addRandomToQueue(jobNo, &q, 100)
		}(i)
	}
	wg.Wait()
	var i int
	for {
		if val, ok := q.Dequeue(); ok {
			fmt.Printf("%d - %d\n", i, val)
			i++
		} else {
			break
		}
	}
}

func addRandomToQueue(jobNo int, q *queue.SafeQueue[int], count int) {
	fmt.Printf("Job %d Started\n", jobNo)
	defer fmt.Printf("Job %d Completed\n", jobNo)
	for i := 0; i < count; i++ {
		q.Enqueue(rand.Int())
	}
}
