package heap

type GenNode interface {
	Key() int
}

type GenSliceHeap[T GenNode] []T

func NewGenSliceHeap[T GenNode](values ...T) *GenSliceHeap[T] {
	h := make(GenSliceHeap[T], 0, len(values))
	h = append(h, values...)
	return &h
}

func (h GenSliceHeap[T]) Len() int           { return len(h) }
func (h GenSliceHeap[T]) Less(i, j int) bool { return h[i].Key() > h[j].Key() }
func (h GenSliceHeap[T]) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *GenSliceHeap[T]) Push(x any) {
	*h = append(*h, x.(T))
}

func (h *GenSliceHeap[T]) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
