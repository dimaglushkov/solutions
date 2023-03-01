package quicksort_test

import (
	"github.com/dimaglushkov/solutions/ADS/quicksort"
	"log"
	"math/rand"
	"sort"
	"testing"
	"time"
)

const (
	iterations           = 20
	minUnsortedDataSize  = 10000
	maxUnsortedDataSize  = 100000
	maxUnsortedDataValue = 1000000
)

var (
	seed int64
	rng  *rand.Rand
)

func init() {
	seed = time.Now().UTC().UnixNano()
	rng = rand.New(rand.NewSource(seed))
	log.Printf("rng seed: %d", seed)
}

func generateUnsortedData(minSize, maxSize, maxValue int) []int {
	var n int
	if minSize != maxSize {
		n = rng.Intn(maxSize-minSize) + minSize
	} else {
		n = minSize
	}
	data := make([]int, n)
	for i := range data {
		data[i] = rng.Intn(maxValue)
	}
	return data
}

func TestQuicksortRandomData(t *testing.T) {
	for tc := 0; tc < iterations; tc++ {
		var d1, d2 []int
		d1 = generateUnsortedData(minUnsortedDataSize, maxUnsortedDataSize, maxUnsortedDataValue)
		d2 = make([]int, len(d1))
		copy(d2, d1)

		quicksort.Quicksort(d1)
		sort.Ints(d2)

		if len(d1) != len(d2) {
			t.Errorf("unequal lens after performing sorting: len(d1)=%d, len(d2)=%d", len(d1), len(d2))
			t.FailNow()
		}

		for i := range d1 {
			if d1[i] != d2[i] {
				t.Errorf("unequal values at position %d: d1[%d]=%d, d2[%d]=%d", i, i, d1[i], i, d2[i])
				t.FailNow()
			}
		}
	}
}

func TestQuicksortHoareRandomData(t *testing.T) {
	for tc := 0; tc < iterations; tc++ {
		var d1, d2 []int
		d1 = generateUnsortedData(minUnsortedDataSize, maxUnsortedDataSize, maxUnsortedDataValue)
		d2 = make([]int, len(d1))
		copy(d2, d1)

		quicksort.QuicksortHoare(d1)
		sort.Ints(d2)

		if len(d1) != len(d2) {
			t.Errorf("unequal lens after performing sorting: len(d1)=%d, len(d2)=%d", len(d1), len(d2))
			t.FailNow()
		}

		for i := range d1 {
			if d1[i] != d2[i] {
				t.Errorf("unequal values at position %d: d1[%d]=%d, d2[%d]=%d", i, i, d1[i], i, d2[i])
				t.FailNow()
			}
		}
	}
}
