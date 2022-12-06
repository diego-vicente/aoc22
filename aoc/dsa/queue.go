package dsa

// A queue of integer is internally represented with a slice
type IntQueue struct {
	rel []int
}

// Create a new IntQueue by initializing the fields
func NewIntQueue() IntQueue {
	return IntQueue{[]int{}}
}

// Add a new element to the end of the queue
func (queue *IntQueue) Add(elem int) {
	queue.rel = append(queue.rel, elem)
}

// Pop the first element added to the queue
func (queue *IntQueue) Pop() int {
	result := queue.rel[0]
	queue.rel = queue.rel[1:]
	return result
}

// Return the first element added to the queue but do not remove it
func (queue IntQueue) Peek() int {
	return queue.rel[0]
}

// Return a slice of the values of an IntQueue
func (queue IntQueue) Values() []int {
	return queue.rel
}

// Return the size of the IntQueue
func (queue IntQueue) Size() int {
	return len(queue.rel)
}
