// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/clear-street/clear-street-cli/internal/mocktest"
)

func TestActiveV1InstrumentsAnalystReportingGetInstrumentAnalystConsensus(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"active:v1:instruments:analyst-reporting", "get-instrument-analyst-consensus",
			"--security-id-source", "CMS",
			"--security-id", "security_id",
			"--from", "'2019-12-27'",
			"--to", "'2019-12-27'",
		)
	})
}
