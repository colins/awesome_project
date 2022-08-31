package awesomeProject_test

import (
	"fmt"
	"strconv"

	. "awesomeProject"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("AddTwoNumbers", func() {
	It("sumsnodes", func() {
		var l1 *ListNode
		//l1 = new(ListNode)
		//l1.Val = 3
		AddNode(3, &l1)
		AddNode(4, &l1)
		AddNode(2, &l1)
		Expect(SumNodes(l1)).To(Equal(342))
	})
	It("add to linked list", func() {
		var l1 *ListNode
		AddNode(3, &l1)
		AddNode(4, &l1)
		AddNode(2, &l1)
		Expect(SumNodes(l1)).To(Equal(342))
		var l2 *ListNode
		AddNode(4, &l2)
		AddNode(6, &l2)
		AddNode(5, &l2)
		Expect(SumNodes(l2)).To(Equal(465))

		Expect(SumNodes(AddTwoNumbers(l1, l2))).To(Equal(807))

		intString := ""
		intString += fmt.Sprintf("%d", 3)
		intString += fmt.Sprintf("%d", 4)
		intString += fmt.Sprintf("%d", 2)
		//intString := strings.Trim(strings.Replace(fmt.Sprint(ints), " ", "", -1), "[]")
		Expect(intString).To(Equal("342"))
		Expect(strconv.Atoi(intString)).To(Equal(342))
	})
	It("does mod 10", func() {
		small := 193 % 10
		rest := 193 - small
		Expect(small).To(Equal(3))
		Expect(rest).To(Equal(190))
		//	[9,9,9,9,9,9,9]
		//		[9,9,9,9]
	})
	It("does does slice stuff", func() {
		var someInts []int
		someInts = make([]int, 0)
		someInts = append(someInts, 9)
		someInts = append(someInts, 12)
		someInts = append(someInts, 15)
		someInts = append(someInts, 19)
		someInts = append([]int{21}, someInts...)
		//for i := len(someInts) - 1; i >= 0; i-- {
		//	fmt.Printf("int %d\n", someInts[i])
		//}
	})
	It("works with mismatched lists", func() {
		var l1 *ListNode
		AddNode(9, &l1)
		AddNode(9, &l1)
		AddNode(9, &l1)
		AddNode(9, &l1)
		AddNode(9, &l1)
		AddNode(9, &l1)
		AddNode(9, &l1)
		Expect(SumNodes(l1)).To(Equal(9999999))
		var l2 *ListNode
		AddNode(9, &l2)
		AddNode(9, &l2)
		AddNode(9, &l2)
		AddNode(9, &l2)
		Expect(SumNodes(l2)).To(Equal(9999))
		Expect(SumNodes(AddTwoNumbers(l1, l2))).To(Equal(10009998))
	})
})
