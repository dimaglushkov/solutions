package main

//func pivotInteger(n int) int {
//	for i := 1; i <= n; i++ {
//		x := (n - i + 1) / 2 * (2*i + (n - i + 1))
//		_ = x
//
//		if i/2*(2+(i-1)) == (n-i+1)/2*(2*i+(n-i)) {
//			return i
//		}
//	}
//	return -1
//
//}

func pivotInteger(n int) int {
	ps := make([]int, n+1)
	ps[0] = 0
	for i := 1; i <= n; i++ {
		ps[i] = ps[i-1] + i
	}

	for i := 1; i <= n; i++ {
		if ps[n]-ps[i]+i == ps[i] {
			return i
		}
	}
	return -1
}

func main() {
	pivotInteger(8)
}
