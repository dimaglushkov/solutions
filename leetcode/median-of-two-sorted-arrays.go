package main

import (
	"sort"
)

// source: https://leetcode.com/problems/median-of-two-sorted-arrays/

// First, we need to make sure that nums1 size is always bigger or equal to nums2 size
// With this, we can take two slices median and compare them.
// If nums1 median is bigger or equal to nums2 median, than we can
// state that the left half of nums2 lies at the left half of the merged slice,
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) < len(nums2) {
		return doFindMedianSortedArray(nums2, nums1)
	} else {
		return doFindMedianSortedArray(nums1, nums2)
	}
}

func doFindMedianSortedArray(nums1 []int, nums2 []int) float64 {
	l1, l2 := len(nums1), len(nums2)

	if l2 == 0 {
		return median(nums1)
	}
	if l2 == 1 {
		var resSlice []int

		if l1 == 1 {
			resSlice = merge(nums1[0], nums2[0])
		} else if l1%2 == 0 {
			resSlice = merge(nums1[l1/2-1], nums1[l1/2], nums2[0])
		} else {
			resSlice = merge(nums1[l1/2-1], nums1[l1/2], nums1[l1/2+1], nums2[0])
		}
		return median(resSlice)
	}
	if l2 == 2 {
		var resSlice []int
		if l1 == 2 {
			resSlice = merge(append(nums2, nums1...)...)
		} else if l1%2 == 0 {
			resSlice = merge(append(nums2, nums1[l1/2-2:l1/2+2]...)...)
		} else {
			resSlice = merge(nums1[l1/2-1], nums1[l1/2], nums1[l1/2+1], nums2[0], nums2[1])
		}
		return median(resSlice)
	}

	mid1, mid2 := (l1-1)/2, (l2-1)/2
	if nums1[mid1] >= nums2[mid2] {
		return doFindMedianSortedArray(nums1[:l1-mid2], nums2[mid2:])
	} else {
		return doFindMedianSortedArray(nums1[mid2:], nums2[:l2-mid2])
	}
}

func merge(val ...int) []int {
	sort.Ints(val)
	return val
}

func median(nums []int) float64 {
	l := len(nums)
	if l%2 == 0 {
		return float64(nums[l/2-1]+nums[l/2]) / 2.
	} else {
		return float64(nums[l/2])
	}
}
