package tree

// 二叉树
type Btree struct {
	Value       rune
	Left, Right *Btree
	NodeIndex   int
}

type Root struct {
	root *Btree
}

func NewBtree(val rune) *Btree {
	return &Btree{
		Value:     0,
		Left:      nil,
		Right:     nil,
		NodeIndex: 1,
	}
}

func (r *Root) CreateTree(val rune) {
	newTree := NewBtree(val)
	if r.root == nil {
		r.root = newTree
	} else {
		CreateLeaf(r.root, newTree)
	}
}

// 创建叶子节点
func CreateLeaf(oldTree, newTree *Btree) {
	if oldTree.Value < newTree.Value {
		if oldTree.Right == nil {
			oldTree.Right = newTree
		} else {
			newTree.NodeIndex++
			CreateLeaf(oldTree.Right, newTree)
		}
	} else {
		if oldTree.Left == nil {
			oldTree.Left = newTree
		} else {
			newTree.NodeIndex++
			CreateLeaf(oldTree.Left, newTree)
		}
	}
}

func (r *Root) Find(btree *Btree, val rune) *Btree {
	if btree.Value > val {
		return r.Find(btree.Left, val)
	} else if btree.Value < val {
		return r.Find(btree.Right, val)
	} else if btree.Value == val {
		return btree
	}
	return btree
}
