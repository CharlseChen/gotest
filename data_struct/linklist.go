package main

import "fmt"

type Node struct {
	Data int32
	Next *Node
}

func main() {
	var list *Node
	list = addNode(list, 1)
	list = addNode(list, 2)
	list = addNode(list, 3)
	list = addNode(list, 4)
	list = addNode(list, 5)
	list = delNode(list, 3)
	list = reverseNode(list)
	cur := list
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
