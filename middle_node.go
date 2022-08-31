package awesomeProject

func MiddleNode(head *ListNode) *ListNode {
	var nodes []*ListNode
	nodes = make([]*ListNode, 0)

	for head != nil {
		nodes = append(nodes, head)
		head = head.Next
	}
	nodeIndex := 0
	if len(nodes)%2 == 1 {
		nodeIndex = 1
	}
	return nodes[(len(nodes)-nodeIndex)/2]
}
