package main

import "sort"

// TLE :(
func countRectangles(rectangles [][]int, points [][]int) []int {
	res := make([]int, len(points))

	sort.Slice(rectangles[:], func(i, j int) bool {
		return rectangles[i][0] < rectangles[j][0]
	})

	buckets := make(map[[2]int][][]int)

loop:
	for _, r := range rectangles {
		for bucket := range buckets {
			if r[0] <= bucket[0] && r[1] <= bucket[1] {
				buckets[[2]int{r[0], r[1]}] = append([][]int{{bucket[0], bucket[1]}}, buckets[bucket]...)
				delete(buckets, bucket)
				continue loop
			}
		}
		buckets[[2]int{r[0], r[1]}] = nil
	}

	for i := 0; i < len(points); i++ {
		for b := range buckets {
			if points[i][0] <= b[0] && points[i][1] <= b[1] {
				res[i] += len(buckets[b]) + 1
			} else {
				for bi, bb := range buckets[b] {
					if points[i][0] <= bb[0] && points[i][1] <= bb[1] {
						res[i] += len(buckets[b]) - bi
						break
					}
				}
			}
		}
	}

	return res
}

func main() {
	countRectangles([][]int{{8, 4}, {10, 8}, {7, 2}, {10, 5}, {4, 7}, {9, 9}, {5, 2}, {1, 5}}, [][]int{{8, 2}, {10, 8}, {8, 8}, {1, 3}})
	countRectangles([][]int{{1, 2}, {2, 3}, {2, 5}}, [][]int{{0, 0}, {2, 1}, {1, 4}})
}
