package tree

import (
	"github.com/pw1992/weixin/kernel"
	"testing"
)

func TestGenTree(t *testing.T) {
	root := &Root{root: nil}
	s := []string{"app", "apple", "approachn", "apparent", "apply", "appearance"}
	for _, v := range s {
		r := []rune(v)
		for _, runeVal := range r {
			root.CreateTree(runeVal)
		}
	}

	kernel.DD(root.Find(root.root, rune('p')))
	//
	//root := &Root{
	//	root: nil,
	//}
	//s := []int{10, 20, 9, 5, 50, 60, 1, 0, 66, 55}
	//for _, v := range s {
	//	root.CreateTree(v)
	//}
	//
	//find := root.Find(root.root, 1)
	//kernel.DD(find)
	//
	//kernel.DD(root.root.Value, root.root.Left.Value, root.root.Left.Left.Left.NodeIndex, root.root.Right.Value, root.root.Left.Left.Value)
}
