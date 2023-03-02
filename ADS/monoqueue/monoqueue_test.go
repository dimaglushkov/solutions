package monoqueue_test

import (
	"github.com/dimaglushkov/solutions/ADS/monoqueue"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"sort"
	"testing"
)

func TestMonoQueue(t *testing.T) {
	mq := monoqueue.NewMonoQueue()
	assert.Equal(t, 0, mq.Len())

	for i := 0; i < 10; i++ {
		mq.Push(i)
	}
	for i := 19; i >= 10; i-- {
		mq.Push(i)
	}

	assert.Equal(t, mq.Len(), 20)
	for i := 19; i >= 0; i-- {
		assert.Equal(t, i, mq.Pop())
	}
}

func TestNewMonoQueue(t *testing.T) {
	values := make([]int, 5)
	for i := range values {
		values[i] = rand.Intn(1000)
	}
	mq := monoqueue.NewMonoQueue(values...)

	for range values {
		x := rand.Intn(1000)
		values = append(values, x)
		mq.Push(x)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(values)))
	for _, i := range values {
		assert.Equal(t, i, mq.Pop())
	}

}
