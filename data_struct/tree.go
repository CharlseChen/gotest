package main

import (
	"container/list"
	"math"
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
	return res
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
	return res
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
	return res
}

// SequenceRange 层序遍历
func (t *TreeNode) SequenceRange(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	lst := []*TreeNode{root}
	res := ([][]int)(nil)
	for len(lst) > 0 {
		temp := ([]int)(nil)
		le := len(lst)
		for i := 0; i < le; i++ {
			e := lst[0]
			lst = lst[1:]
			if e.Left != nil {
				lst = append(lst, e.Left)
			}

			if e.Right != nil {
				lst = append(lst, e.Right)
			}
			temp = append(temp, e.Val)
		}
		res = append(res, temp)
	}
	return res
}

func defs(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	if left.Val != right.Val {
		return false
	}
	return defs(left.Left, right.Right) && defs(left.Right, right.Left)
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return defs(root.Left, root.Right)
}

func isSymmetricV2(root *TreeNode) bool {
	var queue []*TreeNode
	if root != nil {
		queue = append(queue, root.Left, root.Right)
	}
	for len(queue) > 0 {
		left := queue[0]
		right := queue[1]
		queue = queue[2:]
		if left == nil && right == nil {
			continue
		}
		if left == nil || right == nil || left.Val != right.Val {
			return false
		}
		queue = append(queue, left.Left, right.Right, left.Right, right.Left)
	}
	return true
}

func leftSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil {
		return 0
	}
	leftNum := leftSum(root.Left)
	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		return root.Left.Val
	}
	rightNum := leftSum(root.Right)

	return leftNum + rightNum
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 {
		return nil
	}
	//根节点是后序遍历最后一个节点
	rootVal := postorder[0]
	root := &TreeNode{Val: rootVal}

	if len(inorder) == 1 {
		return root
	}

	var delimterIndex int
	for delimterIndex = 0; delimterIndex < len(inorder); delimterIndex++ {
		if inorder[delimterIndex] == rootVal {
			break
		}
	}

	leftInorder := inorder[:delimterIndex]
	rightInorder := inorder[delimterIndex+1:]

	postorder = postorder[:len(postorder)-1]

	leftPostorder := postorder[:len(leftInorder)]
	rightPostorder := postorder[len(rightInorder):]

	root.Left = buildTree(leftInorder, leftPostorder)
	root.Right = buildTree(rightInorder, rightPostorder)
	return root
}

func buildTreeFromPreAndMid(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{Val: preorder[0]}

	if len(inorder) == 1 {
		return root
	}

	var delimterIndex int
	for delimterIndex = 0; delimterIndex < len(inorder); delimterIndex++ {
		if inorder[delimterIndex] == root.Val {
			break
		}
	}

	leftInorder := inorder[:delimterIndex]
	rightInorder := inorder[delimterIndex+1:]

	preorder = preorder[1:]

	leftPreorder := preorder[:len(leftInorder)]
	rightPreorder := preorder[len(rightInorder):]

	root.Left = buildTree(leftInorder, leftPreorder)
	root.Right = buildTree(rightInorder, rightPreorder)
	return root
}

func maxTree(arr []int) *TreeNode {
	if len(arr) == 0 {
		return nil
	}
	var maxV int = arr[0]
	var maxIndex int
	for i, v := range arr {
		if v > maxV {
			maxV = v
			maxIndex = i
			break
		}
	}
	return &TreeNode{
		Val:   maxV,
		Left:  maxTree(arr[:maxIndex]),
		Right: maxTree(arr[maxIndex+1:]),
	}
}

func mergeTree(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil {
		return t2
	}
	if t2 == nil {
		return t1
	}
	t1.Val += t2.Val
	t1.Left = mergeTree(t1.Left, t2.Left)
	t1.Right = mergeTree(t1.Right, t2.Right)
	return t1
}

func mergeTreeV2(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil {
		return t2
	}
	if t2 == nil {
		return t1
	}
	treeQueue := []*TreeNode{t1, t2}
	for len(treeQueue) > 0 {
		e1 := treeQueue[0]
		treeQueue = treeQueue[1:]
		e2 := treeQueue[0]
		treeQueue = treeQueue[1:]
		e1.Val += e2.Val
		if e1.Left != nil && e2.Left != nil {
			treeQueue = append(treeQueue, e1.Left, e2.Left)
		}
		if e1.Right != nil && e2.Right != nil {
			treeQueue = append(treeQueue, e1.Right, e2.Right)
		}
		if e1.Left == nil {
			e1.Left = e2.Left
		}

		if e1.Right == nil {
			e1.Right = e2.Right
		}
	}
	return t1
}

func binarySearch(t *TreeNode, v int) *TreeNode {
	if t == nil {
		return nil
	}
	if t.Val == v {
		return t
	}
	if t.Val > v {
		return binarySearch(t.Left, v)
	}
	return binarySearch(t.Right, v)
}

func binarySearchIteration(t *TreeNode, v int) *TreeNode {
	if t == nil || t.Val == v {
		return t
	}
	for t != nil {
		if t.Val > v {
			t = t.Left
		} else if t.Val < v {
			t = t.Right
		} else {
			return t
		}
	}
	return nil
}

var binarySearchTreeSlice = ([]int)(nil)

func travelBinarySearchTree(root *TreeNode) {
	if root == nil {
		return
	}
	travelBinarySearchTree(root.Left)
	binarySearchTreeSlice = append(binarySearchTreeSlice, root.Val)
	travelBinarySearchTree(root.Right)
}

// 判断是否是二叉搜索树,转换成数组判断
func isRealBinarySearchTree(root *TreeNode) bool {
	if root == nil {
		return true
	}
	travelBinarySearchTree(root)

	for i := 1; i < len(binarySearchTreeSlice); i++ {
		if binarySearchTreeSlice[i] <= binarySearchTreeSlice[i-1] {
			return false
		}
	}
	return true
}

func isRealBinarySearchTreeV2(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return checkBinarySearchTree(root, math.MinInt, math.MaxInt)
}

func checkBinarySearchTree(root *TreeNode, min, max int) bool {
	if root == nil {
		return true
	}
	if root.Val <= min || root.Val >= max {
		return false
	}
	return checkBinarySearchTree(root.Left, min, root.Val) && checkBinarySearchTree(root.Right, root.Val, max)
}

func isRealBinarySearchTreeV3(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var stack []*TreeNode
	var pre *TreeNode
	for len(stack) > 0 || root != nil {
		if root != nil {
			stack = append(stack, root)
			root = root.Left
			continue
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if pre != nil && root.Val <= pre.Val {
			return false
		}
		pre = root
		root = root.Right
	}
	return true
}

func isRealBinarySearchTreeV4(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var pre *TreeNode
	var travel func(root *TreeNode) bool
	travel = func(root *TreeNode) bool {
		if root == nil {
			return true
		}
		leftResult := travel(root.Left)
		if pre != nil && root.Val <= pre.Val {
			return false
		}
		pre = root
		rightResult := travel(root.Right)
		return leftResult && rightResult
	}
	return travel(root)
}
