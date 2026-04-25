package sort

func ShellSort(arr []int) {
	n := len(arr)
	var j int
	for increment := n / 2; increment > 0; increment /= 2 {
		for i := increment; i < n; i++ {
			curr := arr[i]
			for j = i; j >= increment && arr[j-increment] > curr; j -= increment {
				arr[j] = arr[j-increment]
			}
			arr[j] = curr
		}
	}
}

func HeapSort(arr []int) {
	percDown := func(nums []int, i, n int) {
		var curr, child int
		for curr = nums[i]; 2*i+1 < n; i = child {
			child = 2*i + 1
			if child < n-1 && nums[child+1] > nums[child] {
				child++
			}
			if curr < nums[child] {
				nums[i] = nums[child]
			} else {
				break
			}
		}
		nums[i] = curr
	}

	n := len(arr)
	for i := n/2 - 1; i >= 0; i-- {
		percDown(arr, i, n)
	}
	for i := n - 1; i > 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		percDown(arr, 0, i)
	}
}

const cutOff = 10

func QuickSort(arr []int) {
	quickSort(arr, 0, len(arr), choosePivot, partition2)
}

func quickSort(
	arr []int, left, right int,
	pivotFunc func([]int, int, int) int,
	partition func([]int, int, int) (int, int),
) {
	if right-left <= cutOff {
		InsertSort(arr[left:right])
		return
	}
	pivotFunc(arr, left, right)
	i, j := partition(arr, left, right)
	quickSort(arr, left, i, pivotFunc, partition)
	quickSort(arr, j, right, pivotFunc, partition)
}

func QuickSort3Way(arr []int) {
	quickSort(arr, 0, len(arr), choosePivot, partition3)
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
