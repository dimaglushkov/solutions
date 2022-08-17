package main

// source: https://leetcode.com/problems/first-unique-character-in-a-string/

func firstUniqChar(s string) int {
	m := make(map[rune]int)
	d := make(map[rune]bool)
	for i, r := range s {
		if _, ok := d[r]; ok {
			continue
		}
		if _, ok := m[r]; ok {
			delete(m, r)
			d[r] = true
			continue
		}
		m[r] = i
	}

	if len(m) == 0 {
		return -1
	}

	res := 1000000
	for _, i := range m {
		if i < res {
			res = i
		}
	}

	return res

}

func main() {
	firstUniqChar("aabb")
}
