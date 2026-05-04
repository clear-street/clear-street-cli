// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/clear-street/clear-street-cli/internal/mocktest"
)

func TestV1InstrumentsEventsGetAllInstrumentEvents(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:instruments:events", "get-all-instrument-events",
			"--event-type", "EARNINGS",
			"--from-date", "from_date",
			"--instrument-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--to-date", "to_date",
		)
	})
}

func TestV1InstrumentsEventsGetInstrumentEvents(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:instruments:events", "get-instrument-events",
			"--instrument-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--from-date", "from_date",
			"--to-date", "to_date",
		)
	})
}
