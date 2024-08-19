package heap

type MinHeap struct {
	data   []int
	length int
}

func (h *MinHeap) insert(val int) {
	h.data[h.length] = val
	h.length++
}

func (h *MinHeap) delete() int {
	if h.length == 0 {
		return -1
	}

	h.length--

	out := h.data[0]
	if h.length == 0 {
		h.data = []int{}

		return out
	}

	h.data[0] = h.data[h.length]
	h.heapifyDown(0)

	return out
}

func (h *MinHeap) parent(idx int) int {
	return int((idx - 1) / 2)
}

func (h *MinHeap) leftChild(idx int) int {
	return idx*2 + 1
}

func (h *MinHeap) rightChild(idx int) int {
	return idx*2 + 2
}

func (h *MinHeap) heapifyUp(idx int) {
	if idx == 0 {
		return
	}

	pIdx := h.parent(idx)
	pV := h.data[pIdx]
	v := h.data[idx]

	if v < pV {
		h.data[v] = pV
		h.data[pV] = v
		h.heapifyUp(pIdx)
	}
}

func (h *MinHeap) heapifyDown(idx int) {
	lIdx := h.leftChild(idx)
	rIdx := h.rightChild(idx)

	if idx >= h.length || lIdx >= h.length {
		return
	}

	lV := h.data[lIdx]
	rV := h.data[rIdx]
	v := h.data[idx]

	if rV > lV && v > lV {
		h.data[idx] = lV
		h.data[lIdx] = v
		h.heapifyDown(lIdx)
	} else if lV > rV && v > rV {
		h.data[idx] = rV
		h.data[rIdx] = v
		h.heapifyDown(rIdx)
	}
}
