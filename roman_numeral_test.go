package awesomeProject_test

import (
	. "awesomeProject"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("RomanNumeral", func() {
	Describe("basic stuff", func() {
		It("returns 1", func() {
			Expect(RomanToInt("I")).To(Equal(1))
		})
		It("returns 2", func() {
			Expect(RomanToInt("II")).To(Equal(2))
		})
		It("returns 5", func() {
			Expect(RomanToInt("V")).To(Equal(5))
		})
		It("returns 8", func() {
			Expect(RomanToInt("VIII")).To(Equal(8))
		})
		It("returns 4", func() {
			Expect(RomanToInt("IV")).To(Equal(4))
		})
		It("returns 9", func() {
			Expect(RomanToInt("IX")).To(Equal(9))
		})
		It("returns 40", func() {
			Expect(RomanToInt("XL")).To(Equal(40))
		})
		It("returns 90", func() {
			Expect(RomanToInt("XC")).To(Equal(90))
		})
		It("returns 400", func() {
			Expect(RomanToInt("CD")).To(Equal(400))
		})
		It("returns 900", func() {
			Expect(RomanToInt("CM")).To(Equal(900))
		})
		It("returns 58", func() {
			Expect(RomanToInt("LVIII")).To(Equal(58))
		})
		It("returns 1994", func() {
			Expect(RomanToInt("MCMXCIV")).To(Equal(1994))
		})
	})
})
