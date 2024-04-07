package exercise

func partition(nums []int, left, right int) int {
	i, j := left, right
	for i < j {
		for i < j && nums[j] >= nums[left] {
			j--
		}
		for i < j && nums[i] <= nums[left] {
			i++
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	nums[i], nums[left] = nums[left], nums[i]
	return i
}

func QuickSort(nums []int, left, right int) {
	if left >= right {
		return
	}
	index := partition(nums, left, right)
	QuickSort(nums, left, index-1)
	QuickSort(nums, index+1, right)
}

//选取中位数
func getMedian(left, median, right int) int {
	l, m, r := left, median, right
	if (r < m && m < l) || (l < m && m < r) {
		return m
	} else if (m < l && l < r) || (r < l && l < m) {
		return l
	}
	return r
}

func partitionV2(nums []int, left, right int) int {
	med := getMedian(left, (left+right)/2, right)
	nums[left], nums[med] = nums[med], nums[left]
	i, j := left, right
	for i < j {
		for i < j && nums[j] >= nums[left] {
			j-- //从右向左找首个小于基准数的元素
		}
		for i < j && nums[i] <= nums[left] {
			i++ //从左向右找首个大于基准数的元素
		}
		//元素交换
		nums[i], nums[j] = nums[j], nums[i]
	}
	nums[i], nums[left] = nums[left], nums[i]
	return i
}

func QuickSortV2(nums []int, left, right int) {
	if left >= right {
		return
	}
	pivot := partitionV2(nums, left, right)
	QuickSortV2(nums, left, pivot-1)
	QuickSortV2(nums, pivot+1, right)
}

func quickSort(nums []int, left, right int) {
	// 子数组长度为 1 时终止
	for left < right {
		// 哨兵划分操作
		pivot := partitionV2(nums, left, right)
		// 对两个子数组中较短的那个执行快速排序
		if pivot-left < right-pivot {
			quickSort(nums, left, pivot-1) // 递归排序左子数组
			left = pivot + 1               // 剩余未排序区间为 [pivot + 1, right]
		} else {
			quickSort(nums, pivot+1, right) // 递归排序右子数组
			right = pivot - 1               // 剩余未排序区间为 [left, pivot - 1]
		}
	}
}
