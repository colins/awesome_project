package awesomeProject_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/colins/awesome_project"
)

var _ = Describe("ReverseList", func() {
	It("does", func() {
		var l1 *ListNode
		AddNode(5, &l1)
		AddNode(4, &l1)
		AddNode(3, &l1)
		AddNode(2, &l1)
		AddNode(1, &l1)

		Expect(ReverseList(l1).Val).To(Equal(5))
	})
})
