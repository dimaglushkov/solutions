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
