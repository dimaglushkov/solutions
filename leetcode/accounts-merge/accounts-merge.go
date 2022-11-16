package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// source: https://leetcode.com/problems/accounts-merge/

const idPref = "%id"

func toId(x int) string {
	return idPref + strconv.FormatInt(int64(x), 10)
}

func fromId(x string) int {
	y, _ := strconv.ParseInt(x[len(idPref):], 10, 32)
	return int(y)
}

func accountsMerge(accounts [][]string) [][]string {
	parent := make(map[string]string, 0)
	var find func(string) string
	var union func(string, string)
	union = func(x, y string) {
		parent[find(x)] = find(y)
	}
	find = func(x string) string {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}

	for i := range accounts {
		x := toId(i)
		parent[x] = x

		for _, s := range accounts[i][1:] {
			if _, ok := parent[s]; ok {
				union(s, x)
			} else {
				parent[s] = s
				union(s, x)
			}
		}
	}

	tmp := make(map[string][]string, 0)
	for i := range parent {
		if strings.HasPrefix(i, idPref) {
			continue
		}
		x := find(i)
		if tmp[x] == nil {
			tmp[x] = make([]string, 0)
		}
		tmp[x] = append(tmp[x], i)
	}

	res := make([][]string, 0)
	for k, v := range tmp {
		x := make([]string, len(v)+1)
		x[0] = accounts[fromId(k)][0]
		sort.Strings(v)
		for i := range v {
			x[i+1] = v[i]
		}
		res = append(res, x)
	}

	return res
}

func main() {
	testCases := []struct {
		accounts [][]string
		want     [][]string
	}{
		{
			accounts: [][]string{{"John", "johnsmith@mail.com", "john_newyork@mail.com"}, {"John", "johnsmith@mail.com", "john00@mail.com"}, {"Mary", "mary@mail.com"}, {"John", "johnnybravo@mail.com"}},
			want:     [][]string{{"John", "john00@mail.com", "john_newyork@mail.com", "johnsmith@mail.com"}, {"Mary", "mary@mail.com"}, {"John", "johnnybravo@mail.com"}},
		},
		{
			accounts: [][]string{{"Gabe", "Gabe0@m.co", "Gabe3@m.co", "Gabe1@m.co"}, {"Kevin", "Kevin3@m.co", "Kevin5@m.co", "Kevin0@m.co"}, {"Ethan", "Ethan5@m.co", "Ethan4@m.co", "Ethan0@m.co"}, {"Hanzo", "Hanzo3@m.co", "Hanzo1@m.co", "Hanzo0@m.co"}, {"Fern", "Fern5@m.co", "Fern1@m.co", "Fern0@m.co"}},
			want:     [][]string{{"Ethan", "Ethan0@m.co", "Ethan4@m.co", "Ethan5@m.co"}, {"Gabe", "Gabe0@m.co", "Gabe1@m.co", "Gabe3@m.co"}, {"Hanzo", "Hanzo0@m.co", "Hanzo1@m.co", "Hanzo3@m.co"}, {"Kevin", "Kevin0@m.co", "Kevin3@m.co", "Kevin5@m.co"}, {"Fern", "Fern0@m.co", "Fern1@m.co", "Fern5@m.co"}},
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := accountsMerge(tc.accounts)
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
