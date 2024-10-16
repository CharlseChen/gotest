package main

import "container/list"

type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Val   int
}

var tree []*TreeNode

func (t *TreeNode) PreTravel(root *TreeNode) (res []int) {
	var traversal func(node *TreeNode)
	traversal = func(node *TreeNode) {
		if node == nil {
			return
		}
		res = append(res, node.Val)
		traversal(node.Left)
		traversal(node.Right)
	}
	traversal(root)
	return res
}

func (t *TreeNode) MidTravel(root *TreeNode) (res []int) {
	var traversal func(node *TreeNode)
	traversal = func(node *TreeNode) {
		if node == nil {
			return
		}
		traversal(node.Left)
		res = append(res, node.Val)
		traversal(node.Right)
	}
	traversal(root)
	return res
}

func (t *TreeNode) PostTravel(root *TreeNode) (res []int) {
	var traversal func(node *TreeNode)
	traversal = func(node *TreeNode) {
		if node == nil {
			return
		}
		traversal(node.Left)
		traversal(node.Right)
		res = append(res, node.Val)
	}
	traversal(root)
	return res
}

func (t *TreeNode) PreTravelV2(root *TreeNode) []int {
	var ans []int
	if root == nil {
		return ans
	}
	st := list.New()
	st.PushBack(root)
	for st.Len() > 0 {
		node := st.Remove(st.Back()).(*TreeNode)
		ans = append(ans, node.Val)
		if node.Right != nil {
			st.PushBack(node.Right)
		}
		if node.Left != nil {
			st.PushBack(node.Left)
		}
	}
	return ans
}

func (t *TreeNode) MidTravelV2(root *TreeNode) []int {
	ans := ([]int)(nil)
	if root == nil {
		return ans
	}

	st := list.New()
	cur := root

	for cur != nil || st.Len() > 0 {
		if cur != nil {
			st.PushBack(cur)
			cur = cur.Left
		} else {
			cur = st.Remove(st.Back()).(*TreeNode)
			ans = append(ans, cur.Val)
			cur = cur.Right
		}
	}

	return ans
}

func (t *TreeNode) PostTravelV2(root *TreeNode) []int {
	ans := ([]int)(nil)
	if root == nil {
		return ans
	}
	st := list.New()
	st.PushBack(root)
	for st.Len() > 0 {
		node := st.Remove(st.Back()).(*TreeNode)
		ans = append(ans, node.Val)
		if node.Right != nil {
			st.PushBack(node.Right)
		}
		if node.Left != nil {
			st.PushBack(node.Left)
		}
	}
	reverse(ans)
	return ans
}

func reverse(a []int) {
	l, r := 0, len(a)-1
	for l < r {
		a[l], a[r] = a[r], a[l]
		l, r = l+1, r-1
	}
}
