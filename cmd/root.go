package cmd

import (
	"os"
	"time"

	"github.com/Chetana-Thorat/cloud-resource-manager/internal/logger"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var (
	rootCmd = &cobra.Command{
		Use:   "cloudctl",
		Short: "Cloud resource management CLI",
		Long:  "A Go CLI for automating AWS resource provisioning and Kubernetes operations.",
	}

	appLogger    *zap.Logger
	commandStart time.Time
)

// Execute adds all child commands to the root command and runs it.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	var err error
	appLogger, err = logger.New()
	if err != nil {
		panic(err)
	}

	rootCmd.PersistentPreRun = func(cmd *cobra.Command, args []string) {
		commandStart = time.Now()
		appLogger.Info("command started", zap.String("command", cmd.CommandPath()))
	}

	rootCmd.PersistentPostRun = func(cmd *cobra.Command, args []string) {
		appLogger.Info(
			"command finished",
			zap.String("command", cmd.CommandPath()),
			zap.Duration("duration", time.Since(commandStart)),
		)
	}
}
