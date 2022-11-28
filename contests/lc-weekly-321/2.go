package main

func appendCharacters(s string, t string) int {
	i := 0
	rt := []rune(t)
	for _, r := range s {
		if r == rt[i] {
			i++
		}
		if i == len(rt) {
			return 0
		}
	}
	return len(rt) - i
}

func main() {
	appendCharacters("z", "abcde")
	appendCharacters("abcde", "a")
}
