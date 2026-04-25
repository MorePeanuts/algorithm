// Package leetcode0912 solves LeetCode 912. Sort an Array
package leetcode0912

func sortArray(nums []int) []int {
	quickSort(nums, 0, len(nums))
	return nums
}

func quickSort(nums []int, left, right int) {
	if right-left <= 10 {
		insertSort(nums[left:right])
	} else {
		choosePivot(nums, left, right)
		i, j := partition2(nums, left, right)
		quickSort(nums, left, i)
		quickSort(nums, j, right)
	}
}

func insertSort(arr []int) {
	for p := 1; p < len(arr); p++ {
		curr := arr[p]
		j := p
		for ; j > 0 && arr[j-1] > curr; j-- {
			arr[j] = arr[j-1]
		}
		arr[j] = curr
	}
}

func choosePivot(arr []int, left, right int) int {
	mid := left + (right-left)/2
	if arr[left] > arr[mid] {
		arr[left], arr[mid] = arr[mid], arr[left]
	}
	if arr[left] > arr[right-1] {
		arr[left], arr[right-1] = arr[right-1], arr[left]
	}
	if arr[mid] < arr[right-1] {
		arr[mid], arr[right-1] = arr[right-1], arr[mid]
	}
	arr[mid], arr[right-2] = arr[right-2], arr[mid]
	return arr[right-1]
}

func partition2(arr []int, left, right int) (int, int) {
	pivot := arr[right-1]
	i, j := left, right-2
	for i < j {
		for i = i + 1; arr[i] < pivot; i++ {
		}
		for j = j - 1; arr[j] > pivot; j-- {
		}
		if i < j {
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i], arr[right-1] = arr[right-1], arr[i]
	return i, i + 1
}
