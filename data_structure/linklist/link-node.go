package linklist

type LinkNode struct {
	next *LinkNode
	data int
}

func NewLinkNode(data int, next *LinkNode) *LinkNode {
	return &LinkNode{
		data: data,
		next: next,
	}
}

func (node *LinkNode) GetData() int {
	return node.data
}

func (node *LinkNode) Next() *LinkNode {
	return node.next
}
