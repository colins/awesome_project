package awesomeProject_test

import (
	. "github.com/colins/awesome_project"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("SortedArrayToBst", func() {
	It("returns a TreeNode", func() {
		var inputArray = []int{-10, -3, 0, 5, 9}
		Expect(SortedArrayToBST(inputArray).Val).To(Equal(0))
	})
})
