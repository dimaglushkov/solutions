package main

import "fmt"

// source: https://leetcode.com/problems/first-unique-character-in-a-string/

func firstUniqChar(s string) int {
    m := make(map[rune]int)
    for i, r := range s {
        if _, ok := m[r]; ok {
            m[r] = -1
            continue
        }
        m[r] = i
    }

    res := 1000000
    for _, i := range m {
        if i > -1 && i < res {
            res = i
        }
    }

    if res == 1000000 {
        return -1
    }
    return res

}
