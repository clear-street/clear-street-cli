// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/clear-street/clear-street-cli/internal/mocktest"
	"github.com/clear-street/clear-street-cli/internal/requestflag"
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
			"--root", "root",
			"--underlier", "underlier",
			"--underlying-instrument-id", "x",
		)
	})
}

func TestV1InstrumentsGetPredictionCategories(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:instruments", "get-prediction-categories",
		)
	})
}

func TestV1InstrumentsGetPredictionEvents(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:instruments", "get-prediction-events",
			"--categories", "categories",
			"--close-at", "{gt: gt, gte: gte, lt: lt, lte: lte}",
			"--include-markets=true",
			"--include-settled=true",
			"--market-ids", "market_ids",
			"--open-at", "{gt: gt, gte: gte, lt: lt, lte: lte}",
			"--page-size", "1",
			"--page-token", "U3RhaW5sZXNzIHJvY2tz",
			"--search", "search",
			"--series-codes", "series_codes",
			"--sort-by", "OPEN_AT",
			"--sort-direction", "ASC",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(v1InstrumentsGetPredictionEvents)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:instruments", "get-prediction-events",
			"--categories", "categories",
			"--close-at.gt", "gt",
			"--close-at.gte", "gte",
			"--close-at.lt", "lt",
			"--close-at.lte", "lte",
			"--include-markets=true",
			"--include-settled=true",
			"--market-ids", "market_ids",
			"--open-at.gt", "gt",
			"--open-at.gte", "gte",
			"--open-at.lt", "lt",
			"--open-at.lte", "lte",
			"--page-size", "1",
			"--page-token", "U3RhaW5sZXNzIHJvY2tz",
			"--search", "search",
			"--series-codes", "series_codes",
			"--sort-by", "OPEN_AT",
			"--sort-direction", "ASC",
		)
	})
}

func TestV1InstrumentsGetPredictionMarkets(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:instruments", "get-prediction-markets",
			"--event-code", "event_code",
			"--market-ids", "market_ids",
			"--page-size", "1",
			"--page-token", "U3RhaW5sZXNzIHJvY2tz",
			"--statuses", "statuses",
		)
	})
}

func TestV1InstrumentsGetPredictionSeries(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"v1:instruments", "get-prediction-series",
			"--categories", "categories",
			"--page-size", "1",
			"--page-token", "U3RhaW5sZXNzIHJvY2tz",
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
