package test_parallel

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

var _ = DescribeTable("Upgrade a cluster control plane - dry run", func(clusterAlias string) {
	runTest()
},
	Entry(nil, ClusterAliasHypershiftManagedPoliciesUpgrade),
	Entry(nil, ClusterAliasHypershiftArm),
	Entry(nil, ClusterAliasHypershiftAmd),
	Entry(nil, ClusterAliasZeroEgressAmd),
)

var _ = DescribeTable("Upgrade a node pool first - manual upgrade policy - should fail", func(clusterAlias string) {
	runTest()
},
	Entry(nil, ClusterAliasHypershiftManagedPoliciesUpgrade),
	Entry(nil, ClusterAliasHypershiftArm),
	Entry(nil, ClusterAliasHypershiftAmd),
	Entry(nil, ClusterAliasZeroEgressAmd),
)

var _ = DescribeTable("Upgrade a node pool with automatic upgrades and minor version enabled - should fail",
	func(clusterAlias string) {
		runTest()
	},
	Entry(nil, ClusterAliasHypershiftManagedPoliciesUpgrade),
	Entry(nil, ClusterAliasHypershiftArm),
	Entry(nil, ClusterAliasHypershiftAmd),
	Entry(nil, ClusterAliasZeroEgressAmd),
)

var _ = DescribeTable("Upgrade control plane with automatic upgrades and minor version enabled - should fail",
	func(clusterAlias string) {
		runTest()
	},
	Entry(nil, ClusterAliasHypershiftManagedPoliciesUpgrade),
	Entry(nil, ClusterAliasHypershiftArm),
	Entry(nil, ClusterAliasHypershiftAmd),
	Entry(nil, ClusterAliasZeroEgressAmd),
)

var _ = DescribeTable("Upgrade control plane - manual upgrade policy", func(clusterAlias string) {
	runTest()
},
	Entry(nil, ClusterAliasHypershiftManagedPoliciesUpgrade),
	Entry(nil, ClusterAliasHypershiftArm),
	Entry(nil, ClusterAliasZeroEgressAmd),
)

var _ = DescribeTable("Upgrade control plane - manual upgrade policy - migration to multi-arch",
	func(clusterAlias string) {
		runTest()
	},
	Entry(nil, ClusterAliasHypershiftAmd),
)

var _ = DescribeTable("Upgrade multiple node pools - manual upgrade policy", func(clusterAlias string) {
	runTest()
},
	Entry(nil, ClusterAliasHypershiftManagedPoliciesUpgrade),
	Entry(nil, ClusterAliasHypershiftArm),
	Entry(nil, ClusterAliasHypershiftAmd),
	Entry(nil, ClusterAliasZeroEgressAmd),
)

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
