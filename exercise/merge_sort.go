package exercise

/*
归并排序
*/

// 归并排序函数
func MergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	left := MergeSort(arr[:mid])
	right := MergeSort(arr[mid:])

	return merge(left, right)
}

// 合并两个有序数组
func merge(left, right []int) []int {
	i, j := 0, 0
	result := make([]int, 0)

	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	// 将剩余的元素拼接到结果中
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}

func mergeV2(arr []int, left, mid, right int) {
	i, j, k := left, mid+1, 0
	tmp := make([]int, right-left+1)
	for i <= j && j <= right {
		if arr[i] <= arr[j] {
			tmp[k] = arr[i]
			i++
		} else {
			tmp[k] = arr[j]
			j++
		}
		k++
	}
	// 将左子数组和右子数组的剩余元素复制到临时数组中
	for i <= mid {
		tmp[k] = arr[i]
		i++
		k++
	}
	for j <= right {
		tmp[k] = arr[j]
		j++
		k++
	}
	// 将临时数组 tmp 中的元素复制回原数组 nums 的对应区间
	for k := 0; k < len(tmp); k++ {
		arr[left+k] = tmp[k]
	}
}
