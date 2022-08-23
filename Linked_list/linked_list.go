package main

import "fmt"

type node struct {
	data int
	next *node
}

type linkedList struct {
	head   *node
	length int
}

func (l *linkedList) prepend(n *node) {
	second := l.head
	l.head = n
	l.head.next = second
	l.length++
}

func (l *linkedList) printListData() {
	toPrint := l.head
	for l.length != 0 {
		fmt.Printf("%d ", toPrint.data)
		toPrint = toPrint.next
		l.length--
	}
	fmt.Println("\n")
}

func (l *linkedList) deleteWithValue(value int) {
	if l.length == 0 { // for empty list case
		return
	}
	if l.head.data == value { // if deleting value is pointing to head
		l.head = l.head.next
		l.length--

		return
	}
	previousToDelete := l.head
	for previousToDelete.next.data != value {
		if previousToDelete.next.next == nil { // we reached to the end
			return
		}
		previousToDelete = previousToDelete.next

	}
	previousToDelete.next = previousToDelete.next.next
	l.length--
}

func main() {
	myList := linkedList{}
	node1 := &node{data: 40}
	node2 := &node{data: 45}
	node3 := &node{data: 48}
	node4 := &node{data: 4}
	node5 := &node{data: 9}

	myList.prepend(node1)
	myList.prepend(node2)
	myList.prepend(node3)
	myList.prepend(node4)
	myList.prepend(node5)

	//myList.printListData()
	myList.deleteWithValue(100) // value does not exist case
	myList.deleteWithValue(9)   // deleting head element case
	myList.printListData()
	emptyList := linkedList{}
	emptyList.deleteWithValue(10) // deleting with empty list

}
