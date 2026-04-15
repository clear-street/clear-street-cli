// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/clear-street-cli/internal/mocktest"
	"github.com/stainless-sdks/clear-street-cli/internal/requestflag"
)

func TestActiveV1AccountsGetAccountByID(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"active:v1:accounts", "get-account-by-id",
			"--account-id", "0",
		)
	})
}

func TestActiveV1AccountsGetAccounts(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"active:v1:accounts", "get-accounts",
			"--page-size", "1",
			"--page-token", "U3RhaW5sZXNzIHJvY2tz",
		)
	})
}

func TestActiveV1AccountsPatchAccountByID(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"active:v1:accounts", "patch-account-by-id",
			"--account-id", "0",
			"--risk", "{max_notional: '5000000.00'}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(activeV1AccountsPatchAccountByID)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"active:v1:accounts", "patch-account-by-id",
			"--account-id", "0",
			"--risk.max-notional", "5000000.00",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"risk:\n" +
			"  max_notional: '5000000.00'\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"active:v1:accounts", "patch-account-by-id",
			"--account-id", "0",
		)
	})
}
