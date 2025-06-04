// Package cmd /*
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"test-cli/cmd/greet"
	"test-cli/cmd/version"
)

var cfgFile string

// NewRootCmd creates and returns the root command.
func NewRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "test-cli",
		Short: "A sample CLI tool demonstrating modern Go CLI development practices",
		Long:  `test-cli is an example CLI application showcasing production-grade build automation, multi-platform distribution, and comprehensive testing.`,
		RunE: func(cmd *cobra.Command, _ []string) error {
			if _, err := fmt.Fprintf(cmd.OutOrStdout(), "this is test-cli\n"); err != nil {
				return err
			}

			return nil
		},
	}

	// Add persistent flags that will be global for your application
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.test-cli.yaml)")

	// Add local flags
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	// Add subcommands
	rootCmd.AddCommand(version.NewVersionCmd())
	rootCmd.AddCommand(greet.NewGreetCmd())

	// Initialize Viper for configuration
	cobra.OnInitialize(initConfig)

	return rootCmd
}

// rootCmd represents the base command when called without any subcommands.
var rootCmd = NewRootCmd()

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

// init is called automatically by Go when the package is initialized.
func init() {
	// This function is kept for potential future initialization needs
	// All command setup is now done in NewRootCmd()
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory
		home, err := os.UserHomeDir()
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}

		// Search config in home directory with name ".test-cli" (without extension)
		viper.AddConfigPath(home)
		viper.AddConfigPath(".")
		viper.SetConfigType("yaml")
		viper.SetConfigName(".test-cli")
	}

	// Read in environment variables that match
	viper.AutomaticEnv()

	// If a config file is found, read it in
	if err := viper.ReadInConfig(); err == nil {
		_, _ = fmt.Fprintf(os.Stderr, "Using config file: %s\n", viper.ConfigFileUsed())
	}
}
