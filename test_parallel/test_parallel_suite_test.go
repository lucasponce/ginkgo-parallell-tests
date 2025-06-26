package test_parallel

import (
	"strconv"
	"testing"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestParallel(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Test Parallel Suite")
}

var _ = SynchronizedBeforeSuite(
	func() []byte {
		logTime("FIRST PROCESS")

		logTime("BEFORE ALL ---> WAITING CLUSTERS TO BE READY")
		time.Sleep(TestTime)
		logTime("BEFORE ALL ---> CLUSTERS READY")
		
		return nil
	},
	func(sharedVariables []byte) {
		logTime("ALL PROCESSESS")
	},
)

var _ = SynchronizedAfterSuite(
	func() {
		logTime("ALL PROCESSESS")
	},
	func() {
		logTime("FIRST PROCESS")
	},
)

var _ = BeforeEach(func() {
	nProcess := GinkgoParallelProcess()
	logTime("BEFORE PROCESS: " + strconv.Itoa(nProcess))
})

var _ = AfterEach(func() {
	nProcess := GinkgoParallelProcess()
	logTime("AFTER PROCESS: " + strconv.Itoa(nProcess))

})
