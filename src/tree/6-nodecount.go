package tree

// Count total nodes in a binary tree

func nodecount(A *treeNode) int {

	if A == nil {
		return 0
	}

	lefth := solve(A.left)
	righth := solve(A.right)
	return 1 + lefth + righth
}
