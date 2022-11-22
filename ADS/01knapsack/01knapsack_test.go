package knapsack_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	knapsack "github.com/dimaglushkov/solutions/ADS/01knapsack"
)

func TestKnapsack(t *testing.T) {
	testcases := []struct {
		arr  []int
		sum  int
		want bool
	}{
		{
			arr:  []int{1, 5, 5, 10},
			sum:  11,
			want: true,
		},
		{
			arr:  []int{1, 2, 3, 4, 5, 6},
			sum:  11,
			want: true,
		},
		{
			arr:  []int{1, 3, 3, 7},
			sum:  99,
			want: false,
		},
		{
			arr:  []int{1, 3, 3, 7},
			sum:  4,
			want: true,
		},
		{
			arr:  []int{2, 2, 2, 5},
			sum:  3,
			want: false,
		},
	}

	for i, tc := range testcases {
		res := knapsack.Knapsack(tc.arr, tc.sum)
		require.Equal(t, tc.want, res, "Test #%d", i+1)
	}
}
