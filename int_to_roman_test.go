package awesomeProject_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "awesomeProject"
)

var _ = Describe("IntToRoman", func() {
	It("does 1", func() {
		Expect(IntToRoman(1)).To(Equal("I"))
	})
	It("does 2", func() {
		Expect(IntToRoman(2)).To(Equal("II"))
	})
	It("does 3", func() {
		Expect(IntToRoman(3)).To(Equal("III"))
	})
	It("does 4", func() {
		Expect(IntToRoman(4)).To(Equal("IV"))
	})
	It("does 5", func() {
		Expect(IntToRoman(5)).To(Equal("V"))
	})
	It("does 58", func() {
		Expect(IntToRoman(58)).To(Equal("LVIII"))
	})
	It("does 1994", func() {
		Expect(IntToRoman(1994)).To(Equal("MCMXCIV"))
	})
})
