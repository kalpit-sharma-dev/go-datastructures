package tree

type treeNode struct {
	left  *treeNode
	value int
	right *treeNode
}

func treeNode_new(val int) *treeNode {
	var node *treeNode = new(treeNode)
	node.value = val
	node.left = nil
	node.right = nil
	return node
}

func inorderTraversal(A *treeNode) []int {
	var answer []int
	InorderTraversal(A, &answer)
	return answer
}

func InorderTraversal(root *treeNode, list *[]int) {
	if root == nil {
		return
	}
	InorderTraversal(root.left, list)
	*list = append(*list, root.value)

	InorderTraversal(root.right, list)

}
