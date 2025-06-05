// Package greet implements the greet subcommand.
package greet

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
				name = "World"
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
	greetCmd.Flags().StringVarP(&name, "name", "n", "", "Name to greet (default is \"World\")")
	greetCmd.Flags().StringVarP(&format, "format", "f", "Hello, %s!", "Greeting format (use %s for the name)")

	// Bind flags to viper
	_ = viper.BindPFlag("greet.name", greetCmd.Flags().Lookup("name"))
	_ = viper.BindPFlag("greet.format", greetCmd.Flags().Lookup("format"))

	return greetCmd
}
