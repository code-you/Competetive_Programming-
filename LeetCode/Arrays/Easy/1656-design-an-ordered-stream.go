//Ques:- https://leetcode.com/problems/design-an-ordered-stream/
package main

type OrderedStream struct {
	id     int
	stream map[int]string
}

func Constructor(n int) OrderedStream {
	return OrderedStream{
		id:     1,
		stream: make(map[int]string, 0),
	}
}

func (this *OrderedStream) Insert(idKey int, value string) []string {
	this.stream[idKey] = value
	var result []string
	if this.id == idKey {
		for {
			result = append(result, this.stream[this.id])
			this.id += 1
			if _, ok := this.stream[this.id]; !ok {
				break
			}
		}
	}
	return result
}
