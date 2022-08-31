package awesomeProject

func ReverseList(head *ListNode) *ListNode {
	var resultNode *ListNode

	for head != nil {
		AddNode(head.Val, &resultNode)
		head = head.Next
	}
	return resultNode
}
