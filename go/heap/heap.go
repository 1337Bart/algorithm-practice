package heap

type Item struct {
	Node   int
	Weight int
}

type BinaryHeap struct {
	Heap []*Item
}

func NewBinaryHeap() *BinaryHeap {
	return &BinaryHeap{}
}

func (h *BinaryHeap) Insert(item *Item) {
	h.Heap = append(h.Heap, item)
	h.siftUp(len(h.Heap) - 1)
}

func (h *BinaryHeap) siftUp(i int) {
	for i > 0 {
		parent := (i - 1) / 2
		if h.Heap[i].Weight < h.Heap[parent].Weight {
			h.Heap[i], h.Heap[parent] = h.Heap[parent], h.Heap[i]
			i = parent
		} else {
			break
		}
	}
}

func (h *BinaryHeap) DeleteMin() *Item {
	min := h.Heap[0]
	h.Heap[0] = h.Heap[len(h.Heap)-1]
	h.Heap = h.Heap[:len(h.Heap)-1]
	h.siftDown(0)
	return min
}

func (h *BinaryHeap) siftDown(i int) {
	for {
		left := 2*i + 1
		right := 2*i + 2
		min := i

		if left < len(h.Heap) && h.Heap[left].Weight < h.Heap[min].Weight {
			min = left
		}
		if right < len(h.Heap) && h.Heap[right].Weight < h.Heap[min].Weight {
			min = right
		}
		if min == i {
			break
		}

		h.Heap[i], h.Heap[min] = h.Heap[min], h.Heap[i]
		i = min
	}
}
