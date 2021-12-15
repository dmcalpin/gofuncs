package lists

import (
	"sort"
)

// List is a slice of anything,
// from which you can call many
// different functions
type List[T any] []T

// ForEach loops over each item
// and calls func f with the current
// element as the parameter
func (l List[T]) ForEach(f func(T)) List[T] {
	for _, elem := range l {
		f(elem)
	}

	return l
}

// Map loops over each item and calls
// func f with the current element as
// the parameter, assigning the result
// to the current of the List
func (l List[T]) Map(f func(T) T) List[T] {
	for i, elem := range l {
		l[i] = f(elem)
	}

	return l
}

// Sort is a wrapper for sort.Slice.
// It takes a "less" function as its parameter
// which is passed to sort.Slice
func (l List[T]) Sort(f func(int, int) bool) List[T] {
	sort.Slice(l, f)

	return l
}

// Len wraps the built-in len()
// Len is not chainable
func (l List[T]) Len() int {
	return len(l)
}

// Reduce calls a reducer function which
// takes as its parameters the previously
// accumulated value, and the current value
// of the current iteration. A single
// accumulated value is returned.
// This method is not chainable
func (l List[T]) Reduce(f func (prev, curr T) T, initialValue T) T {
	accumulator := initialValue
	if l.Len() == 0 {
		return accumulator
	}
	
	for _, elem := range l {
		accumulator = f(accumulator, elem)
	}

	return accumulator
}
