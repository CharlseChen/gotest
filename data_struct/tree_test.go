package main

import (
	"testing"
	"fmt"
)

func Test_Tree(t *testing.T) {
	n := &TreeNode{
		Val: 6,
	}
	n.InsertNode(2)
	n.InsertNode(3)
	n.InsertNode(1)
	n.InsertNode(4)
	n.InsertNode(8)
	n.InsertNode(5)
	n.InsertNode(9)
	n.InsertNode(10)
	n.InsertNode(11)

	//t.Log(n.Find(99))
	//t.Log(n.Find(1))
	//t.Log(n.Find(4))
	//fmt.Printf("%v", n.MidTravelV2(n))
	//n.Delete(76)
	fmt.Printf("%v", n.PreTravelV3(n))
	//fmt.Printf("%v", n.MidTravelV2(n))
	//fmt.Printf("%v", n.BackTravelV3(n))
	//fmt.Printf("%v\n", n.SequenceRange(n))
}
