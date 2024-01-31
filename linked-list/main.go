package main

import "fmt"

type Node struct {
	data int
	next *Node //pointer to the next node
}

// make a linkedlist struct....
type LinkedList struct {
	head   *Node
	length int
}

// make method to add values to the LinkedList...
// this will be receiver func....
// Time Complexity= O(1)
// this will take constant amount of time every time we run the code..
func (ls *LinkedList) Prepand(n *Node) {
	temp := ls.head     // store the previous head in the temp variable...
	ls.head = n         // assign the new node to the linked list head...
	ls.head.next = temp // assign the temp address value to the head node next parameter...
	ls.length++         // as new node is add increment the length of the linked list....
}

func (ls *LinkedList) RemoveWithValue(value int) {
	if ls.length == 0 {
		// the length is 0
		return
	}
	if ls.head.data == value {
		// if the value is same as the head
		ls.head = ls.head.next
		ls.length--
		return
	}
	prevToDelete := ls.head // the previous node before the node needs ot be deleted...
	for prevToDelete.next.data != value {
		if prevToDelete.next.next == nil { // if the next node null address value
			return
		}
		//checking if the prev node next data !equal  to the value
		// if not then move the for loop to the next value...
		prevToDelete = prevToDelete.next
	}
	// if the value is equal then assign the previous node next to the value after the deleting node
	// prevToDelete.next is the node to be deleted in this case
	// using .next after above syntax we are assigning prevToDelete: the node which is after the deleted node
	prevToDelete.next = prevToDelete.next.next
	ls.length--
}

// function to show the values of the linked list
func (ls LinkedList) Peek() {
	toPrint := ls.head // the head node
	// in the for loop we will print the data
	// assign the toPrint next address
	for ls.length != 0 {
		fmt.Printf("%d ", toPrint.data)
		toPrint = toPrint.next
		ls.length--
	}
	fmt.Printf("\n")
}

func main() {
	// make new linked list
	myList := LinkedList{}
	//create new node...
	node1 := &Node{data: 3}
	node2 := &Node{data: 5}
	node3 := &Node{data: 7}
	node4 := &Node{data: 11}
	node5 := &Node{data: 13}
	myList.Prepand(node1)
	myList.Prepand(node2)
	myList.Prepand(node3)
	myList.Prepand(node4)
	myList.Prepand(node5)
	myList.Peek()
	myList.RemoveWithValue(7)
	fmt.Println("after the delete")
	myList.Peek()
}
