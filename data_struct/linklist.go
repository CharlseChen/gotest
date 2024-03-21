package main

import (
	"fmt"
	"math"
)

type Node struct {
	Data int32
	Next *Node
}

func main() {
	var list *Node
	var list2 *Node
	list = addNode(list, 2)
	list = addNode(list, 4)
	list = addNode(list, 3)
	list2 = addNode(list2, 5)
	list2 = addNode(list2, 6)
	list2 = addNode(list2, 7)
	//list = reverseNode(list)

	cur := addTwoNumbers(list, list2)
	for cur != nil {
		fmt.Println(cur.Data)
		cur = cur.Next
	}
}

func addNode(head *Node, data int32) *Node {
	var newNode = &Node{
		Data: data,
	}
	if head == nil {
		return newNode
	}

	cur := head
	for cur.Next != nil {
		cur = cur.Next
	}
	cur.Next = newNode
	return head
}

func delNode(head *Node, data int32) *Node {
	if head == nil {
		return head
	}
	if head.Data == data {
		return head.Next
	}
	cur := head.Next
	for cur != nil {
		if cur.Next.Data == data {
			cur.Next = cur.Next.Next
			return head
		}
		cur = cur.Next
	}
	return head
}

func reverseNode(head *Node) *Node {
	if head == nil || head.Next == nil {
		return head
	}
	var prev *Node
	var cur = head
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	return prev
}

//列表中所表示的数字比较小时，可以用，大了就用不了
func addTwoNumbers(l1 *Node, l2 *Node) *Node {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	head1 := l1
	i := 0
	var num1 int64
	for head1 != nil {
		if head1.Data > 0 {
			num1 += int64(math.Pow(10, float64(i))) * int64(head1.Data)
		}
		head1 = head1.Next
		i++
	}
	j := 0
	head2 := l2
	var num2 int64
	for head2 != nil {
		if head2.Data > 0 {
			num2 += int64(math.Pow(10, float64(j))) * int64(head2.Data)
		}
		head2 = head2.Next
		j++
	}
	target := num1 + num2
	if target == 0 {
		return &Node{}
	}
	fmt.Println(target, num1, num2)
	var res = &Node{}
	max := math.Max(float64(i), float64(j))
	count := 1
	for ; max >= 0; max-- {
		sq := int64(math.Pow(10, max))
		if sq <= 0 {
			continue
		}
		p := target / sq
		target = target % int64(math.Pow(10, max))
		if p == 0 && count == 1 {
			continue
		}
		count++

		temp := &Node{Data: int32(p)}
		if res.Next != nil {
			t := res.Next
			temp.Next = t
		}
		res.Next = temp
	}
	res = res.Next

	return res
}
