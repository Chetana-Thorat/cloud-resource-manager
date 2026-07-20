package cmd

import (
	"fmt"
	"time"

	"github.com/Chetana-Thorat/cloud-resource-manager/internal/services"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
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
	RunE: func(cmd *cobra.Command, args []string) error {
		start := time.Now()
		pods := k8sProvider.GetPods()

		if appLogger != nil {
			appLogger.Info("pods listed",
				zap.Int("count", len(pods)),
				zap.Duration("duration", time.Since(start)),
			)
		}

		for _, p := range pods {
			fmt.Println(p)
		}
		return nil
	},
}

var k8sDeploymentsCmd = &cobra.Command{
	Use:   "deployments",
	Short: "List deployments",
	RunE: func(cmd *cobra.Command, args []string) error {
		start := time.Now()
		deployments := k8sProvider.GetDeployments()

		if appLogger != nil {
			appLogger.Info("deployments listed",
				zap.Int("count", len(deployments)),
				zap.Duration("duration", time.Since(start)),
			)
		}

		for _, d := range deployments {
			fmt.Println(d)
		}
		return nil
	},
}

var k8sRolloutCmd = &cobra.Command{
	Use:   "rollout [name]",
	Short: "Check rollout status",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		start := time.Now()
		name := args[0]
		status := k8sProvider.RolloutStatus(name)

		if appLogger != nil {
			appLogger.Info("rollout status checked",
				zap.String("deployment", name),
				zap.String("status", status),
				zap.Duration("duration", time.Since(start)),
			)
		}

		fmt.Println(status)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(k8sCmd)
	k8sCmd.AddCommand(k8sPodsCmd)
	k8sCmd.AddCommand(k8sDeploymentsCmd)
	k8sCmd.AddCommand(k8sRolloutCmd)
}
