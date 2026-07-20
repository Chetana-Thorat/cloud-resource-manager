package cmd

import (
	"fmt"
	"time"

	"github.com/Chetana-Thorat/cloud-resource-manager/internal/services"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var cloudProvider = services.NewMockCloudProvider()

var awsCmd = &cobra.Command{
	Use:   "aws",
	Short: "AWS resource operations",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Use one of: create, list, delete")
	},
}

var awsCreateCmd = &cobra.Command{
	Use:   "create [name]",
	Short: "Create a mock AWS resource",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		start := time.Now()
		name := args[0]

		if err := cloudProvider.CreateResource(name); err != nil {
			if appLogger != nil {
				appLogger.Error("create resource failed",
					zap.String("resource", name),
					zap.Error(err),
				)
			}
			return err
		}

		if appLogger != nil {
			appLogger.Info("resource created",
				zap.String("resource", name),
				zap.Duration("duration", time.Since(start)),
			)
		}

		fmt.Println("created:", name)
		return nil
	},
}

var awsListCmd = &cobra.Command{
	Use:   "list",
	Short: "List mock AWS resources",
	RunE: func(cmd *cobra.Command, args []string) error {
		start := time.Now()
		resources := cloudProvider.ListResources()

		if appLogger != nil {
			appLogger.Info("resources listed",
				zap.Int("count", len(resources)),
				zap.Duration("duration", time.Since(start)),
			)
		}

		if len(resources) == 0 {
			fmt.Println("no resources found")
			return nil
		}

		for _, r := range resources {
			fmt.Println(r)
		}
		return nil
	},
}

var awsDeleteCmd = &cobra.Command{
	Use:   "delete [name]",
	Short: "Delete a mock AWS resource",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		start := time.Now()
		name := args[0]

		if err := cloudProvider.DeleteResource(name); err != nil {
			if appLogger != nil {
				appLogger.Error("delete resource failed",
					zap.String("resource", name),
					zap.Error(err),
				)
			}
			return err
		}

		if appLogger != nil {
			appLogger.Info("resource deleted",
				zap.String("resource", name),
				zap.Duration("duration", time.Since(start)),
			)
		}

		fmt.Println("deleted:", name)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(awsCmd)
	awsCmd.AddCommand(awsCreateCmd)
	awsCmd.AddCommand(awsListCmd)
	awsCmd.AddCommand(awsDeleteCmd)
}
