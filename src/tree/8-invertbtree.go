package tree

func invertTree(A *treeNode) *treeNode {

	if A == nil {
		return A
	}
	left := invertTree(A.left)
	right := invertTree(A.right)
	A.left = right
	A.right = left
	return A

}
