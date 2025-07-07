package queue

import "fmt"

type Queue struct {
	data []int
}

func New() *Queue {
	return &Queue{
		make([]int, 0),
	}
}

func (q *Queue) Enqueue(n int) bool {
	q.data = append(q.data, n)
	return true
}

func (q *Queue) Dequeue() (int, bool) {

	if len(q.data) < 1 {
		return -1, false
	}

	n := q.data[0]
	q.data = q.data[1:]

	return n, true

}

func (q *Queue) Info() {
	fmt.Printf("data %v  \n", q.data)
}
