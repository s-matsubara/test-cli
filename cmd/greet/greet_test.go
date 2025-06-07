/*
Copyright © 2023 matsubara

This file contains tests for the greet command.
*/
package greet

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGreetCommand(t *testing.T) {

	tests := []struct {
		name     string
		args     []string
		nameFlag string
		format   string
		expected string
	}{
		{
			name:     "Default greeting",
			args:     []string{},
			nameFlag: "",
			format:   defaultFormat,
			expected: "Hello, World!\n",
		},
		{
			name:     "Custom name via arg",
			args:     []string{"User"},
			nameFlag: "",
			format:   defaultFormat,
			expected: "Hello, User!\n",
		},
		{
			name:     "Custom name via flag",
			args:     []string{},
			nameFlag: "Friend",
			format:   defaultFormat,
			expected: "Hello, Friend!\n",
		},
		{
			name:     "Custom format",
			args:     []string{},
			nameFlag: "Developer",
			format:   "Greetings, %s!",
			expected: "Greetings, Developer!\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create buffer to capture output
			buf := new(bytes.Buffer)

			// Create command
			cmd := NewGreetCmd()
			cmd.SetOut(buf)

			// Set flags and args
			if tt.nameFlag != "" {
				err := cmd.Flags().Set(flagName, tt.nameFlag)
				require.NoError(t, err)
			}
			if tt.format != defaultFormat {
				err := cmd.Flags().Set(flagFormat, tt.format)
				require.NoError(t, err)
			}
			cmd.SetArgs(tt.args)

			// Execute command
			err := cmd.Execute()
			assert.NoError(t, err, "Expected no error when executing greet command")

			// Check output
			output := buf.String()
			assert.Equal(t, tt.expected, output, "Expected output to match")
		})
	}
}

func TestGreetCommandFlags(t *testing.T) {
	cmd := NewGreetCmd()

	// Test name flag
	nameFlag := cmd.Flags().Lookup(flagName)
	assert.NotNil(t, nameFlag, "Expected name flag to be defined")
	assert.Equal(t, "n", nameFlag.Shorthand, "Expected name flag shorthand to be 'n'")

	// Test format flag
	formatFlag := cmd.Flags().Lookup(flagFormat)
	assert.NotNil(t, formatFlag, "Expected format flag to be defined")
	assert.Equal(t, "f", formatFlag.Shorthand, "Expected format flag shorthand to be 'f'")
}
