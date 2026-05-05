// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/clear-street/clear-street-cli/internal/mocktest"
)

func TestV1InstrumentsGetInstrumentByID(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:instruments", "get-instrument-by-id",
			"--instrument-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--include-options-expiry-dates=true",
		)
	})
}

func TestV1InstrumentsGetInstruments(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:instruments", "get-instruments",
			"--easy-to-borrow=true",
			"--id-filter", "id_filter",
			"--instrument-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--instrument-type", "COMMON_STOCK",
			"--is-liquidation-only=true",
			"--is-marginable=true",
			"--is-restricted=true",
			"--is-short-prohibited=true",
			"--is-threshold-security=true",
			"--page-size", "1",
			"--page-token", "U3RhaW5sZXNzIHJvY2tz",
		)
	})
}

func TestV1InstrumentsSearchInstruments(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:instruments", "search-instruments",
			"--q", "q",
			"--asset-class", "asset_class",
			"--country", "country",
			"--currency", "currency",
			"--include-inactive=true",
			"--include-restricted=true",
			"--page-size", "1",
			"--page-token", "U3RhaW5sZXNzIHJvY2tz",
		)
	})
}
