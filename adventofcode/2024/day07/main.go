package main

import (
	"fmt"
	"github.com/dimaglushkov/solutions/adventofcode/util"
	"math/big"
	"strings"
)

func part1(input [][]*big.Int) {
	res := big.NewInt(0)

	var check func(want, have *big.Int, nums []*big.Int, ci int) bool
	check = func(want, have *big.Int, nums []*big.Int, ci int) bool {
		if ci == len(nums) {
			return want.Cmp(have) == 0
		}

		if want.Cmp(have) == -1 {
			return false
		}

		newAdd := new(big.Int).Add(have, nums[ci])
		if have.Cmp(new(big.Int)) == 0 {
			have = big.NewInt(1)
		}
		newMul := new(big.Int).Mul(have, nums[ci])

		return check(want, newAdd, nums, ci+1) || check(want, newMul, nums, ci+1)
	}

	for _, i := range input {
		if check(i[0], new(big.Int), i[1:], 0) {
			res.Add(res, i[0])
		}
	}

	fmt.Println(res)
}

func part2(input [][]*big.Int) {
	res := big.NewInt(0)

	var check func(want, have *big.Int, nums []*big.Int, ci int) bool
	check = func(want, have *big.Int, nums []*big.Int, ci int) bool {
		if ci == len(nums) {
			return want.Cmp(have) == 0
		}

		if want.Cmp(have) == -1 {
			return false
		}

		var addRes, mulRes, appendRes = nums[ci], nums[ci], nums[ci]

		if have.Cmp(new(big.Int)) != 0 {
			addRes = new(big.Int).Add(have, nums[ci])
			mulRes = new(big.Int).Mul(have, nums[ci])
			appendRes, _ = new(big.Int).SetString(string(new(big.Int).Set(nums[ci]).Append([]byte(have.Text(10)), 10)), 10)
		}

		return check(want, addRes, nums, ci+1) || check(want, mulRes, nums, ci+1) || check(want, appendRes, nums, ci+1)
	}

	for _, i := range input {
		if check(i[0], new(big.Int), i[1:], 0) {
			res.Add(res, i[0])
		}
	}

	fmt.Println(res)
}

func main() {
	lines := util.ReadInput(false)
	input := make([][]*big.Int, 0, len(lines))

	for _, line := range lines {
		inp := make([]*big.Int, 0)
		for _, x := range strings.Split(strings.Replace(line, ":", "", 1), " ") {
			inp = append(inp, big.NewInt(int64(util.Atoi(x))))
		}
		input = append(input, inp)

	}

	part1(input)
	part2(input)
}
