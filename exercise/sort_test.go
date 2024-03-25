package exercise

import "testing"

func Test_insertion(t *testing.T) {
	var num = []int{19, 23, 9, 34, 42, 52, 1, 50}
	insertion(num)
	t.Logf("num:%v", num)
}
