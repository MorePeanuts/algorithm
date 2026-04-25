package leetcode0912

func sortArray2(nums []int) []int {
	quickSort2(nums, 0, len(nums))
	return nums
}

func quickSort2(nums []int, left, right int) {
	if right-left <= 10 {
		insertSort(nums[left:right])
	} else {
		choosePivot(nums, left, right)
		i, j := partition3(nums, left, right)
		quickSort2(nums, left, i)
		quickSort2(nums, j, right)
	}
}

func partition3(arr []int, left, right int) (int, int) {
	pivot := arr[right-1]
	for i := left; i < right; {
		switch {
		case arr[i] < pivot:
			arr[left], arr[i] = arr[i], arr[left]
			left++
			i++
		case arr[i] == pivot:
			i++
		case arr[i] > pivot:
			arr[right-1], arr[i] = arr[i], arr[right-1]
			right--
		}
	}
	return left, right
}
