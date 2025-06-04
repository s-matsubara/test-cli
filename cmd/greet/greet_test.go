/*
Copyright © 2023 matsubara

This file contains tests for the greet command.
*/
package greet

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGreetCommand(t *testing.T) {
	// Save original values
	oldName := Name
	oldFormat := Format

	// Restore original values after test
	defer func() {
		Name = oldName
		Format = oldFormat
	}()

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
			format:   "Hello, %s!",
			expected: "Hello, World!\n",
		},
		{
			name:     "Custom name via arg",
			args:     []string{"User"},
			nameFlag: "",
			format:   "Hello, %s!",
			expected: "Hello, User!\n",
		},
		{
			name:     "Custom name via flag",
			args:     []string{},
			nameFlag: "Friend",
			format:   "Hello, %s!",
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
				cmd.Flags().Set("name", tt.nameFlag)
			}
			if tt.format != "Hello, %s!" {
				cmd.Flags().Set("format", tt.format)
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
	nameFlag := cmd.Flags().Lookup("name")
	assert.NotNil(t, nameFlag, "Expected name flag to be defined")
	assert.Equal(t, "n", nameFlag.Shorthand, "Expected name flag shorthand to be 'n'")

	// Test format flag
	formatFlag := cmd.Flags().Lookup("format")
	assert.NotNil(t, formatFlag, "Expected format flag to be defined")
	assert.Equal(t, "f", formatFlag.Shorthand, "Expected format flag shorthand to be 'f'")
}
