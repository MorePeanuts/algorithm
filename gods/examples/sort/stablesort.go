// Package sort
package sort

func InsertSort(arr []int) {
	var j int
	for p := 1; p < len(arr); p++ {
		curr := arr[p]
		for j = p; j > 0 && arr[j-1] > curr; j-- {
			arr[j] = arr[j-1]
		}
		arr[j] = curr
	}
}

func MergeSort(arr []int) {
	tmp := make([]int, len(arr))
	mergeSort(arr, tmp, 0, len(arr))
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
	for pRight < mid {
		tmp[pTmp] = arr[pRight]
		pRight++
		pTmp++
	}
	copy(arr[left:right], tmp[:pTmp])
}
