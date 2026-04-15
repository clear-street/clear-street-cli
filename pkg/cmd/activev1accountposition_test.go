// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/clear-street-cli/internal/mocktest"
)

func TestActiveV1AccountsPositionsClosePosition(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"active:v1:accounts:positions", "close-position",
			"--account-id", "0",
			"--security-id-source", "CMS",
			"--security-id", "security_id",
			"--cancel-orders=false",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("cancel_orders: false")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"active:v1:accounts:positions", "close-position",
			"--account-id", "0",
			"--security-id-source", "CMS",
			"--security-id", "security_id",
		)
	})
}

func TestActiveV1AccountsPositionsClosePositions(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"active:v1:accounts:positions", "close-positions",
			"--account-id", "0",
			"--cancel-orders=false",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("cancel_orders: false")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"active:v1:accounts:positions", "close-positions",
			"--account-id", "0",
		)
	})
}

func TestActiveV1AccountsPositionsGetPositions(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"active:v1:accounts:positions", "get-positions",
			"--account-id", "0",
			"--page-size", "1",
			"--page-token", "U3RhaW5sZXNzIHJvY2tz",
			"--security-id", "string",
			"--security-id-source", "string",
			"--sort-by", "SYMBOL",
			"--sort-direction", "ASC",
		)
	})
}
