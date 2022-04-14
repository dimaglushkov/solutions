package heap

type Node interface {
	Key() int
}

type SliceHeap[T Node] []T

func NewSliceHeap[T Node](values ...T) *SliceHeap[T] {
	h := make(SliceHeap[T], 0, len(values))
	h = append(h, values...)
	return &h
}

func (h SliceHeap[T]) Len() int           { return len(h) }
func (h SliceHeap[T]) Less(i, j int) bool { return h[i].Key() > h[j].Key() }
func (h SliceHeap[T]) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *SliceHeap[T]) Push(x any) {
	*h = append(*h, x.(T))
}

func (h *SliceHeap[T]) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
