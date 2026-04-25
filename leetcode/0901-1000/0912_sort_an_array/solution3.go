package leetcode0912

func sortArray3(nums []int) []int {
	tmp := make([]int, len(nums))
	mergeSort(nums, tmp, 0, len(nums))
	return nums
}

func mergeSort(arr, tmp []int, left, right int) {
	if left < right-1 {
		mid := left + (right-left)/2
		mergeSort(arr, tmp, left, mid)
		mergeSort(arr, tmp, mid, right)
		merge(arr, tmp, left, right)
	}
}

func merge(arr, tmp []int, left, right int) {
	mid := left + (right-left)/2
	pLeft, pRight, pTmp := left, mid, 0
	for pLeft < mid && pRight < right {
		if arr[pLeft] <= arr[pRight] {
			tmp[pTmp] = arr[pLeft]
			pLeft++
		} else {
			tmp[pTmp] = arr[pRight]
			pRight++
		}
		pTmp++
	}
	for pLeft < mid {
		tmp[pTmp] = arr[pLeft]
		pLeft++
		pTmp++
	}
	for pRight < right {
		tmp[pTmp] = arr[pRight]
		pRight++
		pTmp++
	}
	copy(arr[left:right], tmp[:pTmp])
}
