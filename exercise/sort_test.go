package exercise

import (
	"testing"
)

func Test_insertion(t *testing.T) {
	var num = []int{19, 23, 9, 34, 42, 52, 1, 50}
	var num1 = make([]int, 10)
	copy(num1, num)
	insertion(num1)
	t.Logf("num1:%v", num1)

	var num2 = make([]int, 10)
	copy(num2, num)
	BubbleSort(num2)
	t.Logf("num2:%v", num2)

	var num3 = make([]int, 10)
	copy(num3, num)
	SelectionSort(num3)
	t.Logf("num3:%v", num3)

	num4 := make([]int, 10)
	copy(num4, num)
	num4 = MergeSort(num4)
	t.Logf("num4:%v", num4)
}
