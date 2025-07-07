package single


import (
	"testing"
)

func TestSingle(t * testing.T) {
	list := NewList()

	Add(list, 1)
	Add(list, 2)
	Add(list, 3)

	Traverse(list)

	Del(list, 2)

	Traverse(list)
}
