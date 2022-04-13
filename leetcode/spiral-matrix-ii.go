package main

import "fmt"

// source: https://leetcode.com/problems/spiral-matrix-ii/

/*int r1 = 0, r2 = n-1;
        int c1 = 0, c2 = n-1;
        int val = 0;

		// result matrix
        vector<vector<int>> v(n, vector<int> (n));
        while(r1 <= r2 && c1 <= c2)
        {
            // left to right (row will be fixed)
            for(int i = c1; i <= c2; ++i)
                v[r1][i] = ++val;

				// move down(col will be fixed)
            for(int i = r1+1; i <= r2; ++i)
                v[i][c2] = ++val;

            // move right to left
            // move  up
            if(r1 < r2 && c1 < c2)
            {
                // move right to left (row will be fixed)
                for(int i = c2-1; i>c1; --i)
                    v[r2][i] = ++val;

					// move up (col will be fixed)
					for(int i = r2; i>r1; --i)
                    v[i][c1] = ++val;
            }
            ++r1;
            --r2;
            ++c1;
            --c2;
        }
        return v;
    }*/

func generateMatrix(n int) [][]int {
	var rl, rr = 0, n - 1
	var cl, cr = 0, n - 1
	var val = 1
	res := make([][]int, n)
	for i := range res {
		res[i] = make([]int, n)
	}

	for val <= n*n {
		for i := cl; i <= cr; i++ {
			res[rl][i] = val
			val++
		}
		for i := rl + 1; i <= rr; i++ {
			res[i][cr] = val
			val++
		}

		for i := cr - 1; i > cl; i-- {
			res[rr][i] = val
			val++
		}
		for i := rr; i > rl; i-- {
			res[i][cl] = val
			val++
		}

		rl++
		rr--
		cl++
		cr--

	}
	return res
}

func main() {
	// Example 3
	var n3 int = 4
	fmt.Println("Expected: [[1,2,3],[8,9,4],[7,6,5]]	Output: ", generateMatrix(n3))

	// Example 1
	var n1 int = 3
	fmt.Println("Expected: [[1,2,3],[8,9,4],[7,6,5]]	Output: ", generateMatrix(n1))

	// Example 2
	var n2 int = 1
	fmt.Println("Expected: [[1]]	Output: ", generateMatrix(n2))

}
