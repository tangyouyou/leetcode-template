package leetcode

import (
	"fmt"
	"sort"
)

func FindKthLargest(nums []int, k int) int {
	nums = buildHeap(nums)

	fmt.Println("nums:", nums)

	return nums[len(nums) - k]
}

func findKthLargest(nums []int, k int) int {
	sort.Ints(nums)
	return nums[len(nums) - k]
}

func buildHeap(nums []int) []int{
	n := len(nums) - 1
	for root := n / 2; root >= 0; root-- {
		minHeap(root, n, nums)
	}
	return heapSort(nums)
}

func heapSort(c []int) []int {
	n := len(c) - 1
	for end := n; end >= 0; end--{
		if c[0] < c[end] {
			c[end],c[0] = c[0],c[end]
			minHeap(0, end - 1, c)
		}
	}
	return c
}

func minHeap(root int, end int, c []int) {
	for {
		child := 2*root + 1
		if child > end {
			break
		}
		if child + 1 <= end && c[child] > c[child+1] {
			child = child + 1
		}
		if c[root] > c[child] {
			c[root],c[child] = c[child],c[root]
			root = child
		} else {
			break
		}
	}
}
