package main

func minimumCardPickup(cards []int) int {
	var res = 100001
	var memo = make(map[int]int)
	for i, c := range cards {
		if prevI, ok := memo[c]; ok {
			dist := i - prevI + 1
			if dist < res {
				res = dist
			}
		}
		memo[c] = i
	}
	if res == 100001 {
		return -1
	}
	return res
}

func main() {
	println(minimumCardPickup([]int{3, 4, 2, 3, 4, 7})) // 4
	println(minimumCardPickup([]int{1, 0, 5, 3}))       //-1

}
