package heap

type SliceHeap []int

func NewSliceHeap(values ...int) *SliceHeap {
	h := make(SliceHeap, 0, len(values))
	h = append(h, values...)
	return &h
}

func (h SliceHeap) Len() int           { return len(h) }
func (h SliceHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h SliceHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *SliceHeap) Push(x int) {
	*h = append(*h, x)
}

func (h *SliceHeap) Pop() int {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
