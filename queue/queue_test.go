package queue

import (
	"fmt"
	"testing"
)

func TestQueue(t *testing.T) {

	q := New()
	q.Info()

	fmt.Println(q.Dequeue())
	fmt.Println("enqueue: ", q.Enqueue(1))
	fmt.Println("enqueue: ", q.Enqueue(2))
	fmt.Println("enqueue: ", q.Enqueue(3))
	fmt.Println("enqueue: ", q.Enqueue(4))
	q.Info()
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())

	q.Info()

	fmt.Println("enqueue: ", q.Enqueue(5))
	fmt.Println("enqueue: ", q.Enqueue(6))
	fmt.Println("enqueue: ", q.Enqueue(7))
	fmt.Println("enqueue: ", q.Enqueue(8))
	q.Info()
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())

	q.Info()

}
