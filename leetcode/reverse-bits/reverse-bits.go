package main

import (
	"math/bits"
)

// source: https://leetcode.com/problems/reverse-bits/

func reverseBits(num uint32) uint32 {
	return bits.Reverse32(num)
}