package collections

// comparable is any type that can use
// the == or != operators
type Map[T1 comparable, T2 any] map[T1]T2

// Returns all the keys in the map.
// Order not guaranteed to be consistent
func (m Map[T1, T2]) Keys() []T1 {
	tmp := make([]T1, len(m))

	i := 0
	for k, _ := range m {
		tmp[i] = k
		i++
	}

	return tmp
}

// Returns all the values in the map.
// Order not guaranteed to be consistent
func (m Map[T1, T2]) Values() []T2 {
	tmp := make([]T2, len(m))

	i := 0
	for _, v := range m {
		tmp[i] = v
		i++
	}

	return tmp
}

// Filter loops through each key/value
// pair in the map and passes them
// to the comparator function. If it
// returns true, the values are kept
// in the map, if false, they are
// removed
func (m Map[T1, T2]) Filter(compare func(T1, T2) bool) Map[T1, T2] {
	tmp := Map[T1, T2]{}

	for k, v := range m {
		if compare(k, v) {
			tmp[k] = v
		}
	}

	return tmp
}
