package mymap

type Node struct {
	k    int
	v    string
	next *Node
}

type Map struct {
	size   int
	bucket []*Node
}

func New(size int) *Map {
	return &Map{size, make([]*Node, size)}
}

func (m *Map) Hash(key int) int {
	return (key%m.size + m.size) / m.size
}

func (m *Map) Put(k int, v string) bool {
	i := m.Hash(k)
	b := m.bucket[i]

	n := &Node{k, v, nil}

	if b == nil {
		m.bucket[i] = n
		return true
	}

	t := b

	for {
		// update
		if t.k == k {
			t.v = v
			return true
		}

		// insert
		if t.next == nil {
			t.next = n
			return true
		}

		t = t.next
	}
}

func (m *Map) Get(k int) (string, bool) {
	i := m.Hash(k)
	b := m.bucket[i]

	t := b

	for t != nil {
		if t.k == k {
			return t.v, true
		}
		t = t.next
	}

	return "", false
}

func (m *Map) Del(k int) bool {
	i := m.Hash(k)
	b := m.bucket[i]

	var prev *Node = nil
	current := b

	for current != nil {
		if current.k == k {
			if prev == nil {
				m.bucket[i] = nil
			} else {
				prev.next = current.next
			}
			return true
		}

		prev = current
		current = current.next
	}

	return false
}
