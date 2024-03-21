package exercise

func BubbleSort(nums []int) {
	n := len(nums)
	for i := n - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
}

func BubbleUpgrade(nums []int) {
	n := len(nums)
	for i := 0; i < n-1; i++ {
		var flag bool
		for j := 0; j < i; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				flag = true
			}
		}
		//有序的数组未发生交换，直接就返回
		if !flag {
			break
		}
	}
}
