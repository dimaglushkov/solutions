package main

// source: https://leetcode.com/problems/subrectangle-queries/

type SubrectangleQueries [][]int

// NewSubrectangleQueries rename to Constructor for leetcode
func NewSubrectangleQueries(rectangle [][]int) SubrectangleQueries {
	return rectangle
}

func (q *SubrectangleQueries) UpdateSubrectangle(row1 int, col1 int, row2 int, col2 int, newValue int) {
	for i := row1; i <= row2; i++ {
		for j := col1; j <= col2; j++ {
			(*q)[i][j] = newValue
		}
	}
}

func (q *SubrectangleQueries) GetValue(row int, col int) int {
	return (*q)[row][col]
}

/**
 * Your SubrectangleQueries object will be instantiated and called as such:
 * obj := Constructor(rectangle);
 * obj.UpdateSubrectangle(row1,col1,row2,col2,newValue);
 * param_2 := obj.GetValue(row,col);
 */

func main() {
	q := NewSubrectangleQueries([][]int{{1, 2, 1}, {4, 3, 4}, {3, 2, 1}, {1, 1, 1}})
	println(q.GetValue(0, 2))
	q.UpdateSubrectangle(0, 0, 3, 2, 5)
	println(q.GetValue(0, 2))
	println(q.GetValue(3, 1))
	q.UpdateSubrectangle(3, 0, 3, 2, 10)
	println(q.GetValue(3, 1))
	println(q.GetValue(0, 2))
}
