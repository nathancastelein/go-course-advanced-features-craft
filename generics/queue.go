package generics

import "errors"

type Queue struct {
	elements []int
}

func (f *Queue) Push(element int) {
	f.elements = append(f.elements, element)
}

func (f *Queue) Pop() (int, error) {
	if len(f.elements) == 0 {
		var zeroValue int
		return zeroValue, errors.New("no more elements")
	}
	element, newSlice := f.elements[0], f.elements[1:]
	f.elements = newSlice
	return element, nil
}
