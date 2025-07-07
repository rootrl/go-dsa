package heap

import (
	"fmt"
	"testing"
)

func TestHeap(t *testing.T) {

	arr := []int{30, 80, 10, 50, 20, 90, 5}

	h := Heapify(arr)
	fmt.Println(h)

	h.Insert(3)
	fmt.Println(h)

	fmt.Println(h.ExtractMin())
	fmt.Println(h.ExtractMin())
	fmt.Println(h.ExtractMin())
	fmt.Println(h.ExtractMin())
	fmt.Println(h.ExtractMin())
	fmt.Println(h.ExtractMin())
	fmt.Println(h.ExtractMin())
	fmt.Println(h.ExtractMin())
	fmt.Println(h.ExtractMin())
}
