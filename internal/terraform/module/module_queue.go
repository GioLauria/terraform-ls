package module

import (
	"container/heap"
)

type moduleQueue []*module

func newModuleQueue() moduleQueue {
	q := moduleQueue{}
	heap.Init(&q)
	return q
}

var _ heap.Interface = &moduleQueue{}

func (q moduleQueue) Len() int {
	return len(q)
}

func (q moduleQueue) Less(i, j int) bool {
	leftOpen, rightOpen := 0, 0

	if q[i].HasOpenFiles() {
		leftOpen = 1
	}
	if q[j].HasOpenFiles() {
		rightOpen = 1
	}

	return leftOpen < rightOpen
}

func (q *moduleQueue) Swap(i, j int) {
	// TODO
}

func (q *moduleQueue) Pop() interface{} {
	old := *q
	n := len(old)
	item := old[n-1]
	old[n-1] = nil // avoid memory leak
	*q = old[0 : n-1]
	return item
}

func (q *moduleQueue) Push(x interface{}) {
	module := x.(*module)
	*q = append(*q, module)
}
