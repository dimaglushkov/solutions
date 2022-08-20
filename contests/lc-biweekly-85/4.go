package main

// Memory limit exceeded
// probably have to change sumArr logic to create nlog n matrix instead of n*n
func sumArr(x []int) [][]int {
	s := make([][]int, len(x))

	for i := 0; i < len(x); i++ {
		s[i] = make([]int, len(x))
		s[i][i] = x[i]
	}

	for i := 0; i < len(x); i++ {
		for j := i + 1; j < len(x); j++ {
			s[i][j] = s[i][j-1] + x[j]
		}
	}
	return s
}

func maximumSegmentSum(nums []int, removeQueries []int) []int64 {
	sums := sumArr(nums)
	res := make([]int64, len(removeQueries))

	segments := [][]int{{0, len(nums) - 1}}

	for i, rem := range removeQueries {
		for j := range segments {
			if segments[j][0] != -1 && rem >= segments[j][0] && rem <= segments[j][1] {
				if rem == segments[j][0] {
					segments[j][0]++
				} else if rem == segments[j][1] {
					segments[j][1]--
				} else {
					end := segments[j][1]
					segments[j][1] = rem - 1
					segments = append(segments, []int{rem + 1, end})
				}

				if segments[j][0] > segments[j][1] {
					segments[j][0], segments[j][1] = -1, -1
				}
				break
			}
		}
		var max = -1
		for _, s := range segments {
			if s[0] != -1 && sums[s[0]][s[1]] > max {
				max = sums[s[0]][s[1]]
			}
		}
		res[i] = int64(max)

	}
	res[len(res)-1] = 0
	return res
}

func main() {
	maximumSegmentSum([]int{1, 2, 5, 6, 1}, []int{0, 3, 2, 4, 1})
	maximumSegmentSum([]int{3, 2, 11, 1}, []int{3, 2, 1, 0})

}
