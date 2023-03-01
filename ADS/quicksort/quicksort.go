package quicksort

import "math/rand"

func Quicksort(data []int) {
	var doSort func(data []int, l, r int)
	doSort = func(data []int, l, r int) {
		if l < r {
			p := partition(data, l, r)
			doSort(data, l, p-1)
			doSort(data, p+1, r)
		}
	}
	doSort(data, 0, len(data)-1)
}

// partition implements algorithm described in CLRS
func partition(data []int, l, r int) int {
	i, pivot := l-1, data[r]
	for j := l; j < r; j++ {
		if data[j] <= pivot {
			i++
			data[i], data[j] = data[j], data[i]
		}
	}
	data[i+1], data[r] = data[r], data[i+1]
	return i + 1
}

// randomPartition implements randomized partition algorithm by randomly selecting pivot
func randomPartition(data []int, l, r int) int {
	x := rand.Intn(r-l) + l
	data[x], data[r] = data[r], data[x]

	i, pivot := l-1, data[r]
	for j := l; j < r; j++ {
		if data[j] <= pivot {
			i++
			data[i], data[j] = data[j], data[i]
		}
	}
	data[i+1], data[r] = data[r], data[i+1]
	return i + 1
}

func QuicksortHoare(data []int) {
	var doSort func(data []int, l, r int)
	doSort = func(data []int, l, r int) {
		if l < r {
			p := hoarePartition(data, l, r)
			doSort(data, l, p)
			doSort(data, p+1, r)
		}
	}
	doSort(data, 0, len(data)-1)
}

// hoarePartition implements hoare partition
func hoarePartition(data []int, i, j int) int {
	pivot := data[i]
	for {
		for data[i] < pivot {
			i++
		}
		for data[j] > pivot {
			j--
		}
		if i >= j {
			return j
		}
		data[i], data[j] = data[j], data[i]
		i++
		j--
	}
}
