package binary_tree

import "container/list"

type TreeNode struct {
	left  *TreeNode
	right *TreeNode
	value int
}

func NewTreeNode(left *TreeNode, right *TreeNode, value int) *TreeNode {
	return &TreeNode{left: left, right: right, value: value}
}

func (node *TreeNode) GetValue() int {
	return node.value
}

func (node *TreeNode) GetLeftLeaves() *TreeNode {
	return node.left
}

func (node *TreeNode) GetRightLeaves() *TreeNode {
	return node.right
}

func (node *TreeNode) AddLeft(p *TreeNode) {
	node.left = p
}

func (node *TreeNode) AddRight(p *TreeNode) {
	node.right = p
}

func (node *TreeNode) Add(p *TreeNode) {
	if node.left == nil {
		node.AddLeft(p)
	}
	if node.right == nil {
		node.AddRight(p)
	}
}

// 二叉树操作
type BinaryTree struct {
	root *TreeNode
}

func NewBinaryTree(root *TreeNode) *BinaryTree {
	return &BinaryTree{root: root}
}

// 层序遍历,也叫广度优先搜索 (BFS)
// 广度优先搜索，其核心是利用队列的先进先出的特性，一层一层的遍历二叉树，将其每一层的节点遍历
func (tree *BinaryTree) LevelOrder() []int {
	quee := list.New()     // 定义一个队列，用来遍历节点
	nums := make([]int, 0) // 定义一个切片，用来存储遍历的值
	quee.PushBack(tree.root)
	for quee.Len() > 0 {
		node := quee.Remove(quee.Front()).(*TreeNode)
		nums = append(nums, node.value)
		if node.left != nil {
			quee.PushBack(node.left)
		}
		if node.right != nil {
			quee.PushBack(node.right)
		}
	}
	return nums
}

// 前、中、后续遍历，也叫深度优先遍历（DFS） 它体现了一种“先走到尽头，再回溯继续”的遍历方式
// 深度优先遍历，通常使用递归实现

// 前序遍历
func PreOrder(node *TreeNode, result *[]int) {
	if node == nil {
		return
	}
	// 访问优先级：根节点 -> 左子树 -> 右子树
	*result = append(*result, node.value)
	PreOrder(node.left, result)
	PreOrder(node.right, result)
}

// 中序遍历
func InOrder(node *TreeNode, result *[]int) {
	if node == nil {
		return
	}
	// 访问优先级：左子树 -> 根节点 -> 右子树
	InOrder(node.left, result)
	*result = append(*result, node.value)
	InOrder(node.right, result)
}

// 后续遍历(优先访问所有子节点， 再回溯到根节点)
func PostOrder(node *TreeNode, result *[]int) {
	if node == nil {
		return
	}

	// 访问优先级：左子树 -> 右子树 -> 根节点
	PostOrder(node.left, result)
	PostOrder(node.right, result)
	*result = append(*result, node.value)
}
