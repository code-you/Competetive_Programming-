package main

import "fmt"

type Stack struct {
	items []int
}

//push

func (s *Stack) push(i int) { //push will add value at the end
	s.items = append(s.items, i)
}

//pop
func (s *Stack) pop() int { // pop will remove value at the end and returns removed value
	l := len(s.items) - 1
	removedItem := s.items[l]
	s.items = s.items[:l]
	return removedItem
}
func main() {
	myStack := Stack{}
	fmt.Println(myStack)
	myStack.push(3)
	myStack.push(4)
	myStack.push(5)
	myStack.push(6)
	fmt.Println(myStack)
	fmt.Println(myStack.pop())
}
