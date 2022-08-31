package awesomeProject_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "awesomeProject"
)

var _ = Describe("RepeatedCharacter", func() {
	It("returns 'c' for abccbaacz", func() {
		Expect(RepeatedCharacter("abccbaacz")).To(Equal(byte('c')))
	})
	It("returns 'd' for abcdd", func() {
		Expect(RepeatedCharacter("abcdd")).To(Equal(byte('d')))
	})

})
