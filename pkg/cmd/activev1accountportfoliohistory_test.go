// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/clear-street/clear-street-cli/internal/mocktest"
)

func TestActiveV1AccountsPortfolioHistoryGetPortfolioHistory(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"active:v1:accounts:portfolio-history", "get-portfolio-history",
			"--account-id", "0",
			"--start-date", "'2019-12-27'",
			"--end-date", "'2019-12-27'",
		)
	})
}
