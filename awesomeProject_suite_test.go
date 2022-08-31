package awesomeProject_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestAwesomeProject(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "AwesomeProject Suite")
}
