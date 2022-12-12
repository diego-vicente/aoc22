package dsa

import (
	"container/heap"
)

// An Item is something we manage in a priority queue.
type Item[T any] struct {
	Value    T   // The value of the item; arbitrary.
	Priority int // The priority of the item in the queue (less is higher).
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
//
// Note: please use the `container/heap` methods when pushing/popping items!
type PriorityQueue[T any] []*Item[T]

// Return the size of the PriorityQueue
func (pq PriorityQueue[T]) Len() int {
	return len(pq)
}

// Less is used to sort the PriorityQueue
func (pq PriorityQueue[T]) Less(i, j int) bool {
	return pq[i].Priority < pq[j].Priority
}

// Swap two elements of the PriorityQueue using their indices
func (pq PriorityQueue[T]) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

// Add an Item[T] to the PriorityQueue
func (pq *PriorityQueue[T]) Push(x any) {
	n := len(*pq)
	item := x.(Item[T])
	item.index = n
	*pq = append(*pq, &item)
}

// Pop the Item[T] with the lowest priority value
func (pq *PriorityQueue[T]) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// Update modifies the priority and value of an Item[T] in the queue.
func (pq *PriorityQueue[T]) Update(item *Item[T], value T, priority int) {
	item.Value = value
	item.Priority = priority
	heap.Fix(pq, item.index)
}
