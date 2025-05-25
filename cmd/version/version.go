/*
Copyright © 2023 matsubara
*/
package version

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	Version   = "dev"
	Commit    = "none"
	BuildDate = "unknown"
)

func NewVersionCmd() *cobra.Command {
	versionCmd := &cobra.Command{
		Use:   "version",
		Short: "Print the version information",
		Long:  `Print the version, commit, and build date information for the CLI`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Fprintf(cmd.OutOrStdout(), "Version: %s\n", Version)
			fmt.Fprintf(cmd.OutOrStdout(), "Commit: %s\n", Commit)
			fmt.Fprintf(cmd.OutOrStdout(), "Build Date: %s\n", BuildDate)
		},
	}

	return versionCmd
}
