package list

type GenListNode[T any] struct {
	Val  T
	Next *GenListNode[T]
}

func NewGenListNode[T any](values ...T) *GenListNode[T] {
	head := new(GenListNode[T])
	p := head
	for _, v := range values {
		p.Next = &GenListNode[T]{Val: v}
	}
	return head.Next
}

