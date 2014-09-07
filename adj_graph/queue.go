package main

type Queue []interface{}

func (q *Queue) enqueue(elem interface{}) {
	*q = append(*q, elem)
}

func (q *Queue) isEmpty() bool {
	if len(*q) == 0 {
		return true
	}
	return false
}

func (q *Queue) dequeue() interface{} {
	if q.isEmpty() {
		panic("Cannot dequeue because queue is empty.")
	}
	elem := (*q)[0]
	copy((*q)[0:], (*q)[1:])
	(*q)[len(*q)-1] = nil
	(*q) = (*q)[:len(*q)-1]
	return elem
}
