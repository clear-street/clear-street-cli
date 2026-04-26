// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/clear-street/clear-street-cli/internal/mocktest"
)

func TestActiveV1InstrumentsGetInstrumentByID(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"active:v1:instruments", "get-instrument-by-id",
			"--security-id-source", "CMS",
			"--security-id", "security_id",
			"--include-options-expiry-dates=true",
		)
	})
}

func TestActiveV1InstrumentsGetInstruments(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"active:v1:instruments", "get-instruments",
			"--easy-to-borrow=true",
			"--id-filter", "id_filter",
			"--is-liquidation-only=true",
			"--is-marginable=true",
			"--is-restricted=true",
			"--is-short-prohibited=true",
			"--is-threshold-security=true",
			"--page-size", "1",
			"--page-token", "U3RhaW5sZXNzIHJvY2tz",
			"--security-id", "string",
			"--security-id-source", "string",
			"--security-type", "COMMON_STOCK",
		)
	})
}

func TestActiveV1InstrumentsSearch(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"active:v1:instruments", "search",
			"--q", "q",
			"--asset-class", "asset_class",
			"--country", "country",
			"--currency", "currency",
			"--cursor", "cursor",
			"--include-inactive=true",
			"--include-restricted=true",
			"--limit", "0",
		)
	})
}
