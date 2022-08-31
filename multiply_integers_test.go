package awesomeProject_test

import (
	. "awesomeProject"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("MultiplyIntegers", func() {
	It("Does stuff", func() {
		Expect(Multiply("", "")).To(Equal(""))
	})
	It("math", func() {
		Expect((123 * 6) + (123 * 50) + (123 * 400)).To(Equal(56088))
		Expect(123 * 456).To(Equal(56088))
	})
})
