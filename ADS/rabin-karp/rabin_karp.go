package rabin_karp

import "math"

const (
	MOD = 1_000_003
	D   = 7
)

func pow(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}

// RabinKarp implements rolling hash algorithm to search
// for a string in a given text
func RabinKarp(text, pattern string) []int {
	if len(text) < len(pattern) {
		return nil
	}

	if len(pattern) == 0 {
		var ans []int
		for i := range text {
			ans = append(ans, i)
		}
		ans = append(ans, len(ans))
		return ans
	}

	n := len(pattern)
	hash, pHash := 0, 0
	for i := range pattern {
		hash += int(text[i]) * pow(D, n-1-i)
		hash %= MOD

		pHash += int(pattern[i]) * pow(D, n-1-i)
		pHash %= MOD
	}

	var ans []int
	if text[:n] == pattern {
		ans = append(ans, 0)
	}

	dm := pow(D, n-1)
	for i := 1; i <= len(text)-n; i++ {
		x := text[i : i+n]
		_ = x
		hash -= int(text[i-1]) * dm
		hash *= D
		hash += int(text[i+n-1])
		hash %= MOD

		if hash == pHash {
			if text[i:i+n] == pattern {
				ans = append(ans, i)
			}
		}
	}

	if len(ans) == 0 {
		return nil
	}
	return ans
}
