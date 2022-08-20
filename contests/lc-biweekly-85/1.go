package main

func minimumRecolors(blocks string, k int) int {
	res := 101
	n := len(blocks)
	for i := 0; i <= n-k; i++ {
		cnt := 0
		for j := i; j < i+k; j++ {
			if blocks[j] == 'W' {
				cnt++
			}
		}
		if cnt < res {
			res = cnt
		}
	}

	return res
}

func main() {
	println(minimumRecolors("WBBWWBBWBW", 7))
	println(minimumRecolors("WBWBBBW", 2))
	println(minimumRecolors("BWWWBB", 6))
	println(minimumRecolors("W", 1))
}
