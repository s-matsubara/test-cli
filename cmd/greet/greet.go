// Package greet /*
package greet

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// Name is the name to greet
	Name string
	// Format is the greeting format
	Format string
)

// NewGreetCmd creates and returns the greet command
func NewGreetCmd() *cobra.Command {
	greetCmd := &cobra.Command{
		Use:   "greet",
		Short: "Greet the user with a customizable message",
		Long:  `Greet the user with a customizable message. You can specify the name and format of the greeting.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			// If name is not provided via flag, check if it's in args
			if Name == "" && len(args) > 0 {
				Name = args[0]
			}

			// If name is still not provided, use default
			if Name == "" {
				Name = "World"
			}

			// Format the greeting
			greeting := fmt.Sprintf(Format, Name)

			// Print the greeting
			if _, err := fmt.Fprintf(cmd.OutOrStdout(), "%s\n", greeting); err != nil {
				return err
			}

			return nil
		},
	}

	// Add local flags
	greetCmd.Flags().StringVarP(&Name, "name", "n", "", "Name to greet (default is \"World\")")
	greetCmd.Flags().StringVarP(&Format, "format", "f", "Hello, %s!", "Greeting format (use %s for the name)")

	// Bind flags to viper
	_ = viper.BindPFlag("greet.name", greetCmd.Flags().Lookup("name"))
	_ = viper.BindPFlag("greet.format", greetCmd.Flags().Lookup("format"))

	return greetCmd
}
