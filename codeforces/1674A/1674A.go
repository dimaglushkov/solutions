package main

import (
    "bufio"
    "fmt"
    "math"
    "os"
)

// source: https://codeforces.com/contest/1674/problem/A?locale=en

var in, out = bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)

func solve() string {
    var x, y int
    fmt.Fscan(in, &x, &y)

    if x > y {
        return "0 0"
    }
    if x == y {
        return "1 1"
    }
    for b := 2; b <= y/x; b++ {
        for a := 1; int(math.Pow(float64(b), float64(a)))*x <= y; a++ {
            tmp := int(math.Pow(float64(b), float64(a))) * x
            if tmp == y {
                return fmt.Sprintf("%d %d", a, b)
            }
        }
    }
    return "0 0"
}

func main() {
    defer out.Flush()
    var t int
    for fmt.Fscanln(in, &t); t > 0; t-- {
        fmt.Fprintln(out, solve())
    }
}
