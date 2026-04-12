package sort

import (
	"math"
	"slices"

	"github.com/MorePeanuts/algorithm/gods/queues/circularqueue"
)

// CountingSort performs counting sort on an array of non-negative integers.
// Time Complexity: O(n + k), where k is the range of input values
// Space Complexity: O(k)
func CountingSort(arr []int) []int {
	if len(arr) == 0 {
		return []int{}
	}
	maxNum := slices.Max(arr)
	bucket := make([]int, maxNum+1)
	for _, num := range arr {
		bucket[num]++
	}
	res := []int{}
	for num, count := range bucket {
		for ; count > 0; count-- {
			res = append(res, num)
		}
	}
	return res
}

// BucketSort performs bucket sort on an array of non-negative integers.
// It uses 10 buckets by default and uses insertion sort for each bucket.
// Time Complexity: O(n + k), where k is the number of buckets (average case)
// Space Complexity: O(n + k)
func BucketSort(arr []int) []int {
	if len(arr) == 0 {
		return []int{}
	}

	maxNum := slices.Max(arr)
	bucketCount := 10
	if maxNum < 10 {
		bucketCount = maxNum + 1
	}

	// Create buckets
	buckets := make([][]int, bucketCount)
	for i := range buckets {
		buckets[i] = []int{}
	}

	// Distribute elements into buckets
	interval := float64(maxNum+1) / float64(bucketCount)
	for _, num := range arr {
		bucketIndex := int(float64(num) / interval)
		// Handle edge case where num == maxNum
		if bucketIndex >= bucketCount {
			bucketIndex = bucketCount - 1
		}
		buckets[bucketIndex] = append(buckets[bucketIndex], num)
	}

	// Sort each bucket and concatenate results
	res := make([]int, 0, len(arr))
	for _, bucket := range buckets {
		// Use insertion sort for small buckets (or slices.Sort for simplicity)
		slices.Sort(bucket)
		res = append(res, bucket...)
	}

	return res
}

// RadixSort performs LSD (Least Significant Digit) radix sort on an array of non-negative integers.
// Time Complexity: O(d * (n + k)), where d is the number of digits, k is the radix (10)
// Space Complexity: O(n + k)
func RadixSort(arr []int) {
	if len(arr) == 0 {
		return
	}

	buckets := make([]*circularqueue.Queue[int], 0, 10)
	for range 10 {
		buckets = append(buckets, circularqueue.New[int]())
	}

	maxNum := slices.Max(arr)
	digits := 0
	temp := maxNum
	for temp != 0 {
		digits++
		temp /= 10
	}
	// Handle case where maxNum is 0
	if digits == 0 {
		digits = 1
	}

	for i := 0; i < digits; i++ {
		divisor := int(math.Pow10(i))
		for _, num := range arr {
			digit := (num / divisor) % 10
			buckets[digit].Enqueue(num)
		}
		arr = arr[:0]
		for _, bucket := range buckets {
			for !bucket.IsEmpty() {
				num, _ := bucket.Dequeue()
				arr = append(arr, num)
			}
		}
	}
}