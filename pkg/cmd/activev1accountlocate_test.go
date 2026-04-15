// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/clear-street-cli/internal/mocktest"
	"github.com/stainless-sdks/clear-street-cli/internal/requestflag"
)

func TestActiveV1AccountsLocatesCreateLocateRequest(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"active:v1:accounts:locates", "create-locate-request",
			"--account-id", "0",
			"--body", "{quantity: 500, symbol: AAPL, comments: Locate for earnings play, reference_id: my-locate-batch-001}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(activeV1AccountsLocatesCreateLocateRequest)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"active:v1:accounts:locates", "create-locate-request",
			"--account-id", "0",
			"--body.quantity", "500",
			"--body.symbol", "AAPL",
			"--body.comments", "Locate for earnings play",
			"--body.reference-id", "my-locate-batch-001",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"- quantity: 500\n" +
			"  symbol: AAPL\n" +
			"  comments: Locate for earnings play\n" +
			"  reference_id: my-locate-batch-001\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"active:v1:accounts:locates", "create-locate-request",
			"--account-id", "0",
		)
	})
}

func TestActiveV1AccountsLocatesGetLocateRequests(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"active:v1:accounts:locates", "get-locate-requests",
			"--account-id", "0",
			"--page-size", "1",
			"--page-token", "U3RhaW5sZXNzIHJvY2tz",
			"--reference-id", "reference_id",
			"--status", "PENDING",
		)
	})
}

func TestActiveV1AccountsLocatesUpdateLocateRequest(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"active:v1:accounts:locates", "update-locate-request",
			"--account-id", "0",
			"--accept=true",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("accept: true")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"active:v1:accounts:locates", "update-locate-request",
			"--account-id", "0",
		)
	})
}
