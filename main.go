package main

import (
	// "Go_DSA_Practice/queue"

	question "Go_DSA_Practice/interview-questions"
	"fmt"
)

func main() {
	// q := queue.SafeQueue[int]{}
	// //for loop to run 5 go routine
	// //will use the go routine to add in parallel
	// //using the wait group to wait for the go routine to finish
	// wg := sync.WaitGroup{}
	// for i := 0; i < 5; i++ {
	// 	wg.Add(1)
	// 	go func(jobNo int) {
	// 		defer wg.Done()
	// 		addRandomToQueue(jobNo, &q, 100)
	// 	}(i)
	// }
	// wg.Wait()
	// var i int
	// for {
	// 	if val, ok := q.Dequeue(); ok {
	// 		fmt.Printf("%d - %d\n", i, val)
	// 		i++
	// 	} else {
	// 		break
	// 	}
	// }

	// interview questions practice
	fac := question.Factorial(5)
	fab := question.Fibonacci(10)
	fmt.Println(fac)
	fmt.Println(fab)
}

// func addRandomToQueue(jobNo int, q *queue.SafeQueue[int], count int) {
// 	fmt.Printf("Job %d Started\n", jobNo)
// 	defer fmt.Printf("Job %d Completed\n", jobNo)
// 	for i := 0; i < count; i++ {
// 		q.Enqueue(rand.Int())
// 	}
// }
