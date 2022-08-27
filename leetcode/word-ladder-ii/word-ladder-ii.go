package main

import (
	"fmt"
	"math"
)

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	l := len(wordList)
	wl := len(beginWord)
	var src, dst int
	data := make([][]byte, l+1)
	foundEndWord := false
	foundBeginWord := false
	for i := 0; i < l; i++ {
		if wordList[i] == endWord {
			dst = i
			foundEndWord = true
		}
		if wordList[i] == beginWord {
			src = i
			foundBeginWord = true
		}

		data[i] = []byte(wordList[i])
	}

	if !foundEndWord {
		return [][]string{}
	}
	if !foundBeginWord {
		data[l] = []byte(beginWord)
		src = l
		wordList = append(wordList, beginWord)
	}

	adj := make([][]int, l+1)
	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			if findLink(i, j, data, wl) {
				adj[i] = append(adj[i], j)
				adj[j] = append(adj[j], i)
			}
		}
	}
	if !foundBeginWord {
		for i := 0; i < l; i++ {
			if findLink(i, l, data, wl) {
				adj[l] = append(adj[l], i)
			}
		}
	}

	//remove loops from endword
	adj[dst] = nil

	q := make([]int, 0)
	q = append(q, src)
	cost := make([]int, l+1)
	parent := make([][]int, l+1)
	for i := 0; i <= l; i++ {
		cost[i] = math.MaxInt32
		//parent[i]=make([]int,0)
	}
	parent[src] = append(parent[src], -1)
	cost[src] = 0

	for len(q) > 0 {
		top := q[0]

		// if top==dst{
		//     return cost[dst]+1
		// }
		q = q[1:]
		for _, t := range adj[top] {
			if cost[t] > 1+cost[top] {
				cost[t] = 1 + cost[top]
				q = append(q, t)
				parent[t] = make([]int, 0)
				parent[t] = append(parent[t], top)
			} else if cost[t] == 1+cost[top] {
				//q = append(q,t)
				parent[t] = append(parent[t], top)
			}
		}
	}

	res := make([][]string, 0)
	path(&res, []string{}, parent, src, dst, wordList)

	for i := range res {
		reverse(&res[i])
	}

	return res
}

func path(res *[][]string, p []string, parent [][]int, src, dst int, wordList []string) {
	if dst == -1 {
		*res = append(*res, p)
		return
	}
	// fmt.Println(dst)
	p = append(p, wordList[dst])
	for _, v := range parent[dst] {
		path(res, p, parent, src, v, wordList)
	}
	p = p[:len(p)-1]
}

func findLink(i, j int, data [][]byte, sz int) bool {
	cnt := 0
	for k := 0; k < sz; k++ {
		if data[i][k] == data[j][k] {
			cnt++
		}
	}
	return cnt == sz-1
}

func reverse(lst *[]string) {
	l := len(*lst)

	i, j := 0, l-1
	for i <= j {
		(*lst)[i], (*lst)[j] = (*lst)[j], (*lst)[i]
		i++
		j--
	}
}

func main() {
	// Example 1
	var beginWord1 string = "hit"
	var endWord1 string = "cog"
	var wordList1 = []string{"hot", "dot", "dog", "lot", "log", "cog"}
	fmt.Println("Expected: [[\"hit\",\"hot\",\"dot\",\"dog\",\"cog\"],[\"hit\",\"hot\",\"lot\",\"log\",\"cog\"]]	Output: ", findLadders(beginWord1, endWord1, wordList1))

	// Example 2
	var beginWord2 string = "hit"
	var endWord2 string = "cog"
	var wordList2 = []string{"hot", "dot", "dog", "lot", "log"}
	fmt.Println("Expected: []	Output: ", findLadders(beginWord2, endWord2, wordList2))

}
