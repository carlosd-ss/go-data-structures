package main

import (
	"fmt"
)

type Stack struct {
	items []int
}

func (s *Stack) Push(i int) {
	s.items = append(s.items, i)

}
func (s *Stack) Pop() int {
	if len(s.items) > 0 {
		lastindex := len(s.items) - 1
		remove := s.items[lastindex]
		s.items = s.items[:lastindex]
		return remove
	}
	return 0
}
func (s *Stack) Size() int {
	return len(s.items)
}
func (s *Stack) Empty() bool {
	return len(s.items) == 0
}

func (s *Stack) Stackpop() int {
	ultimoindex := len(s.items) - 1
	s.items = s.items[ultimoindex:]
	return ultimoindex
}

func main() {
	stack := Stack{}
	fmt.Println("Stack: ", stack)

	//Push
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	fmt.Println("Push: ", stack)

	//Pop
	stack.Pop()
	fmt.Println("Pop: ", stack)

	//Size
	stack.Size()
	fmt.Println("Size: ", stack)

	//Stackpop
	stack.Stackpop()
	fmt.Println("Stackpop: ", stack)

	//Empty
	stack.Empty()
}
