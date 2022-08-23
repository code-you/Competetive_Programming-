package main

import "fmt"

type Queue struct {
	items []int
}

// Enqueue add elements at the end
func (q *Queue) Enqueue(i int) {
	q.items = append(q.items, i)

}

// Dequeue removes elements from first
func (q *Queue) Dequeue() int {
	removeItem := q.items[0]
	q.items = q.items[1:]
	return removeItem
}

func main() {
	myQueue := Queue{}
	myQueue.Enqueue(3)
	myQueue.Enqueue(4)
	myQueue.Enqueue(5)
	myQueue.Enqueue(6)
	myQueue.Enqueue(7)
	myQueue.Enqueue(8)
	fmt.Println(myQueue)
	myQueue.Dequeue()
	fmt.Println(myQueue)

}
