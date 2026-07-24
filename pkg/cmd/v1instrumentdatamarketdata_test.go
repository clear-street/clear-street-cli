// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/clear-street/clear-street-cli/internal/mocktest"
)

func TestV1InstrumentDataMarketDataGetDailySummaries(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:instrument-data:market-data", "get-daily-summaries",
			"--instrument-ids", "instrument_ids",
		)
	})
}

func TestV1InstrumentDataMarketDataGetSnapshots(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:instrument-data:market-data", "get-snapshots",
			"--instrument-id", "x",
		)
	})
}
