package disjoint_set

type DisjointSet struct {
	parent []int
}

func New() DisjointSet {
	return DisjointSet{parent: make([]int, 0)}
}

func (ds *DisjointSet) union(x, y int) {
	ds.parent[ds.find(x)] = ds.find(y)
}

func (ds *DisjointSet) find(x int) int {
	if ds.parent[x] != x {
		ds.parent[x] = ds.find(ds.parent[x])
	}
	return ds.parent[x]
}

func example(data [][]any) {
	parent := make([]int, len(data))
	var find func(int) int
	var union func(int, int)

	union = func(x, y int) {
		parent[find(x)] = find(y)
	}
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}

	for i := range data {
		parent[i] = i
	}

	for i := range data {
		union(data[i][0].(int), data[i][1].(int))
	}
}
