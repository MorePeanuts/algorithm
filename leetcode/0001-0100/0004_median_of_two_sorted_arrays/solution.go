// Package leetcode0004 solves LeetCode 4. Median of Two Sorted Arrays
package leetcode0004

import "math"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}
	len1, len2 := len(nums1), len(nums2)
	l, r := -1, len1
	for l < r-1 {
		// p1 and p2 are the rightmost points on the left side of the median.
		p1 := (l + r) / 2
		p2 := (len1+len2+1)/2 - (p1 + 1) - 1
		if nums1[p1] > nums2[p2+1] {
			// Constraint not satisfied
			r = p1
		} else {
			l = p1
		}
	}
	// l == r - 1, p1 cannot be r, otherwise the constraint condition is not satisfied
	p1 := l
	p2 := (len1+len2+1)/2 - (p1 + 1) - 1
	// -1 <= p1 <= len1-1 and -1 <= p2 <= len2-1
	nums1Left, nums1Right := 0, 0
	nums2Left, nums2Right := 0, 0
	if p1 > -1 {
		nums1Left = nums1[p1]
	} else {
		nums1Left = math.MinInt
	}
	if p1 < len1-1 {
		nums1Right = nums1[p1+1]
	} else {
		nums1Right = math.MaxInt
	}
	if p2 > -1 {
		nums2Left = nums2[p2]
	} else {
		nums2Left = math.MinInt
	}
	if p2 < len2-1 {
		nums2Right = nums2[p2+1]
	} else {
		nums2Right = math.MaxInt
	}
	leftMax := max(nums1Left, nums2Left)
	rightMin := min(nums1Right, nums2Right)
	if (len1+len2)%2 == 0 {
		return float64(leftMax+rightMin) / 2
	} else {
		return float64(leftMax)
	}
}
