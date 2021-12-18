package math

import "math"

// number is a type constraint which is any
// primitive Go number
type number interface {
	intX | floatX | uintX
}

type floatX interface {
	float32 | float64
}

type intX interface {
	int | int8 | int16 | int32 | int64
}

type uintX interface {
	uint | uint8 | uint16 | uint32 | uint64
}

func Abs[T number](num T) T {
	return T(math.Abs(float64(num)))
}

func Acos[T floatX](num T) T {
	return T(math.Acos(float64(num)))
}

func Max[T number](a, b T) T {
	return T(math.Max(float64(a), float64(b)))
}

func Mod[T number](a, b T) int {
	return int(math.Mod(float64(a), float64(b)))
}
