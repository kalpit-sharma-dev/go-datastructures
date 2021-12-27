package tree

//Find Height of tree
// Height of tree = max nodes in a single path

func solve(A *treeNode) int {

	if A == nil {
		return 0
	}

	lefth := solve(A.left)
	righth := solve(A.right)
	return 1 + max(lefth, righth)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
