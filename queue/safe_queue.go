package queue

import "sync"

// create node struct private
type node[T any] struct {
	val  T
	next *node[T]
}

// create queue struct Public
// if we use this q without the mutex ,it will work
// properly for the multiple go routines
// mutex lock the value
type SafeQueue[T any] struct {
	front *node[T]
	rear  *node[T]
	mutex sync.RWMutex // mutex
}

// function to add value to the queue
// val will always be added from rear
// will use mutex
// lock the mutex when fun is called and unloack when complete
func (q *SafeQueue[T]) Enqueue(val T) {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	//create a node
	node := node[T]{val, nil}
	//check if the q is empty
	if q.rear == nil {
		q.front = &node
		q.rear = &node
		return
	}
	q.rear.next = &node  //old rear value pointer updated to the next node
	q.rear = q.rear.next //set the new rear value ....by using the old rear value pointer
}

//func to remove value from the queue
//value will always be removed from the front

func (q *SafeQueue[T]) Dequeue() (T, bool) {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	var val T
	//check if the q is empty
	if q.front == nil {
		return val, false
	}
	val = q.front.val      //assignn the front value
	q.front = q.front.next //move the fronnt to the next value
	if q.front == nil {
		q.rear = nil
	}
	return val, true
}

// funnc to peek the value
func (q *SafeQueue[T]) Peek() (T, bool) {
	q.mutex.RLock() //read lock
	defer q.mutex.RUnlock()
	var val T
	if q.front == nil {
		return val, false
	}
	val = q.front.val
	return val, true
}
