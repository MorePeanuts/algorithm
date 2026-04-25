package leetcode0215

func findKthLargest3(nums []int, k int) int {
	return quickSelect(nums, 0, len(nums), k)
}

func quickSelect(nums []int, left, right, k int) int {
	if right-left <= 10 {
		insertSort(nums[left:right])
		return nums[left+k-1]
	} else {
		choosePivot(nums, left, right)
		i, j := partition(nums, left, right)
		if k <= i-left {
			return quickSelect(nums, left, i, k)
		} else if k <= j-left {
			return nums[i]
		} else {
			return quickSelect(nums, j, right, k-j+left)
		}
	}
}

func insertSort(nums []int) {
	for i := 1; i < len(nums); i++ {
		curr := nums[i]
		j := i
		for ; j > 0 && nums[j-1] < curr; j-- {
			nums[j] = nums[j-1]
		}
		nums[j] = curr
	}
}

func choosePivot(nums []int, left, right int) {
	mid := left + (right-left)/2
	if nums[left] > nums[mid] {
		nums[left], nums[mid] = nums[mid], nums[left]
	}
	if nums[left] > nums[right-1] {
		nums[left], nums[right-1] = nums[right-1], nums[left]
	}
	if nums[mid] < nums[right-1] {
		nums[mid], nums[right-1] = nums[right-1], nums[mid]
	}
	nums[mid], nums[right-2] = nums[right-2], nums[mid]
}

func partition(nums []int, left, right int) (int, int) {
	pivot := nums[right-1]
	for i := left; i < right; {
		switch {
		case nums[i] > pivot:
			nums[left], nums[i] = nums[i], nums[left]
			left++
			i++
		case nums[i] == pivot:
			i++
		case nums[i] < pivot:
			nums[right-1], nums[i] = nums[i], nums[right-1]
			right--
		}
	}
	return left, right
}
