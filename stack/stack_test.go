package stack

import (
	"fmt"
	"testing"
)

func TestStatck(t *testing.T) {
	s := NewStack(3)

	fmt.Println(s.Pop())
	fmt.Println(s.Push(1))
	fmt.Println(s.Push(2))
	fmt.Println(s.Push(3))
	fmt.Println(s.Push(4))

	s.Info()

	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())

	s.Info()
}
