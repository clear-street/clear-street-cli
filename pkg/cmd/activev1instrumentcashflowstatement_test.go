// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/clear-street/clear-street-cli/internal/mocktest"
)

func TestActiveV1InstrumentsCashFlowStatementsGetInstrumentCashFlowStatements(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"active:v1:instruments:cash-flow-statements", "get-instrument-cash-flow-statements",
			"--security-id-source", "CMS",
			"--security-id", "security_id",
			"--from-date", "from_date",
			"--page-size", "1",
			"--page-token", "U3RhaW5sZXNzIHJvY2tz",
			"--to-date", "to_date",
		)
	})
}
