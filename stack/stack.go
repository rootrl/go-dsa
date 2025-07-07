package stack

import "fmt"

type stack struct {
	data []int
	sp   int
	c    int
}

func NewStack(c int) *stack {
	return &stack{make([]int, c), -1, c}
}

func (s *stack) Pop() int {
	if s.sp == -1 {
		return -1
	}

	n := s.data[s.sp]

	s.sp--

	return n
}

func (s *stack) Push(n int) bool {
	if s.sp >= s.c-1 {
		return false
	}

	s.sp++
	s.data[s.sp] = n

	return true
}

func (s *stack) Info() {
	fmt.Printf(" \n sp: %d, c: %d, data: %v \n", s.sp, s.c, s.data)
}
