// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/clear-street/clear-street-cli/internal/mocktest"
)

func TestV1InstrumentsGetInstrumentByID(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:instruments", "get-instrument-by-id",
			"--instrument-id", "x",
			"--include-options-expiry-dates=true",
		)
	})
}

func TestV1InstrumentsGetInstruments(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:instruments", "get-instruments",
			"--easy-to-borrow=true",
			"--instrument-id", "x",
			"--instrument-type", "COMMON_STOCK",
			"--is-liquidation-only=true",
			"--is-marginable=true",
			"--is-ptp=true",
			"--is-short-prohibited=true",
			"--is-threshold-security=true",
			"--page-size", "1",
			"--page-token", "U3RhaW5sZXNzIHJvY2tz",
		)
	})
}

func TestV1InstrumentsGetOptionContracts(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:instruments", "get-option-contracts",
			"--contract-id", "x",
			"--contract-type", "CALL",
			"--expiry", "'2019-12-27'",
			"--is-settle-on-open=true",
			"--page-size", "1",
			"--page-token", "U3RhaW5sZXNzIHJvY2tz",
			"--underlier", "underlier",
			"--underlying-instrument-id", "x",
		)
	})
}

func TestV1InstrumentsSearchInstruments(t *testing.T) {
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
			"--include-ptp=true",
			"--page-size", "1",
			"--page-token", "U3RhaW5sZXNzIHJvY2tz",
		)
	})
}
