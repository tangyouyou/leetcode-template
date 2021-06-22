package leetcode

//给你两个有序整数数组 nums1 和 nums2，请你将 nums2 合并到 nums1 中，使 nums1 成为一个有序数组。
//初始化 nums1 和 nums2 的元素数量分别为 m 和 n 。你可以假设 nums1 的空间大小等于 m + n，这样它就有足够的空间保存来自 nums2 的元素。


//输入：nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
//输出：[1,2,2,3,5,6]

func merge(nums1 []int, m int, nums2 []int, n int)  {
	iMax := m + n
	left := m - 1
	right := n -1
	i := 1

	for left >= 0 && right >= 0 {
		if nums1[left] < nums2[right] {
			nums1[iMax-i] = nums2[right]
			i++
			right--
		} else {
			nums1[iMax-i] = nums1[left]
			i++
			left--
		}
	}

	if left == -1 {
		for i := right; i >= 0; i-- {
			nums1[i] = nums2[i]
		}
	}
}