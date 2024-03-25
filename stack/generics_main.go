package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type Node[T constraints.Ordered] struct {
	data T
	next *Node[T]
}

type G_Stack[T constraints.Ordered] struct {
	head   *Node[T]
	length int
}

func (s *G_Stack[T]) GS_Prepand(n *Node[T]) {
	if s.length == 0 {
		s.head = n
		s.head.next = nil
		s.length++
		return
	}
	temp := s.head

	s.head = n
	s.head.next = temp
	s.length++
}

func (s G_Stack[T]) Peek() {
	toPrint := s.head // the head node
	// in the for loop we will print the data
	// assign the toPrint next address
	for s.length != 0 {
		fmt.Print(toPrint.data)
		toPrint = toPrint.next
		s.length--
	}
	fmt.Printf("\n")
}
func main() {

	// node1 := &Node[int]{data: 3}
	// node2 := &Node[int]{data: 5}
	// node3 := &Node[int]{data: 7}
	// intlist.GS_Prepand(node1)
	// intlist.GS_Prepand(node2)
	// intlist.GS_Prepand(node3)
	// intlist.Peek()
	intList := G_Stack[int]{}
	node1 := &Node[int]{data: 3}
	node2 := &Node[int]{data: 5}
	node3 := &Node[int]{data: 7}
	intList.GS_Prepand(node1)
	intList.GS_Prepand(node2)
	intList.GS_Prepand(node3)
	intList.Peek()

	stringList := G_Stack[string]{}
	nodeA := &Node[string]{data: "A"}
	nodeB := &Node[string]{data: "B"}
	nodeC := &Node[string]{data: "C"}
	stringList.GS_Prepand(nodeA)
	stringList.GS_Prepand(nodeB)
	stringList.GS_Prepand(nodeC)
	stringList.Peek()
}
