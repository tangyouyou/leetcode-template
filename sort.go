// 冒泡
func bubbleSort(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}

	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[i] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}

	return nums
}

func HeapSort(arr []int) {
	if len(arr) == 0 {
		return
	}

	// 构建堆，最大长度为 length-1
	length := len(arr)
	s := length / 2
	for i := s; i >= 0; i-- {
		heap(arr, i, length - 1)
	}
	// 调整堆，从0开始调整，每次移除一个元素
	for i := length - 1; i > 0; i--{
		arr[i],arr[0] = arr[0],arr[i]
		heap(arr, 0, i-1)
	}
}

func heap(arr []int, i, end int) {
	l := 2 * i + 1
	if l > end {
		return
	}
	n := l
	r := 2 * i + 2
	if r <= end && arr[r] > arr[l] {
		n = r
	}
	if arr[i] > arr[n] {
		return
	}
	arr[i], arr[n] = arr[n], arr[i]
	heap(arr, n, end)
}

// 快速排序
func quickSort(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	tmp := make([]int, 0)
	pos := nums[0]
	left := make([]int, 0)
	right := make([]int, 0)
	for i := 1; i < len(nums); i++ {
		if nums[i] <= pos {
			left = append(left, nums[i])
		} else if nums[i] > pos {
			right = append(right, nums[i])
		}
	}

	tmp = append(quickSort(left), pos)
	tmp = append(tmp, quickSort(right)...)

	return tmp
}

// 计数排序
func countingSort(nums []int, maxValue int) []int {
	if len(nums) == 0 {
		return []int{}
	}

	bucket := make([]int, maxValue+1)
	for i := 0; i < len(nums); i++ {
		bucket[nums[i]]++
	}

	index := 0
	for i := 0; i < len(bucket); i++ {
		for bucket[i] > 0 {
			nums[index] = i
			index++
			bucket[i]--
		}
	}

	return nums
}