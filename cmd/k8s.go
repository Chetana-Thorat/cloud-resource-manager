package cmd

import (
	"fmt"

	"github.com/Chetana-Thorat/cloud-resource-manager/internal/services"
	"github.com/spf13/cobra"
)

var k8sProvider = services.NewMockKubernetesProvider()

var k8sCmd = &cobra.Command{
	Use:   "k8s",
	Short: "Kubernetes operations",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Use one of: pods, deployments, rollout")
	},
}

var k8sPodsCmd = &cobra.Command{
	Use:   "pods",
	Short: "List pods",
	Run: func(cmd *cobra.Command, args []string) {
		for _, p := range k8sProvider.GetPods() {
			fmt.Println(p)
		}
	},
}

var k8sDeploymentsCmd = &cobra.Command{
	Use:   "deployments",
	Short: "List deployments",
	Run: func(cmd *cobra.Command, args []string) {
		for _, d := range k8sProvider.GetDeployments() {
			fmt.Println(d)
		}
	},
}

var k8sRolloutCmd = &cobra.Command{
	Use:   "rollout [name]",
	Short: "Check rollout status",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(k8sProvider.RolloutStatus(args[0]))
	},
}

func init() {
	rootCmd.AddCommand(k8sCmd)
	k8sCmd.AddCommand(k8sPodsCmd)
	k8sCmd.AddCommand(k8sDeploymentsCmd)
	k8sCmd.AddCommand(k8sRolloutCmd)
}
