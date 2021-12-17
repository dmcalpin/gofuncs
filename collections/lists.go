package collections

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
	tmp := List[T]{}
	for _, elem := range l {
		tmp = append(tmp, f(elem))
	}

	return tmp
}

// Sort is a wrapper for sort.Slice.
// It takes a compare function as its parameter
// which is used to determine order.
// a < b for ascending order
// a > b for descending order
func (l List[T]) Sort(compare func(T, T) bool) List[T] {
	tmp := make(List[T], len(l))
	copy(tmp, l)

	tmp = bubbleSort(tmp, compare)

	return tmp
}

// I know... just need an alternative to sort.Slice which
// has a compare function that isn't indexes
// TODO: use a better algorithm
func bubbleSort[T any](list List[T], compare func(T, T) bool) List[T] {
	for i := 0; i < len(list)-1; i++ {
		for j := 0; j < len(list)-i-1; j++ {
			if compare(list[j+1], list[j]) {
				list[j], list[j+1] = list[j+1], list[j]
			}
		}
	}
	return list
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
func (l List[T]) Reduce(f func(prev, curr T) T, initialValue T) T {
	accumulator := initialValue
	if l.Len() == 0 {
		return accumulator
	}

	for _, elem := range l {
		accumulator = f(accumulator, elem)
	}

	return accumulator
}

// Filter calls a filter function for
// each element in the List. If the
// return value of the filter function is
// true, then the element is kept in
// the List, otherwise it is removed
func (l List[T]) Filter(f func(T) bool) List[T] {
	tmp := List[T]{}
	for _, elem := range l {
		if f(elem) {
			tmp = append(tmp, elem)
		}

	}

	return tmp
}

// Fill fills the List with val starting at start,
// ending at one before end. This allows for
// easy use of cap(myList) or len(myList).
func (l List[T]) Fill(val T, start, end int) List[T] {
	tmp := make(List[T], len(l))
	copy(tmp, l)

	for i := start; i < end; i++ {
		tmp[i] = val
	}

	return tmp
}

// Reverse returns a List with the items
// in reverse order.
func (l List[T]) Reverse() List[T] {
	tmp := make(List[T], len(l))

	for i, elem := range l {
		tmp[len(l)-1-i] = elem
	}

	return tmp
}
