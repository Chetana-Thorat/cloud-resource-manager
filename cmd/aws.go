package cmd

import (
	"fmt"
	"os"

	"github.com/Chetana-Thorat/cloud-resource-manager/internal/services"
	"github.com/spf13/cobra"
)

var awsCmd = &cobra.Command{
	Use:   "aws",
	Short: "AWS resource operations",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Use one of: create, list, delete")
	},
}

var cloudProvider = services.NewMockCloudProvider()

var awsCreateCmd = &cobra.Command{
	Use:   "create [name]",
	Short: "Create a mock AWS resource",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if err := cloudProvider.CreateResource(args[0]); err != nil {
			fmt.Println("error:", err)
			os.Exit(1)
		}
		fmt.Println("created:", args[0])
	},
}

var awsListCmd = &cobra.Command{
	Use:   "list",
	Short: "List mock AWS resources",
	Run: func(cmd *cobra.Command, args []string) {
		resources := cloudProvider.ListResources()
		if len(resources) == 0 {
			fmt.Println("no resources found")
			return
		}
		for _, r := range resources {
			fmt.Println(r)
		}
	},
}

var awsDeleteCmd = &cobra.Command{
	Use:   "delete [name]",
	Short: "Delete a mock AWS resource",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if err := cloudProvider.DeleteResource(args[0]); err != nil {
			fmt.Println("error:", err)
			os.Exit(1)
		}
		fmt.Println("deleted:", args[0])
	},
}

func init() {
	rootCmd.AddCommand(awsCmd)
	awsCmd.AddCommand(awsCreateCmd)
	awsCmd.AddCommand(awsListCmd)
	awsCmd.AddCommand(awsDeleteCmd)
}
