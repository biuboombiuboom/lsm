package storage

//TreeNode t
type TreeNode struct {
	dp    int8
	Val   Value
	Left  *TreeNode
	Right *TreeNode
}

//Value s
type Value struct {
	Key   int64
	Value []byte
}

func (root *TreeNode) getDp() int8 {
	if root == nil {
		return 0
	}
	return root.dp
}

func (root *TreeNode) isBlance() bool {
	diffDP := root.Left.getDp() - root.Right.getDp()
	return diffDP == 1 || diffDP == -1

}

func insert(root *TreeNode, val Value) {
	if root == nil {
		root = &TreeNode{
			Val: val,
		}
		return
	}
	if root.Val.Key == val.Key {
		return
	}
	if val.Key < root.Val.Key {
		insert(root.Left, val)
	}
	insert(root.Right, val)
}

//旋转链表
func (root *TreeNode) rotation(left bool) *TreeNode {
	var turn *TreeNode
	if left {
		turn = root.Left
		root.Left = turn.Right
	} else {
		turn = root.Right
		root.Right = turn.Left
	}
	root = turn
	root.dp = max(root.Left.dp, root.Right.dp) + 1
	turn.dp = max(turn.Left.dp, turn.Right.dp) + 1
	return turn
}

func max(dp1 int8, dp2 int8) int8 {
	var zeroDp int8
	if (dp1 - dp2) > zeroDp {
		return dp1
	}
	return dp2
}
