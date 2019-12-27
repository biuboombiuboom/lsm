package storage

func init() {
	root = &TreeNode{}
}

var (
	root *TreeNode
)

const (
	maxDP = 100
)

//NewValue put a new vlue to tree
func NewValue(value Value) {
	insert(root, value)
	if !root.isBlance() {
		leftDeflection := value.Key < root.Left.Val.Key
		root.rotation(leftDeflection)
	}
}
