/*
Copyright © 2023 matsubara
*/
package cmd

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRootCommand(t *testing.T) {
	buf := new(bytes.Buffer)
	rootCmd.SetOutput(buf)

	err := rootCmd.Execute()
	assert.NoError(t, err, "Expected no error when executing root command")

	output := buf.String()
	assert.Contains(t, output, "this is test-cli", "Expected output to contain 'this is test-cli'")
}

func TestRootCommandFlags(t *testing.T) {
	assert.NotNil(t, rootCmd.Flags().Lookup("toggle"), "Expected toggle flag to be defined")
	assert.Equal(t, "t", rootCmd.Flags().Lookup("toggle").Shorthand, "Expected toggle flag shorthand to be 't'")
}
