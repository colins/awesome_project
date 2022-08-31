package awesomeProject_test

import (
	. "awesomeProject"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("PlusOne", func() {
	It("adds one to simple array", func() {
		var input = []int{1, 2, 3}
		var result = []int{1, 2, 4}
		Expect(PlusOne(input)).To(Equal(result))
	})
	It("adds one to simple array", func() {
		var input = []int{4, 3, 2, 1}
		var result = []int{4, 3, 2, 2}
		Expect(PlusOne(input)).To(Equal(result))
	})
	It("adds one to simple array", func() {
		var input = []int{9}
		var result = []int{1, 0}
		Expect(PlusOne(input)).To(Equal(result))
	})
	It("adds one to simple array", func() {
		var input = []int{1, 9}
		var result = []int{2, 0}
		Expect(PlusOne(input)).To(Equal(result))
	})
	It("adds one to simple array", func() {
		var input = []int{9, 9}
		var result = []int{1, 0, 0}
		Expect(PlusOne(input)).To(Equal(result))
	})
})
