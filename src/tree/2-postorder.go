package tree

func postorderTraversal(A *treeNode) []int {

	var answer []int
	PostOrderTraverse(A, &answer)
	return answer
}

func PostOrderTraverse(root *treeNode, list *[]int) {
	if root == nil {
		return
	}
	PostOrderTraverse(root.left, list)
	PostOrderTraverse(root.right, list)
	*list = append(*list, root.value)
}
