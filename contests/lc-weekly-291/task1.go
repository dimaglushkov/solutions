package main

import "sort"

func removeDigit(number string, digit byte) string {
	var variants = []string{}

	for i := range number {
		if number[i] == digit {
			variants = append(variants, number[:i]+number[i+1:])
		}
	}

	sort.Strings(variants)

	return variants[len(variants)-1]
}

func main() {
	removeDigit("9239", '9')
}
