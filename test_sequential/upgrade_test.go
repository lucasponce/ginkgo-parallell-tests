package test_sequential

import (
	"fmt"
	"time"

	. "github.com/onsi/ginkgo/v2"
)

const (
	// Alias per identifying clusters
	ClusterAliasHypershiftManagedPoliciesUpgrade = "Cluster-ClusterAliasHypershiftManagedPoliciesUpgrade"
	ClusterAliasHypershiftArm                    = "Cluster-ClusterAliasHypershiftArm"
	ClusterAliasHypershiftAmd                    = "Cluster-ClusterAliasHypershiftAmd"
	ClusterAliasZeroEgressAmd                    = "Cluster-ClusterAliasZeroEgressAmd"

	TestTime = 10 * time.Second
)

var _ = Describe("Upgrade control plane and node pool test", Ordered,
	Label("upgrade-only", "slow", "hypershift"), func() {
		BeforeAll(func() {
			logTime("BEFORE ALL ---> WAITING CLUSTERS TO BE READY")
			time.Sleep(TestTime)
			logTime("BEFORE ALL ---> CLUSTERS READY")
		})
		Context("Upgrade tests", func() {
			DescribeTable("Upgrade a cluster control plane - dry run", func(clusterAlias string) {
				runTest()
			},
				Entry(nil, ClusterAliasHypershiftManagedPoliciesUpgrade),
				Entry(nil, ClusterAliasHypershiftArm),
				Entry(nil, ClusterAliasHypershiftAmd),
				Entry(nil, ClusterAliasZeroEgressAmd),
			)
			DescribeTable("Upgrade a node pool first - manual upgrade policy - should fail", func(clusterAlias string) {
				runTest()
			},
				Entry(nil, ClusterAliasHypershiftManagedPoliciesUpgrade),
				Entry(nil, ClusterAliasHypershiftArm),
				Entry(nil, ClusterAliasHypershiftAmd),
				Entry(nil, ClusterAliasZeroEgressAmd),
			)
			DescribeTable("Upgrade a node pool with automatic upgrades and minor version enabled - should fail",
				func(clusterAlias string) {
					runTest()
				},
				Entry(nil, ClusterAliasHypershiftManagedPoliciesUpgrade),
				Entry(nil, ClusterAliasHypershiftArm),
				Entry(nil, ClusterAliasHypershiftAmd),
				Entry(nil, ClusterAliasZeroEgressAmd),
			)
			DescribeTable("Upgrade control plane with automatic upgrades and minor version enabled - should fail",
				func(clusterAlias string) {
					runTest()
				},
				Entry(nil, ClusterAliasHypershiftManagedPoliciesUpgrade),
				Entry(nil, ClusterAliasHypershiftArm),
				Entry(nil, ClusterAliasHypershiftAmd),
				Entry(nil, ClusterAliasZeroEgressAmd),
			)
			DescribeTable("Upgrade control plane - manual upgrade policy", func(clusterAlias string) {
				runTest()
			},
				Entry(nil, ClusterAliasHypershiftManagedPoliciesUpgrade),
				Entry(nil, ClusterAliasHypershiftArm),
				Entry(nil, ClusterAliasZeroEgressAmd),
			)
			DescribeTable("Upgrade control plane - manual upgrade policy - migration to multi-arch",
				func(clusterAlias string) {
					runTest()
				},
				Entry(nil, ClusterAliasHypershiftAmd),
			)
			DescribeTable("Upgrade multiple node pools - manual upgrade policy", func(clusterAlias string) {
				runTest()
			},
				Entry(nil, ClusterAliasHypershiftManagedPoliciesUpgrade),
				Entry(nil, ClusterAliasHypershiftArm),
				Entry(nil, ClusterAliasHypershiftAmd),
				Entry(nil, ClusterAliasZeroEgressAmd),
			)
		})
	})

func logTime(message string) {
	fmt.Printf("[%s] %s\n", time.Now().Format(time.RFC3339), message)
}

func runTest() {
	cSpecReport := CurrentSpecReport()
	testName := cSpecReport.FullText()
	logTime("START    === " + testName)
	time.Sleep(TestTime)
	logTime("FINISH   === " + testName)
}
