package heap

type heap struct {
	data []int
}

func New(arr []int) *heap {
	return &heap{arr}
}

func Heapify(arr []int) *heap {

	h := New(arr)
	p := len(arr)/2 - 1

	for i := p; i >= 0; i-- {
		h.SiftDown(i)
	}

	return h
}

func (h *heap) SiftDown(i int) {

	size := len(h.data)

	for {
		m := i
		li := 2*i + 1
		ri := 2*i + 2

		if li < size && h.data[li] <= h.data[m] {
			m = li
		}

		if ri < size && h.data[ri] <= h.data[m] {
			m = ri
		}

		if m == i || h.data[i] == h.data[m] {
			break
		}

		h.data[i], h.data[m] = h.data[m], h.data[i]

		i = m
	}
}

func (h *heap) SiftUp(i int) {

	for i > 0 {
		p := (i - 1) / 2
		if h.data[p] < h.data[i] {
			break
		}
		h.data[i], h.data[p] = h.data[p], h.data[i]
		i = p
	}
}

func (h *heap) Insert(n int) {
	h.data = append(h.data, n)
	h.SiftUp(len(h.data) - 1)
}

func (h *heap) ExtractMin() int {

	l := len(h.data)

	if l == 0 {
		return -1
	}

	m := h.data[0]

	last := h.data[l-1]
	h.data = h.data[:l-1]

	if len(h.data) > 0 {
		h.data[0] = last
		h.SiftDown(0)
	}

	return m
}
