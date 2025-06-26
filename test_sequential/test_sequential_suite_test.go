package test_sequential_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestSequential(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Test Sequential Suite")
}
