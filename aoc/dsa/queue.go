package dsa

// A queue is internally represented with a slice
type Queue[T any] struct {
	rel []T
}

// Create a new Queue by initializing the fields
func NewQueue[T any]() Queue[T] {
	return Queue[T]{[]T{}}
}

// Add a new element to the end of the Queue
func (queue *Queue[T]) Add(elem T) {
	queue.rel = append(queue.rel, elem)
}

// Pop the first element added to the Queue
func (queue *Queue[T]) Pop() T {
	result := queue.rel[0]
	queue.rel = queue.rel[1:]
	return result
}

// Return the first element added to the Queue but do not remove it
func (queue Queue[T]) Peek() T {
	return queue.rel[0]
}

// Return a slice of the values of an Queue
func (queue Queue[T]) Values() []T {
	return queue.rel
}

// Return the size of the Queue
func (queue Queue[T]) Size() int {
	return len(queue.rel)
}
