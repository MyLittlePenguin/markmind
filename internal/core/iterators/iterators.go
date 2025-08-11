package iterators

import "markmind/internal/data/entities"

type Iterator[T any] func(yield func(int, T) bool)

func Iter[T any](list []T) Iterator[T] {
	return func(yield func(int, T) bool) {
		for i, it := range list {
			if !yield(i, it) {
				return
			}
		}
	}
}

func Range[T entities.Number](min, max T) Iterator[T] {
	return func(yield func(int, T) bool) {
		i := 0
		for value := min; value <= max; value++ {
			if !yield(i, value) {
				return
			}
			i++
		}
	}
}

func Map[I, O any](
	in []I,
	mapper func(i int, value I) O,
) Iterator[O] {
	return func(yield func(int, O) bool) {
		for i, it := range in {
			if !yield(i, mapper(i, it)) {
				return
			}
		}
	}
}

func MapIter[I, O any](
	in Iterator[I],
	mapper func(i int, value I) O,
) Iterator[O] {
	return func(yield func(int, O) bool) {
		for i, it := range in {
			if !yield(i, mapper(i, it)) {
				return
			}
		}
	}
}

func (self Iterator[T]) Filter(filter func(i int, value T) bool) Iterator[T] {
	return func(yield func(int, T) bool) {
		newIdx := 0
		for i, it := range self {
			if filter(i, it) {
				if !yield(newIdx, it) {
					return
				}
				newIdx++
			}
		}
	}
}

type Bloated[T any] Iterator[Iterator[T]]

func (self Bloated[T]) Flatten() Iterator[T] {
	return func(yield func(int, T) bool) {
		i := 0
		for _, outer := range self {
			for _, inner := range outer {
				if !yield(i, inner) {
				}
				i++
			}
		}
	}
}

func Flatten[T any](bloated Iterator[Iterator[T]]) Iterator[T] {
	return Bloated[T](bloated).Flatten()
}

func (self Iterator[T]) Any(matcher func(i int, value T) bool) bool {
	for i, it := range self {
		if matcher(i, it) {
			return true
		}
	}
	return false
}

func (self Iterator[T]) Length() int {
	length := 0
	for i := range self {
		i *= 1
		length++
	}
	return length
}

func (self Iterator[T]) Collect() []T {
	result := make([]T, 0)
	for _, it := range self {
		result = append(result, it)
	}
	return result
}
