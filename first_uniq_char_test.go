package awesomeProject_test

import (
	. "github.com/colins/awesome_project"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("FirstUniqChar", func() {
	It("returns 0 for leetcode", func() {
		Expect(FirstUniqChar("leetcode")).To(Equal(0))
	})
	It("returns 2 for loveleetcode", func() {
		Expect(FirstUniqChar("loveleetcode")).To(Equal(2))
	})
})
