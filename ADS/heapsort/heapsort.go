package heapsort

// HeapSort implements in-place heap-sorting
func HeapSort(data []int) {
	CreateHeap(data)
	for i := len(data) - 1; i > 0; i-- {
		data[0], data[i] = data[i], data[0]
		Heapify(data, 0, i)
	}
}

func CreateHeap(data []int) {
	n := len(data)
	for i := n/2 - 1; i >= 0; i-- {
		Heapify(data, i, n)
	}
}

func Heapify(data []int, i, n int) {
	iMax, iLeft, iRight := i, 2*i+1, 2*i+2
	if iLeft < n && data[iLeft] > data[iMax] {
		iMax = iLeft
	}

	if iRight < n && data[iRight] > data[iMax] {
		iMax = iRight
	}

	if iMax != i {
		data[i], data[iMax] = data[iMax], data[i]
		Heapify(data, iMax, n)
	}
}
