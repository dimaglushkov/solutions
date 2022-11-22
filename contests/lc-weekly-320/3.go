package main

// WA
func minimumFuelCost(roads [][]int, seats int) int64 {
	graph := make(map[int][]int)
	for _, r := range roads {
		if _, ok := graph[r[0]]; !ok {
			graph[r[0]] = make([]int, 0, 2)
		}
		if _, ok := graph[r[1]]; !ok {
			graph[r[1]] = make([]int, 0, 2)
		}
		graph[r[0]] = append(graph[r[0]], r[1])
		graph[r[1]] = append(graph[r[1]], r[0])
	}

	var cnt int
	var farestCity func(src, dst int)
	farestCity = func(src, dst int) {
		cnt++
		var newSrc, newDst int
		newSrc = dst
		if len(graph[dst]) == 1 {
			return
		}
		if graph[dst][0] == src {
			newDst = graph[dst][1]
		} else {
			newDst = graph[dst][0]
		}
		farestCity(newSrc, newDst)
	}

	var res int64
	for _, c := range graph[0] {
		cnt = 0
		farestCity(0, c)
		for i := 1; i <= cnt; i++ {
			res += int64(i / seats)
			if i%seats != 0 {
				res++
			}
		}
	}

	return res
}

func main() {
	minimumFuelCost([][]int{{3, 1}, {3, 2}, {1, 0}, {0, 4}, {0, 5}, {4, 6}}, 2)
}
