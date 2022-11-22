package main

func unequalTriplets(nums []int) int {
	n := len(nums)
	res := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if nums[i] != nums[j] {
				for k := j; k < n; k++ {
					if nums[i] != nums[k] && nums[j] != nums[k] {
						res++
					}
				}
			}

		}
	}
	return res
}

func main() {

}
