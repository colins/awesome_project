package awesomeProject_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "awesomeProject"
)

var _ = Describe("MiddleNode", func() {
	It("returns middle node", func() {
		var l1 *ListNode
		AddNode(5, &l1)
		AddNode(4, &l1)
		AddNode(3, &l1)
		AddNode(2, &l1)
		AddNode(1, &l1)

		Expect(MiddleNode(l1).Val).To(Equal(3))
	})
})
