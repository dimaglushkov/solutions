package rabin_karp_test

import (
	rabin_karp "github.com/dimaglushkov/solutions/ADS/rabin-karp"
	"strings"
	"testing"
)

// TestRabinKarpLen implements pretty simple tests to check if returning value
// in RabinKarp implementation returns an answer of len equal to strings.Count().
// However, strings.Count() does not count overlapping substrings and RabinKarp does
// hence, these tests are weak. Additionally, it only checks for a len of a returned value
// and doesn't check if the actual values are correct
func TestRabinKarpLen(t *testing.T) {
	tcs := []struct {
		text    string
		pattern string
	}{
		{
			text:    "abacdbdcnabaxabajsbcjdbsjasnfasf",
			pattern: "aba",
		},
		{
			text:    "aaaaaaaa",
			pattern: "a",
		},
		{
			text:    "abcd",
			pattern: "xxxxxx",
		},
		{
			text:    "abcd",
			pattern: "x",
		},
		{
			text:    "abcdbdcnajsbcjdbsjasnfasf",
			pattern: "a",
		},
		{
			text:    "",
			pattern: "",
		},
		{
			text:    "abcdbdcnajsbcjdbsjasnfasf",
			pattern: "",
		},
	}
	for i, tc := range tcs {
		ans := rabin_karp.RabinKarp(tc.text, tc.pattern)
		if expected := strings.Count(tc.text, tc.pattern); len(ans) != expected {
			t.Errorf("wrong answer len for test case %d: expected %d, actual %d", i+1, expected, len(ans))
		}
	}
}
