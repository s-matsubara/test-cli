// Package greet implements the greet subcommand.
package greet

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	flagName   = "name"
	flagFormat = "format"

	defaultName   = "World"
	defaultFormat = "Hello, %s!"
)

// NewGreetCmd creates and returns the greet command.
func NewGreetCmd() *cobra.Command {
	var (
		name   string
		format string
	)

	greetCmd := &cobra.Command{
		Use:   "greet",
		Short: "Greet the user with a customizable message",
		Long:  `Greet the user with a customizable message. You can specify the name and format of the greeting.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			// If name is not provided via flag, check if it's in args
			if name == "" && len(args) > 0 {
				name = args[0]
			}

			// If name is still not provided, use default
			if name == "" {
				name = defaultName
			}

			// Format the greeting
			greeting := fmt.Sprintf(format, name)

			// Print the greeting
			if _, err := fmt.Fprintf(cmd.OutOrStdout(), "%s\n", greeting); err != nil {
				return err
			}

			return nil
		},
	}

	// Add local flags
	greetCmd.Flags().StringVarP(&name, flagName, "n", "", fmt.Sprintf("Name to greet (default is %q)", defaultName))
	greetCmd.Flags().StringVarP(&format, flagFormat, "f", defaultFormat, "Greeting format (use %s for the name)")

	// Bind flags to viper and handle possible errors
	cobra.CheckErr(viper.BindPFlag("greet."+flagName, greetCmd.Flags().Lookup(flagName)))
	cobra.CheckErr(viper.BindPFlag("greet."+flagFormat, greetCmd.Flags().Lookup(flagFormat)))

	return greetCmd
}
