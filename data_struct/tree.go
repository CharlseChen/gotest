package main

import (
	"container/list"
)

type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Val   int
}

func (t *TreeNode) InsertNode(v int) {
	if v == 0 {
		return
	}
	if t.Val == v {
		return
	}

	if t.Val < v {
		if t.Right != nil {
			t.Right.InsertNode(v)
		} else {
			t.Right = &TreeNode{
				Val: v,
			}
		}
	} else {
		if t.Left != nil {
			t.Left.InsertNode(v)
		} else {
			t.Left = &TreeNode{
				Val: v,
			}
		}
	}
}

func (t *TreeNode) Find(v int) *TreeNode {
	if t == nil {
		return nil
	}
	if t.Val == v {
		return t
	}
	if t.Val < v {
		return t.Right.Find(v)
	}
	return t.Left.Find(v)
}

func (t *TreeNode) Delete(v int) {
	/* 删除节点 */
	cur := t
	// 若树为空，直接提前返回
	if cur == nil {
		return
	}
	// 待删除节点之前的节点位置
	var pre *TreeNode = nil
	// 循环查找，越过叶节点后跳出
	for cur != nil {
		if cur.Val == v {
			break
		}
		pre = cur
		if cur.Val < v {
			// 待删除节点在右子树中
			cur = cur.Right
		} else {
			// 待删除节点在左子树中
			cur = cur.Left
		}
	}
	// 若无待删除节点，则直接返回
	if cur == nil {
		return
	}
	// 子节点数为 0 或 1
	if cur.Left == nil || cur.Right == nil {
		var child *TreeNode = nil
		// 取出待删除节点的子节点
		if cur.Left != nil {
			child = cur.Left
		} else {
			child = cur.Right
		}
		// 删除节点 cur
		if cur != t {
			if pre.Left == cur {
				pre.Left = child
			} else {
				pre.Right = child
			}
		} else {
			// 若删除节点为根节点，则重新指定根节点
			t = child
		}
		// 子节点数为 2
	} else {
		// 获取中序遍历中待删除节点 cur 的下一个节点
		tmp := cur.Right
		for tmp.Left != nil {
			tmp = tmp.Left
		}
		// 递归删除节点 tmp
		t.Delete(tmp.Val)
		// 用 tmp 覆盖 cur
		cur.Val = tmp.Val
	}
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
		if node.Left != nil {
			st.PushBack(node.Left)
		}
		if node.Right != nil {
			st.PushBack(node.Right)
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

func (t *TreeNode) PreTravelV3(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var stack = list.New() //栈
	var res []int          //结果集
	stack.PushBack(root)
	var node *TreeNode
	for stack.Len() > 0 {
		e := stack.Back()
		stack.Remove(e)     //弹出元素
		if e.Value == nil { // 如果为空，则表明是需要处理中间节点
			e = stack.Back() //弹出元素（即中间节点）
			stack.Remove(e)  //删除中间节点
			node = e.Value.(*TreeNode)
			res = append(res, node.Val) //将中间节点加入到结果集中
			continue                    //继续弹出栈中下一个节点
		}
		node = e.Value.(*TreeNode)
		//压栈顺序：右左中
		if node.Right != nil {
			stack.PushBack(node.Right)
		}
		if node.Left != nil {
			stack.PushBack(node.Left)
		}
		stack.PushBack(node) //中间节点压栈后再压入nil作为中间节点的标志符
		stack.PushBack(nil)
	}
	return nil
}

func (t *TreeNode) MidTravelV3(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var stack = list.New() //栈
	var res []int          //结果集
	stack.PushBack(root)
	var node *TreeNode
	for stack.Len() > 0 {
		e := stack.Back()
		stack.Remove(e)     //弹出元素
		if e.Value == nil { // 如果为空，则表明是需要处理中间节点
			e = stack.Back() //弹出元素（即中间节点）
			stack.Remove(e)  //删除中间节点
			node = e.Value.(*TreeNode)
			res = append(res, node.Val) //将中间节点加入到结果集中
			continue                    //继续弹出栈中下一个节点
		}
		node = e.Value.(*TreeNode)
		//压栈顺序：右中左
		if node.Right != nil {
			stack.PushBack(node.Right)
		}
		stack.PushBack(node) //中间节点压栈后再压入nil作为中间节点的标志符
		stack.PushBack(nil)
		if node.Left != nil {
			stack.PushBack(node.Left)
		}
	}
	return nil
}

func (t *TreeNode) BackTravelV3(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var stack = list.New() //栈
	var res []int          //结果集
	stack.PushBack(root)
	var node *TreeNode
	for stack.Len() > 0 {
		e := stack.Back()
		stack.Remove(e)     //弹出元素
		if e.Value == nil { // 如果为空，则表明是需要处理中间节点
			e = stack.Back() //弹出元素（即中间节点）
			stack.Remove(e)  //删除中间节点
			node = e.Value.(*TreeNode)
			res = append(res, node.Val) //将中间节点加入到结果集中
			continue                    //继续弹出栈中下一个节点
		}
		node = e.Value.(*TreeNode)
		//压栈顺序：中右左
		stack.PushBack(node) //中间节点压栈后再压入nil作为中间节点的标志符
		stack.PushBack(nil)

		if node.Right != nil {
			stack.PushBack(node.Right)
		}

		if node.Left != nil {
			stack.PushBack(node.Left)
		}
	}
	return nil
}
