package awesomeProject

import (
	"fmt"
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func SetNodeVal(val int, node **ListNode) {
	var newNode *ListNode
	newNode = new(ListNode)
	newNode.Val = val
	*node = newNode
}
func SumNodes(node *ListNode) int {
	buildAString := ""
	for node != nil {
		buildAString = fmt.Sprintf("%d%s", node.Val, buildAString)
		node = node.Next
	}
	intRet, _ := strconv.Atoi(buildAString)
	return intRet
}
func AddNode(val int, headNode **ListNode) {
	var newNode *ListNode
	newNode = new(ListNode)
	newNode.Val = val
	newNode.Next = *headNode
	*headNode = newNode
}

func OldAddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var resultNode *ListNode
	resultNode = new(ListNode)

	l1Result := SumNodes(l1)
	l2Result := SumNodes(l2)
	result := l1Result + l2Result
	resultString := fmt.Sprintf("%d", result)
	for _, ch := range resultString {
		intToSend, _ := strconv.Atoi(string(ch))
		AddNode(intToSend, &resultNode)
	}
	return resultNode
}
func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var resultNode *ListNode
	node1 := l1
	node2 := l2
	remainder := 0
	var quotient int
	var resultInts []int
	resultInts = make([]int, 0)
	for node1 != nil || node2 != nil {
		node1Val := 0
		node2Val := 0
		if node1 != nil {
			node1Val = node1.Val
		}
		if node2 != nil {
			node2Val = node2.Val
		}
		quotient = (node1Val + node2Val + remainder) % 10
		remainder = ((node1Val + node2Val + remainder) - quotient) / 10
		resultInts = append(resultInts, quotient)
		if node1 != nil {
			node1 = node1.Next
		}
		if node2 != nil {
			node2 = node2.Next
		}
	}
	if remainder > 0 {
		resultInts = append(resultInts, remainder)
	}
	for i := len(resultInts) - 1; i >= 0; i-- {
		AddNode(resultInts[i], &resultNode)
	}
	return resultNode
}
