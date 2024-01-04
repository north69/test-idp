package cmd

import (
	"context"
	"github.com/spf13/cobra"
	"test-idp/internal/config"
)

var rootCmd = &cobra.Command{
	Use:           "test-idp",
	SilenceUsage:  true,
	SilenceErrors: true,
}

func Run(ctx context.Context, cnf *config.Config) error {
	rootCmd.AddCommand(
		startServerCmd(cnf),
	)

	return rootCmd.ExecuteContext(ctx)
}
