package main

import "testing"

func Test_Tree(t *testing.T) {
	n := &TreeNode{
		Val: 1,
	}
	n.InsertNode(2)
	n.InsertNode(3)
	n.InsertNode(99)
	n.InsertNode(87)
	n.InsertNode(76)

	t.Log(n.Find(99))
	t.Log(n.Find(1))
	t.Log(n.Find(4))

}
