/*
Copyright © 2023 matsubara
*/
package main

import (
	"os"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMainIntegration(t *testing.T) {
	if os.Getenv("INTEGRATION_TEST") != "true" {
		t.Skip("Skipping integration test; set INTEGRATION_TEST=true to run")
	}

	cmd := exec.Command("go", "build", "-o", "test-cli-test")
	err := cmd.Run()
	assert.NoError(t, err, "Failed to build test binary")
	defer func() {
		if removeErr := os.Remove("test-cli-test"); removeErr != nil {
			t.Logf("failed to remove test binary: %v", removeErr)
		}
	}()

	output, err := exec.Command("./test-cli-test").CombinedOutput()
	assert.NoError(t, err, "Failed to run test binary")
	assert.Contains(t, string(output), "this is test-cli", "Expected output to contain 'this is test-cli'")
}
