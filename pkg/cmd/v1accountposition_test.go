// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/clear-street/clear-street-cli/internal/mocktest"
)

func TestV1AccountsPositionsClosePosition(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:accounts:positions", "close-position",
			"--account-id", "0",
			"--instrument-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--cancel-orders=false",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("cancel_orders: false")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"v1:accounts:positions", "close-position",
			"--account-id", "0",
			"--instrument-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestV1AccountsPositionsClosePositions(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:accounts:positions", "close-positions",
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
			"v1:accounts:positions", "close-positions",
			"--account-id", "0",
		)
	})
}

func TestV1AccountsPositionsGetPositions(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:accounts:positions", "get-positions",
			"--account-id", "0",
			"--instrument-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--page-size", "1",
			"--page-token", "U3RhaW5sZXNzIHJvY2tz",
			"--sort-by", "SYMBOL",
			"--sort-direction", "ASC",
		)
	})
}
