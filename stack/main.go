package main

// declare a struct....
// it has a field fo type slice of int
type Stack struct {
	items []int
}

// make a push func to add data to the stack
// this func will append the new value to the slice
func (st *Stack) Push(value int) {
	st.items = append(st.items, value)
}

// func to delete data form the stack
// find the length of the stack
// find the element to delete
// update the stack items
func (st *Stack) Pop() int {
	l := len(st.items) - 1
	toRemove := st.items[l]
	st.items = st.items[:l]
	return toRemove
}
