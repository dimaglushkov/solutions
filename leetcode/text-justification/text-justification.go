package main

import (
	"fmt"
	"strings"
)

// source: https://leetcode.com/problems/text-justification/

func fullJustify(words []string, maxWidth int) []string {
	ans := make([]string, 0)
	textWords := make([][]string, 0)
	textWordsLens := make([]int, 0)
	i := 0

	// collecting words
	for i < len(words) {
		textWords = append(textWords, make([]string, 0))
		ln := len(textWords) - 1
		strLen := 0
		for i < len(words) && strLen+len(textWords[ln])-1+len(words[i]) < maxWidth {
			textWords[ln] = append(textWords[ln], words[i])
			strLen += len(words[i])
			i++
		}
		textWordsLens = append(textWordsLens, strLen)
	}

	// counting spaces and building strings
	for i = 0; i < len(textWords); i++ {
		// if there is only one word or this is the last row, then left alignemnt
		if len(textWords[i]) == 1 || i == len(textWords)-1 {
			str := strings.Join(textWords[i], " ")
			for len(str) < maxWidth {
				str += " "
			}
			ans = append(ans, str)
			continue
		}
		strSpaces := make([]string, len(textWords[i]))
		strSpacesNum := maxWidth - textWordsLens[i]
		for j := 0; strSpacesNum > 0; strSpacesNum-- {
			strSpaces[j] += " "
			j++
			if j == len(strSpaces)-1 {
				j = 0
			}
		}
		// building the string
		str := ""
		for j := range textWords[i] {
			str += textWords[i][j] + strSpaces[j]
		}
		ans = append(ans, str)
	}

	return ans
}

func main() {
	testCases := []struct {
		words    []string
		maxWidth int
		want     []string
	}{
		{
			words:    []string{"This", "is", "an", "example", "of", "text", "justification."},
			maxWidth: 16,
			want: []string{
				"This    is    an",
				"example  of text",
				"justification.  ",
			},
		},
		{
			words:    []string{"What", "must", "be", "acknowledgment", "shall", "be"},
			maxWidth: 16,
			want: []string{
				"What   must   be",
				"acknowledgment  ",
				"shall be        ",
			},
		},
		{
			words:    []string{"Science", "is", "what", "we", "understand", "well", "enough", "to", "explain", "to", "a", "computer.", "Art", "is", "everything", "else", "we", "do"},
			maxWidth: 20,
			want: []string{
				"Science  is  what we",
				"understand      well",
				"enough to explain to",
				"a  computer.  Art is",
				"everything  else  we",
				"do                  ",
			},
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := fullJustify(tc.words, tc.maxWidth)
		status := "ERROR"
		if fmt.Sprint(x) == fmt.Sprint(tc.want) {
			status = "OK"
			successes++
		}
		fmt.Println(status, "	Expected: ", tc.want, " Actual: ", x)
	}
	if l := len(testCases); successes == len(testCases) {
		fmt.Printf("===\nSUCCESS: %d of %d tests ended successfully\n", successes, l)
	} else {
		fmt.Printf("===\nFAIL: %d tests failed\n", l-successes)
	}

}
