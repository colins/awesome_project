package awesomeProject_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "awesomeProject"
)

var _ = Describe("Rotate", func() {
	It("rotates", func() {
		nums := []int{1, 2, 3, 4, 5, 6, 7}
		result := []int{5, 6, 7, 1, 2, 3, 4}
		Rotate(nums, 3)
		Expect(nums).To(Equal(result))
	})
	It("rotates single", func() {
		nums := []int{1}
		result := []int{1}
		Rotate(nums, 0)
		Expect(nums).To(Equal(result))
	})
	It("rotates single", func() {
		nums := []int{-1}
		result := []int{-1}
		Rotate(nums, 2)
		Expect(nums).To(Equal(result))
	})
	It("over rotates", func() {
		nums := []int{1, 2}
		result := []int{2, 1}
		Rotate(nums, 3)
		Expect(nums).To(Equal(result))
	})
	It("over rotates", func() {
		nums := []int{1, 2}
		result := []int{2, 1}
		Rotate(nums, 5)
		Expect(nums).To(Equal(result))
	})
})
