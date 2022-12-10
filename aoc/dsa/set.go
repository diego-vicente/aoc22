package dsa

// A set is represented by a map of its values
type Set[T comparable] struct {
	rel map[T]bool
}

// Create a new Set and initialize all fields
func NewSet[T comparable]() Set[T] {
	return Set[T]{map[T]bool{}}
}

// Create a new Set from a slice of values
func NewSetFrom[T comparable](values []T) Set[T] {
	set := Set[T]{map[T]bool{}}

	for _, elem := range values {
		set.Add(elem)
	}

	return set
}

// Add a new element to the Set
func (set Set[T]) Add(elem T) {
	set.rel[elem] = true
}

// Remove an element from an Set
func (set Set[T]) Remove(elem T) {
	delete(set.rel, elem)
}

// Return a slice of the values fo an Set
func (set Set[T]) Values() []T {
	var values []T

	for value := range set.rel {
		values = append(values, value)
	}

	return values
}

// Return the number of elements in the Set
func (set Set[T]) Size() int {
	return len(set.rel)
}

// Return true if the Set contains the element
func (set Set[T]) Contains(elem T) bool {
	_, ok := set.rel[elem]
	return ok
}

// Return the union of two given Sets
func (set Set[T]) Union(other Set[T]) Set[T] {
	union := NewSet[T]()

	for elem := range set.rel {
		union.Add(elem)
	}

	for elem := range other.rel {
		union.Add(elem)
	}

	return union
}

// Return the intersection of two given Sets
func (set Set[T]) Intersection(other Set[T]) Set[T] {
	intersect := NewSet[T]()

	for elem := range set.rel {
		if other.Contains(elem) {
			intersect.Add(elem)
		}
	}

	return intersect
}
