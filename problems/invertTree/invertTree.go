package invertTree

import . "sample/utils/binarytree"

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Right, root.Left = invertTree(root.Left), invertTree(root.Right)

	return root
}
