package util

import (
	"container/heap"
)

// An Item is a container for the actual values and their priority within the queue.
// The value can be any type that satisfies the comparable constraint.
type Item[T comparable] struct {
	Value    T   // The Value of the item; arbitrary.
	Priority int // The Priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Item(s).
// Based on example included in the docs for "container/heap" modified to support storing a generic value.
type PriorityQueue[T comparable] []*Item[T]

func (pq PriorityQueue[T]) Len() int { return len(pq) }

func (pq PriorityQueue[T]) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, Priority so we use greater than here.
	return pq[i].Priority < pq[j].Priority
}

func (pq PriorityQueue[T]) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue[T]) Push(x any) {
	n := len(*pq)
	item := x.(*Item[T])
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue[T]) PushValue(v T, priority int) *Item[T] {
	item := &Item[T]{
		Priority: priority,
		Value:    v,
	}
	heap.Push(pq, item)
	return item
}

func (pq *PriorityQueue[T]) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // don't stop the GC from reclaiming the item eventually
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue[T]) PopValue() (T, int) {
	item := heap.Pop(pq).(*Item[T])
	return item.Value, item.Priority
}

// UpdatePriority modifies the Priority an item in the queue.
func (pq *PriorityQueue[T]) UpdatePriority(value T, priority int) {
	for _, item := range *pq {
		if item.Value == value {
			item.Priority = priority
			heap.Fix(pq, item.index)
			return
		}
	}
}
