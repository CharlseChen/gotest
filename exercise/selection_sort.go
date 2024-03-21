package exercise

func SelectionSort(num []int) {
	for i, _ := range num {
		k := i
		for j := i + 1; i < len(num); i++ {
			if num[i] > num[j] {
				k = j
			}
		}
		num[i], num[k] = num[k], num[i]
	}
}
