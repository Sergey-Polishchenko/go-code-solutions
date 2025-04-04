package testdata

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewListNode(values ...int) *ListNode {
	node := &ListNode{}
	subNode := node
	for i, v := range values {
		subNode.Val = v
		if i != len(values)-1 {
			subNode.Next = &ListNode{}
			subNode = subNode.Next
		}
	}
	return node
}
