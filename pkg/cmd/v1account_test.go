// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/clear-street/clear-street-cli/internal/mocktest"
	"github.com/clear-street/clear-street-cli/internal/requestflag"
)

func TestV1AccountsGetAccountBalances(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:accounts", "get-account-balances",
			"--account-id", "0",
			"--top-margin-contributors-limit", "1",
		)
	})
}

func TestV1AccountsGetAccountByID(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:accounts", "get-account-by-id",
			"--account-id", "0",
		)
	})
}

func TestV1AccountsGetAccounts(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:accounts", "get-accounts",
			"--page-size", "1",
			"--page-token", "U3RhaW5sZXNzIHJvY2tz",
		)
	})
}

func TestV1AccountsGetPortfolioHistory(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:accounts", "get-portfolio-history",
			"--account-id", "0",
			"--start-date", "'2019-12-27'",
			"--end-date", "'2019-12-27'",
		)
	})
}

func TestV1AccountsPatchAccountByID(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:accounts", "patch-account-by-id",
			"--account-id", "0",
			"--risk", "{max_notional: '5000000.00'}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(v1AccountsPatchAccountByID)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:accounts", "patch-account-by-id",
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
			"v1:accounts", "patch-account-by-id",
			"--account-id", "0",
		)
	})
}
