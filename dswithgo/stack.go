package dswithgo

import (
	"errors"
	"fmt"
)

type Stack struct {
	data   []int
	Length int
}

func (s *Stack) Push(num int) {
	s.data = append(s.data, num)
	s.Length = len(s.data)
}

func (s *Stack) Pop() (int, error) {
	if s.Length == 0 {
		return 0, errors.New("the stack is empty")
	}
	popedValue := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	s.Length = len(s.data)
	return popedValue, nil
}

func (s *Stack) Print() {
	fmt.Println(s.data, " Length: ", s.Length)
}
