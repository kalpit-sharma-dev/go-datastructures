package tree

func preorderTraversal(A *treeNode) []int {
	var answer []int
	PreorderTraversal(A, &answer)
	return answer
}

func PreorderTraversal(root *treeNode, list *[]int) {
	if root == nil {
		return
	}
	*list = append(*list, root.value)
	PreorderTraversal(root.left, list)

	PreorderTraversal(root.right, list)

}
