package exercise

func insertion(num []int) {
	if len(num) <= 1 {
		return
	}
	//已排好序的区间为[0,i-1]
	for i := 1; i < len(num); i++ {
		base := num[i] //num[i]用于存放在已排好序的挤出来的一个最大的数
		j := i - 1
		for j >= 0 && num[j] > base {
			num[j+1] = num[j]
			j--
		}
		num[j+1] = base
	}
}
