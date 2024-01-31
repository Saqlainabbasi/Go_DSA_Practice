package main

import "fmt"

type Node struct {
	data int
	next *Node
	prev *Node
}

// declare Linked list struct
type DLinkedList struct {
	head   *Node
	length int
}

// function/method to create the list
// receiver function
func (dl *DLinkedList) AddToStart(n *Node) {
	if dl.head == nil {
		n.next = nil
		n.prev = nil
		dl.head = n
	} else {
		n.prev = nil
		dl.head.prev = n
		n.next = dl.head
		dl.head = n
	}
	fmt.Println(dl)
	dl.length++
}

// func to add value at the last.......
func (dl *DLinkedList) AddToLast(n *Node) {
	if dl.head == nil {
		n.next = nil
		n.prev = nil
		dl.head = n
	} else {
		n.prev = nil
		n.next = dl.head
		// dl.head.next =
		dl.head.prev = n
		dl.head = n
	}
	dl.length++
}

// func to add to specific index....
func (dl *DLinkedList) AddToIndex(loc int, n *Node) {
	if dl.head == nil {
		fmt.Println("Empty List")
		return
	} else {
		temp := dl.head
		for i := 0; i < loc; i++ {
			temp = temp.next
			if temp == nil {
				fmt.Println("Specific location not found")
				return
			}
		}
		// as the loop ends the temp node will be the node before the specific location
		n.next = temp.next
		n.prev = temp
		temp.next = n
		temp.next.prev = n
		dl.length++
		fmt.Println("New Node Inserted")
	}
}

// dunc to display  data
// O(n)
func (dl DLinkedList) Display() {
	toPrint := dl.head // node to print
	for dl.length != 0 {
		fmt.Printf("%d ", toPrint.data)
		toPrint = toPrint.next
		dl.length--
	}
	fmt.Printf("\n")
}

// main func
func main() {
	// make a new linked list
	myList := DLinkedList{}
	node1 := &Node{data: 1}
	node2 := &Node{data: 3}
	node3 := &Node{data: 5}
	node4 := &Node{data: 7}
	node5 := &Node{data: 11}
	node6 := &Node{data: 13}
	node7 := &Node{data: 17}
	node8 := &Node{data: 19}
	myList.AddToStart(node1)
	myList.AddToStart(node2)
	myList.AddToStart(node3)
	myList.AddToStart(node4)
	// myList.AddToStart(node5)
	myList.Display()
	myList.AddToLast(node5)
	myList.AddToLast(node6)
	myList.AddToLast(node7)
	myList.AddToLast(node8)
	fmt.Println("Add to last")
	node9 := &Node{data: 31}
	myList.AddToIndex(6, node9)
	myList.Display()

}
