package awesomeProject_test

import (
	. "awesomeProject"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("TwoSum", func() {
	It("works out hashmaps", func() {
		var ints map[int]int
		ints = make(map[int]int)
		ints[0] = 99
		if myInt, ok := ints[0]; ok {
			Expect(myInt).To(Equal(99))
		}
	})
	It("finds indices of added ints", func() {
		inputArray := []int{2, 7, 11, 15}
		outputArray := []int{0, 1}
		Expect(TwoSum(inputArray, 9)).To(Equal(outputArray))
	})
	It("does more added ints", func() {
		inputArray := []int{2, 7, 11, 15, -8, 23, 5}
		outputArray := []int{2, 4}
		Expect(TwoSum(inputArray, 3)).To(Equal(outputArray))
	})
})
