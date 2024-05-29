package generics

import "errors"

type Queue[T any] struct {
	elements []T
}

func (f *Queue[T]) Push(element T) {
	f.elements = append(f.elements, element)
}

func (f *Queue[T]) Pop() (T, error) {
	if len(f.elements) == 0 {
		var zeroValue T
		return zeroValue, errors.New("no more elements")
	}
	element, newSlice := f.elements[0], f.elements[1:]
	f.elements = newSlice
	return element, nil
}
