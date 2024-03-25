package exercise

func SelectionSort(num []int) {
	for i, _ := range num {
		k := i
		for j := i + 1; j < len(num); j++ {
			if num[k] > num[j] {
				k = j
			}
		}
		num[i], num[k] = num[k], num[i]
	}
}
